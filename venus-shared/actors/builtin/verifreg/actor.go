// FETCHED FROM LOTUS: builtin/verifreg/actor.go.template

package verifreg

import (
	"fmt"

	actorstypes "github.com/filecoin-project/go-state-types/actors"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	"github.com/filecoin-project/go-state-types/cbor"

	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"

	builtin2 "github.com/filecoin-project/specs-actors/v2/actors/builtin"

	builtin3 "github.com/filecoin-project/specs-actors/v3/actors/builtin"

	builtin4 "github.com/filecoin-project/specs-actors/v4/actors/builtin"

	builtin5 "github.com/filecoin-project/specs-actors/v5/actors/builtin"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"

	builtin7 "github.com/filecoin-project/specs-actors/v7/actors/builtin"

	builtin9 "github.com/filecoin-project/go-state-types/builtin"

	verifregtypes "github.com/filecoin-project/go-state-types/builtin/v9/verifreg"
	"github.com/filecoin-project/venus/venus-shared/actors"
	"github.com/filecoin-project/venus/venus-shared/actors/adt"
	types "github.com/filecoin-project/venus/venus-shared/internal"
)

var (
	Address = builtin9.VerifiedRegistryActorAddr
	Methods = builtin9.MethodsVerifiedRegistry
)

func Load(store adt.Store, act *types.Actor) (State, error) {
	if name, av, ok := actors.GetActorMetaByCode(act.Code); ok {
		if name != actors.VerifregKey {
			return nil, fmt.Errorf("actor code is not verifreg: %s", name)
		}

		switch av {

		case actorstypes.Version8:
			return load8(store, act.Head)

		case actorstypes.Version9:
			return load9(store, act.Head)

		}
	}

	switch act.Code {

	case builtin0.VerifiedRegistryActorCodeID:
		return load0(store, act.Head)

	case builtin2.VerifiedRegistryActorCodeID:
		return load2(store, act.Head)

	case builtin3.VerifiedRegistryActorCodeID:
		return load3(store, act.Head)

	case builtin4.VerifiedRegistryActorCodeID:
		return load4(store, act.Head)

	case builtin5.VerifiedRegistryActorCodeID:
		return load5(store, act.Head)

	case builtin6.VerifiedRegistryActorCodeID:
		return load6(store, act.Head)

	case builtin7.VerifiedRegistryActorCodeID:
		return load7(store, act.Head)

	}

	return nil, fmt.Errorf("unknown actor code %s", act.Code)
}

func MakeState(store adt.Store, av actorstypes.Version, rootKeyAddress address.Address) (State, error) {
	switch av {

	case actorstypes.Version0:
		return make0(store, rootKeyAddress)

	case actorstypes.Version2:
		return make2(store, rootKeyAddress)

	case actorstypes.Version3:
		return make3(store, rootKeyAddress)

	case actorstypes.Version4:
		return make4(store, rootKeyAddress)

	case actorstypes.Version5:
		return make5(store, rootKeyAddress)

	case actorstypes.Version6:
		return make6(store, rootKeyAddress)

	case actorstypes.Version7:
		return make7(store, rootKeyAddress)

	case actorstypes.Version8:
		return make8(store, rootKeyAddress)

	case actorstypes.Version9:
		return make9(store, rootKeyAddress)

	}
	return nil, fmt.Errorf("unknown actor version %d", av)
}

type State interface {
	cbor.Marshaler

	RootKey() (address.Address, error)
	VerifiedClientDataCap(address.Address) (bool, abi.StoragePower, error)
	VerifierDataCap(address.Address) (bool, abi.StoragePower, error)
	RemoveDataCapProposalID(verifier address.Address, client address.Address) (bool, uint64, error)
	ForEachVerifier(func(addr address.Address, dcap abi.StoragePower) error) error
	ForEachClient(func(addr address.Address, dcap abi.StoragePower) error) error
	GetAllocation(clientIdAddr address.Address, allocationId verifregtypes.AllocationId) (*verifregtypes.Allocation, bool, error)
	GetAllocations(clientIdAddr address.Address) (map[verifregtypes.AllocationId]verifregtypes.Allocation, error)
	GetClaim(providerIdAddr address.Address, claimId verifregtypes.ClaimId) (*verifregtypes.Claim, bool, error)
	GetClaims(providerIdAddr address.Address) (map[verifregtypes.ClaimId]verifregtypes.Claim, error)
	GetState() interface{}
}
