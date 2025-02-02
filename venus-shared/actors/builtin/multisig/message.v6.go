// FETCHED FROM LOTUS: builtin/multisig/message.go.template

package multisig

import (
	"fmt"

	"github.com/filecoin-project/go-address"
	"github.com/filecoin-project/go-state-types/abi"

	builtin6 "github.com/filecoin-project/specs-actors/v6/actors/builtin"
	init6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/init"
	multisig6 "github.com/filecoin-project/specs-actors/v6/actors/builtin/multisig"

	builtintypes "github.com/filecoin-project/go-state-types/builtin"
	"github.com/filecoin-project/venus/venus-shared/actors"
	init_ "github.com/filecoin-project/venus/venus-shared/actors/builtin/init"
	types "github.com/filecoin-project/venus/venus-shared/internal"
)

type message6 struct{ message0 }

func (m message6) Create(
	signers []address.Address, threshold uint64,
	unlockStart, unlockDuration abi.ChainEpoch,
	initialAmount abi.TokenAmount,
) (*types.Message, error) {

	lenAddrs := uint64(len(signers))

	if lenAddrs < threshold {
		return nil, fmt.Errorf("cannot require signing of more addresses than provided for multisig")
	}

	if threshold == 0 {
		threshold = lenAddrs
	}

	if m.from == address.Undef {
		return nil, fmt.Errorf("must provide source address")
	}

	// Set up constructor parameters for multisig
	msigParams := &multisig6.ConstructorParams{
		Signers:               signers,
		NumApprovalsThreshold: threshold,
		UnlockDuration:        unlockDuration,
		StartEpoch:            unlockStart,
	}

	enc, actErr := actors.SerializeParams(msigParams)
	if actErr != nil {
		return nil, actErr
	}

	// new actors are created by invoking 'exec' on the init actor with the constructor params
	execParams := &init6.ExecParams{
		CodeCID:           builtin6.MultisigActorCodeID,
		ConstructorParams: enc,
	}

	enc, actErr = actors.SerializeParams(execParams)
	if actErr != nil {
		return nil, actErr
	}

	return &types.Message{
		To:     init_.Address,
		From:   m.from,
		Method: builtintypes.MethodsInit.Exec,
		Params: enc,
		Value:  initialAmount,
	}, nil
}
