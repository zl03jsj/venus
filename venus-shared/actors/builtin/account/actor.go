// FETCHED FROM LOTUS: builtin/account/actor.go.template

package account

import (
	"fmt"

	actorstypes "github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/venus/venus-shared/actors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/cbor"

	"github.com/filecoin-project/venus/venus-shared/actors/adt"
	types "github.com/filecoin-project/venus/venus-shared/internal"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"

	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	builtin9 "github.com/filecoin-project/go-state-types/builtin"
)

var Methods = builtin9.MethodsAccount

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := actors.GetActorMetaByCode(act.Code); ok {
		if name != actors.AccountKey {
			return nil, fmt.Errorf("actor code is not account: %s", name)
		}

		switch av {

		case actorstypes.Version8:
			return load8(store, act.Head)

		case actorstypes.Version9:
			return load9(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.AccountActorCodeID:
		return load0(store, act.Head)

	case builtin2.AccountActorCodeID:
		return load2(store, act.Head)

	case builtin3.AccountActorCodeID:
		return load3(store, act.Head)

	case builtin4.AccountActorCodeID:
		return load4(store, act.Head)

	case builtin5.AccountActorCodeID:
		return load5(store, act.Head)

	case builtin6.AccountActorCodeID:
		return load6(store, act.Head)

	case builtin7.AccountActorCodeID:
		return load7(store, act.Head)

	}

	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

func MakeState(store adt.Store, av actorstypes.Version, addr address.Address) (State, error) {
	switch av {

	case actorstypes.Version0:
		return make0(store, addr)

	case actorstypes.Version2:
		return make2(store, addr)

	case actorstypes.Version3:
		return make3(store, addr)

	case actorstypes.Version4:
		return make4(store, addr)

	case actorstypes.Version5:
		return make5(store, addr)

	case actorstypes.Version6:
		return make6(store, addr)

	case actorstypes.Version7:
		return make7(store, addr)

	case actorstypes.Version8:
		return make8(store, addr)

	case actorstypes.Version9:
		return make9(store, addr)

	}
	return nil, fmt.Errorf("unknown actor version %d", av)
}

type State interface {
	cbor.Marshaler

	PubkeyAddress() (address.Address, error)
	GetState() interface{}
}
