package client

import (
	govclient "github.com/okxtuta/go-anatha/x/gov/client"
	"github.com/okxtuta/go-anatha/x/params/client/cli"
	"github.com/okxtuta/go-anatha/x/params/client/rest"
)

// ProposalHandler handles param change proposals
var ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
