package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forbole/soljuno/actions-proxy/apis/types"
	"github.com/forbole/soljuno/solana/client"
)

func RegisterAPIs(r *gin.Engine, proxy client.ClientProxy) {
	group := r.Group("/api")
	group.POST("/epoch_info", func(c *gin.Context) {
		info, err := proxy.GetEpochInfo()
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		c.JSON(http.StatusOK, info)
	})

	group.POST("/epoch_schedule", func(c *gin.Context) {
		schedule, err := proxy.GetEpochSchedule()
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		c.JSON(http.StatusOK, schedule)
	})

	group.POST("/inflation_rate", func(c *gin.Context) {
		inflation, err := proxy.GetInflationRate()
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		c.JSON(http.StatusOK, inflation)
	})

	group.POST("/inflation_governor", func(c *gin.Context) {
		governor, err := proxy.GetInflationGovernor()
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		c.JSON(http.StatusOK, governor)
	})

}
