package statemanger

import (
	"context"
	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/venus/pkg/chain"
	"github.com/filecoin-project/venus/pkg/consensus"
	"github.com/filecoin-project/venus/pkg/fork"
	"github.com/filecoin-project/venus/pkg/specactors/builtin/market"
	"github.com/filecoin-project/venus/pkg/specactors/builtin/paych"
	appstate "github.com/filecoin-project/venus/pkg/state"
	"github.com/filecoin-project/venus/pkg/state/tree"
	"github.com/filecoin-project/venus/pkg/types"
	"github.com/filecoin-project/venus/pkg/vm"
	"github.com/filecoin-project/venus/pkg/vm/gas"
	"github.com/ipfs/go-cid"
	"go.opencensus.io/trace"
	"golang.org/x/xerrors"
	"sync"
)

// stateManagerAPI defines the methods needed from StateManager
// todo remove this code and add private interface in market and paychanel package
type IStateManager interface {
	ResolveToKeyAddress(ctx context.Context, addr address.Address, ts *types.TipSet) (address.Address, error)
	GetPaychState(ctx context.Context, addr address.Address, ts *types.TipSet) (*types.Actor, paych.State, error)
	Call(ctx context.Context, msg *types.UnsignedMessage, ts *types.TipSet) (*vm.Ret, error)
	GetMarketState(ctx context.Context, ts *types.TipSet) (market.State, error)
}

type stateComputeResult struct {
	stateRoot, receipt cid.Cid
}

var _ IStateManager = &Stmgr{}

type Stmgr struct {
	cs *chain.Store
	cp consensus.StateTransformer

	fork         fork.IFork
	gasSchedule  *gas.PricesSchedule
	syscallsImpl vm.SyscallsImpl

	// Compute StateRoot parallel safe
	stCache      map[types.TipSetKey]stateComputeResult
	chsWorkingOn map[types.TipSetKey]chan struct{}
	stLk         sync.Mutex
}

func NewStateManger(cs *chain.Store, cp consensus.StateTransformer, fork fork.IFork,
	gasSchedule *gas.PricesSchedule,
	syscallsImpl vm.SyscallsImpl) *Stmgr {
	return &Stmgr{cs: cs, fork: fork, cp: cp,
		gasSchedule:  gasSchedule,
		syscallsImpl: syscallsImpl,

		stCache:      make(map[types.TipSetKey]stateComputeResult),
		chsWorkingOn: make(map[types.TipSetKey]chan struct{}, 1),
	}
}

func (o *Stmgr) ResolveToKeyAddress(ctx context.Context, addr address.Address, ts *types.TipSet) (address.Address, error) {
	switch addr.Protocol() {
	case address.BLS, address.SECP256K1:
		return addr, nil
	case address.Actor:
		return address.Undef, xerrors.New("cannot resolve actor address to key address")
	default:
	}
	if ts == nil {
		ts = o.cs.GetHead()
	}
	_, view, err := o.ParentStateView(ctx, ts)
	if err != nil {
		return address.Undef, err
	}
	return view.ResolveToKeyAddr(ctx, addr)
}

func (o *Stmgr) GetPaychState(ctx context.Context, addr address.Address, ts *types.TipSet) (*types.Actor, paych.State, error) {
	_, view, err := o.ParentStateView(ctx, ts)
	if err != nil {
		return nil, nil, err
	}
	act, err := view.LoadActor(ctx, addr)
	if err != nil {
		return nil, nil, err
	}
	actState, err := view.LoadPaychState(ctx, act)
	if err != nil {
		return nil, nil, err
	}
	return act, actState, nil
}

func (o *Stmgr) GetMarketState(ctx context.Context, ts *types.TipSet) (market.State, error) {
	_, view, err := o.ParentStateView(ctx, ts)
	if err != nil {
		return nil, err
	}
	actState, err := view.LoadMarketState(ctx)
	if err != nil {
		return nil, err
	}
	return actState, nil
}

