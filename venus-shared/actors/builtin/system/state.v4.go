// FETCHED FROM LOTUS: builtin/system/state.go.template

package system

import (
	"github.com/ipfs/go-cid"
	"golang.org/x/xerrors"

	"github.com/filecoin-project/venus/venus-shared/actors/adt"

	system4 "github.com/filecoin-project/specs-actors/v4/actors/builtin/system"
)

var _ State = (*state4)(nil)

func load4(store adt.Store, root cid.Cid) (State, error) {
	out := state4{store: store}
	err := store.Get(store.Context(), root, &out)
	if err != nil {
		return nil, err
	}
	return &out, nil
}

func make4(store adt.Store) (State, error) {
	out := state4{store: store}
	out.State = system4.State{}
	return &out, nil
}

type state4 struct {
	system4.State
	store adt.Store
}

func (s *state4) GetState() interface{} {
	return &s.State
}

func (s *state4) GetBuiltinActors() cid.Cid {

	return cid.Undef

}

func (s *state4) SetBuiltinActors(c cid.Cid) error {

	return xerrors.New("cannot set manifest cid before v8")

}
