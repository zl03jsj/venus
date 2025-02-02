// FETCHED FROM LOTUS: builtin/system/state.go.template

package system

import (
	"github.com/ipfs/go-cid"

	"github.com/filecoin-project/venus/venus-shared/actors/adt"

	system9 "github.com/filecoin-project/go-state-types/builtin/v9/system"
)

var _ State = (*state9)(nil)

func load9(store adt.Store, root cid.Cid) (State, error) {
	out := state9{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make9(store adt.Store, builtinActors cid.Cid) (State, error) {
	out := state9{store: store}
	out.State = system9.State{
		BuiltinActors: builtinActors,
	}
	return &out, nil
}

type state9 struct {
	system9.State
	store adt.Store
}

func (s *state9) GetState() interface{} {
	return &s.State
}

func (s *state9) GetBuiltinActors() cid.Cid {

	return s.State.BuiltinActors

}

func (s *state9) SetBuiltinActors(c cid.Cid) error {

	s.State.BuiltinActors = c
	return nil

}