func (s *Stmgr) ParentStateTsk(ctx context.Context, tsk types.TipSetKey) (*types.TipSet, *tree.State, error) {
	ts, err := s.cs.GetTipSet(tsk)
	if err != nil {
		return nil, nil, xerrors.Errorf("loading tipset %s: %w", tsk, err)
	}
	return s.ParentState(ctx, ts)
}

func (s *Stmgr) ParentState(ctx context.Context, ts *types.TipSet) (*types.TipSet, *tree.State, error) {
	if ts == nil {
		ts = s.cs.GetHead()
	}
	parent, err := s.cs.GetTipSet(ts.Parents())
	if err != nil {
		return nil, nil, xerrors.Errorf("find tipset(%s) parent failed:%w",
			ts.Key().String(), err)
	}

	if stateRoot, _, err := s.RunStateTransition(ctx, parent); err != nil {
		return nil, nil, xerrors.Errorf("runstateTransition failed:%w", err)
	} else if !stateRoot.Equals(ts.At(0).ParentStateRoot) {
		return nil, nil, xerrors.Errorf("runstateTransition error, %w", consensus.ErrStateRootMismatch)
	}

	state, err := tree.LoadState(ctx, s.cs.Store(ctx), ts.At(0).ParentStateRoot)
	return parent, state, err
}

func (c *Stmgr) TipsetState(ctx context.Context, ts *types.TipSet) (cid.Cid, cid.Cid, error) {
	return c.RunStateTransition(ctx, ts)
}

func (c *Stmgr) RunStateTransition(ctx context.Context, ts *types.TipSet) (root cid.Cid, receipts cid.Cid, err error) {
	ctx, span := trace.StartSpan(ctx, "Exected.RunStateTransition")
	defer span.End()

	key := ts.Key()
	c.stLk.Lock()

	workingCh, exist := c.chsWorkingOn[key]

	if exist {
		c.stLk.Unlock()
		select {
		case <-workingCh:
			c.stLk.Lock()
		case <-ctx.Done():
			return cid.Undef, cid.Undef, ctx.Err()
		}
	}

	if meta, _ := c.cs.GetTipsetMetadata(ts); meta != nil {
		c.stLk.Unlock()
		return meta.TipSetStateRoot, meta.TipSetReceipts, nil
	}

	workingCh = make(chan struct{})
	c.chsWorkingOn[key] = workingCh
	c.stLk.Unlock()

	defer func() {
		c.stLk.Lock()
		delete(c.chsWorkingOn, key)
		if !root.Equals(cid.Undef) {
			if e := c.cs.PutTipSetMetadata(ctx, &chain.TipSetMetadata{
				TipSetStateRoot: root,
				TipSet:          ts,
				TipSetReceipts:  receipts,
			}); e != nil {
				err = e
			}
		}
		c.stLk.Unlock()
		close(workingCh)
	}()

	if ts.Height() == 0 {
		return ts.Blocks()[0].ParentStateRoot, ts.Blocks()[0].ParentMessageReceipts, nil
	}

	if root, receipts, err = c.cp.RunStateTransition(ctx, ts); err != nil {
		return cid.Undef, cid.Undef, err
	}

	return root, receipts, nil
}

// ctx context.Context, ts *types.TipSet, addr address.Address
func (c *Stmgr) GetActorAtTsk(ctx context.Context, addr address.Address, tsk types.TipSetKey) (*types.Actor, error) {
	ts, err := c.cs.GetTipSet(tsk)
	if err != nil {
		return nil, err
	}
	return c.GetActorAt(ctx, addr, ts)
}

func (c *Stmgr) GetActorAt(ctx context.Context, addr address.Address, ts *types.TipSet) (*types.Actor, error) {
	_, state, err := c.ParentState(ctx, ts)
	if err != nil {
		return nil, err
	}

	actor, find, err := state.GetActor(ctx, addr)
	if err != nil {
		return nil, err
	}

	if !find {
		return nil, types.ErrActorNotFound
	}
	return actor, nil
}

