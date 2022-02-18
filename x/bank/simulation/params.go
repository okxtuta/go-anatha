package simulation

// DONTCOVER

import (
	"fmt"
	"math/rand"

	"github.com/okxtuta/go-anatha/x/bank/internal/types"
	"github.com/okxtuta/go-anatha/x/simulation"
)

const keySendEnabled = "sendenabled"

// ParamChanges defines the parameters that can be modified by param change proposals
// on the simulation
func ParamChanges(r *rand.Rand) []simulation.ParamChange {
	return []simulation.ParamChange{
		simulation.NewSimParamChange(types.ModuleName, keySendEnabled,
			func(r *rand.Rand) string {
				return fmt.Sprintf("%v", GenSendEnabled(r))
			},
		),
	}
}
