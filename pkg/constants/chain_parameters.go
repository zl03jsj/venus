package constants

import (
	"math"

	"github.com/filecoin-project/go-state-types/abi"
	builtin0 "github.com/filecoin-project/specs-actors/actors/builtin"
)

const (
	DefaultConfidence          = uint64(5)
	DefaultMessageWaitLookback = abi.ChainEpoch(100) // in most cases, this should be enough to avoid races.
	LookbackNoLimit            = abi.ChainEpoch(-1)
)

const BlockMessageLimit = 10000

// Epochs
const TicketRandomnessLookback = abi.ChainEpoch(1)

// expect blocks number in a tipset
var ExpectedLeadersPerEpoch = builtin0.ExpectedLeadersPerEpoch

// BlockGasLimit is the maximum amount of gas that can be used to execute messages in a single block.
const (
	BlockGasLimit          = 10_000_000_000
	BlockGasTarget         = BlockGasLimit / 2
	BaseFeeMaxChangeDenom  = 8 // 12.5%
	InitialBaseFee         = 100e6
	MinimumBaseFee         = 100
	PackingEfficiencyNum   = 4
	PackingEfficiencyDenom = 5
)

const MainNetBlockDelaySecs = uint64(builtin0.EpochDurationSeconds)

// todo move this value to config
var InsecurePoStValidation = false

const (
	NoTimeout = math.MaxInt64
	NoHeight  = abi.ChainEpoch(-1)
)
