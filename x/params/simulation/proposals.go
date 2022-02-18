package simulation

import (
	simappparams "github.com/okxtuta/go-anatha/simapp/params"
	"github.com/okxtuta/go-anatha/x/simulation"
)

// OpWeightSubmitParamChangeProposal app params key for param change proposal
const OpWeightSubmitParamChangeProposal = "op_weight_submit_param_change_proposal"

// ProposalContents defines the module weighted proposals' contents
func ProposalContents(paramChanges []simulation.ParamChange) []simulation.WeightedProposalContent {
	return []simulation.WeightedProposalContent{
		{
			AppParamsKey:       OpWeightSubmitParamChangeProposal,
			DefaultWeight:      simappparams.DefaultWeightParamChangeProposal,
			ContentSimulatorFn: SimulateParamChangeProposalContent(paramChanges),
		},
	}
}
