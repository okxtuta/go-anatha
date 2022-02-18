package client

import (
	"github.com/okxtuta/go-anatha/x/distribution/client/cli"
	"github.com/okxtuta/go-anatha/x/distribution/client/rest"
	govclient "github.com/okxtuta/go-anatha/x/gov/client"
)

// param change proposal handler
var (
	ProposalHandler = govclient.NewProposalHandler(cli.GetCmdSubmitProposal, rest.ProposalRESTHandler)
)
