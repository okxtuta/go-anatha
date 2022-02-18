package client

import (
	"github.com/gorilla/mux"

	"github.com/okxtuta/go-anatha/client/context"
	"github.com/okxtuta/go-anatha/client/rpc"
)

// Register routes
func RegisterRoutes(cliCtx context.CLIContext, r *mux.Router) {
	rpc.RegisterRPCRoutes(cliCtx, r)
}
