package client

import (
	govclient "github.com/okxtuta/go-anatha/x/gov/client"
	"github.com/okxtuta/go-anatha/x/upgrade/client/cli"
	"github.com/okxtuta/go-anatha/x/upgrade/client/rest"
)

var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitUpgradeProposal, rest.ProposalRESTHandler)
