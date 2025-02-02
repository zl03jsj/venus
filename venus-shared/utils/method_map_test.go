package utils

import (
	"testing"

	tf "github.com/filecoin-project/venus/pkg/testhelpers/testflags"
	"github.com/filecoin-project/venus/venus-shared/actors"
	"github.com/filecoin-project/venus/venus-shared/types"
	"github.com/stretchr/testify/assert"
)

func TestMethodMap(t *testing.T) {
	tf.UnitTest(t)

	t.Run("Default to load mainnet v8 actors", func(t *testing.T) {
		for _, actorsMetadata := range actors.EmbeddedBuiltinActorsMetadata {
			if actorsMetadata.Network == string(types.NetworkNameMain) {
				for _, actor := range actorsMetadata.Actors {
					_, ok := MethodsMap[actor]
					assert.True(t, ok)
				}
			}
		}
	})

	t.Run("ReLoad butterflynet v8 actors", func(t *testing.T) {
		for _, actorsMetadata := range actors.EmbeddedBuiltinActorsMetadata {
			if actorsMetadata.Network == string(types.NetworkNameButterfly) {
				for _, actor := range actorsMetadata.Actors {
					_, ok := MethodsMap[actor]
					assert.False(t, ok)
				}
			}
		}

		assert.Nil(t, actors.SetNetworkBundle(int(types.NetworkButterfly)))
		ReloadMethodsMap()
		for _, actorsMetadata := range actors.EmbeddedBuiltinActorsMetadata {
			if actorsMetadata.Network == string(types.NetworkNameButterfly) {
				for _, actor := range actorsMetadata.Actors {
					_, ok := MethodsMap[actor]
					assert.True(t, ok)
				}
			}
		}
	})
}
