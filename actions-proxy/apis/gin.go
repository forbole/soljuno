package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forbole/soljuno/actions-proxy/apis/types"
	"github.com/forbole/soljuno/solana/client"
)

func RegisterAPIs(r *gin.Engine, proxy client.ClientProxy) {
	r.Group("/api").
		GET("/epoch_info", func(c *gin.Context) {
			epochInfo, err := proxy.GetEpochInfo()
			if err != nil {
				c.JSON(http.StatusBadRequest, types.NewError(err))
				return
			}
			c.JSON(http.StatusOK, types.NewEpochInfo(epochInfo))
		})

}
