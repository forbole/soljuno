package proxy

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/forbole/soljuno/actions-proxy/apis"
	"github.com/forbole/soljuno/solana/client"
)

func StartProxyServer(proxy client.ClientProxy, port int) {
	// Setup the rest server
	r := gin.Default()
	r.Use(gin.Recovery()) // Set panic errors to be 500
	apis.RegisterAPIs(r, proxy)

	// Run the server
	if port == 0 {
		port = 3000
	}
	r.Run(fmt.Sprintf(":%d", port))
}