func (c *Stmgr) RunStateTransitionV2(ctx context.Context, ts *types.TipSet) (cid.Cid, cid.Cid, error) {
	ctx, span := trace.StartSpan(ctx, "Exected.RunStateTransition")
	defer span.End()

	var state stateComputeResult
	var err error
	key := ts.Key()
	c.stLk.Lock()

	cmptCh, exist := c.chsWorkingOn[key]

	if exist {
		c.stLk.Unlock()
		select {
		case <-cmptCh:
			c.stLk.Lock()
		case <-ctx.Done():
			return cid.Undef, cid.Undef, ctx.Err()
		}
	}

	if state, exist = c.stCache[key]; exist {
		c.stLk.Unlock()
		return state.stateRoot, state.receipt, nil
	}

	if meta, _ := c.cs.GetTipsetMetadata(ts); meta != nil {
		c.stLk.Unlock()
		return meta.TipSetStateRoot, meta.TipSetReceipts, nil
	}

	cmptCh = make(chan struct{})
	c.chsWorkingOn[key] = cmptCh
	c.stLk.Unlock()

	defer func() {
		c.stLk.Lock()
		delete(c.chsWorkingOn, key)
		if !state.stateRoot.Equals(cid.Undef) {
			c.stCache[key] = state
		}
		c.stLk.Unlock()
		close(cmptCh)
	}()

	if ts.Height() == 0 {
		return ts.Blocks()[0].ParentStateRoot, ts.Blocks()[0].ParentMessageReceipts, nil
	}

	if state.stateRoot, state.receipt, err = c.cp.RunStateTransition(ctx, ts); err != nil {
		return cid.Undef, cid.Undef, err
	} else if err = c.cs.PutTipSetMetadata(ctx, &chain.TipSetMetadata{
		TipSet:          ts,
		TipSetStateRoot: state.stateRoot,
		TipSetReceipts:  state.receipt,
	}); err != nil {
		return cid.Undef, cid.Undef, err
	}

	return state.stateRoot, state.receipt, nil
}

func (c *Stmgr) ParentStateViewTsk(ctx context.Context, tsk types.TipSetKey) (*types.TipSet, *appstate.View, error) {
	var ts *types.TipSet

	if tsk.Equals(types.EmptyTSK) {
		ts = c.cs.GetHead()
	}

	return c.ParentStateView(ctx, ts)
}

func (c *Stmgr) ParentStateView(ctx context.Context, ts *types.TipSet) (*types.TipSet, *appstate.View, error) {
	if ts == nil {
		ts = c.cs.GetHead()
	}

	parent, err := c.cs.GetTipSet(ts.Parents())
	if err != nil {
		return nil, nil, err
	}

	if _, _, err := c.RunStateTransition(ctx, parent); err != nil {
		return nil, nil, xerrors.Errorf("runStateTransition failed:%w", err)
	}

	view, err := c.cs.ParentStateView(ts)
	if err != nil {
		return nil, nil, xerrors.Errorf("parentStateView failed:%w", err)
	}
	return ts, view, nil
}

func (c *Stmgr) StateViewTsk(ctx context.Context, tsk types.TipSetKey) (cid.Cid, *appstate.View, error) {
	ts, err := c.cs.GetTipSet(tsk)
	if err != nil {
		return cid.Undef, nil, err
	}
	return c.StateView(ctx, ts)
}

func (c *Stmgr) StateView(ctx context.Context, ts *types.TipSet) (cid.Cid, *appstate.View, error) {
	stateCid, _, err := c.RunStateTransition(ctx, ts)
	if err != nil {
		return cid.Undef, nil, err
	}

	view, err := c.cs.StateView(ts)
	if err != nil {
		return cid.Undef, nil, err
	}
	return stateCid, view, nil
}
