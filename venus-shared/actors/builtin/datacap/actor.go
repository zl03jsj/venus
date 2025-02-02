// FETCHED FROM LOTUS: builtin/datacap/actor.go.template

package datacap

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"
	actorstypes "github.com/filecoin-project/go-state-types/actors"
	builtin9 "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/go-state-types/cbor"

	"github.com/filecoin-project/venus/venus-shared/actors"
	"github.com/filecoin-project/venus/venus-shared/actors/adt"
	types "github.com/filecoin-project/venus/venus-shared/internal"
)

var (
	Address = builtin9.DatacapActorAddr
	Methods = builtin9.MethodsDatacap
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := actors.GetActorMetaByCode(act.Code); ok {
		if name != actors.DatacapKey {
			return nil, fmt.Errorf("actor code is not datacap: %s", name)
		}

		switch av {

		case actorstypes.Version9:
			return load9(store, act.Head)

		}
	}

	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

func MakeState(store adt.Store, av actorstypes.Version, governor address.Address, bitwidth uint64) (State, error) {
	switch av {

	case actorstypes.Version9:
		return make9(store, governor, bitwidth)

	default:
		return nil, fmt.Errorf("datacap actor only valid for actors v9 and above, got %d", av)

	}
}

type State interface {
	cbor.Marshaler

	ForEachClient(func(addr address.Address, dcap abi.StoragePower) error) error
	VerifiedClientDataCap(address.Address) (bool, abi.StoragePower, error)
	Governor() (address.Address, error)
	GetState() interface{}
}
