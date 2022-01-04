package apis

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/forbole/soljuno/actions-proxy/apis/types"
	"github.com/forbole/soljuno/solana/client"
	clienttypes "github.com/forbole/soljuno/solana/client/types"
	"github.com/forbole/soljuno/solana/parser/manager"
	solanatypes "github.com/forbole/soljuno/types"
)

func RegisterAPIs(r *gin.Engine, proxy client.ClientProxy) {
	parserManager := manager.NewDefaultManager()
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

	group.POST("/tx_metas", func(c *gin.Context) {
		var playload types.TxMetaPayload
		if err := c.BindJSON(&playload); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		metas, err := proxy.GetSignaturesForAddress(
			playload.Input.Address,
			clienttypes.GetSignaturesForAddressConfig{
				Limit:  playload.Input.Limit,
				Before: playload.Input.Before,
				Until:  playload.Input.Until,
			},
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		c.JSON(http.StatusOK, metas)
	})

	group.POST("/tx", func(c *gin.Context) {
		var playload types.TxPayload
		if err := c.BindJSON(&playload); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		encodedTx, err := proxy.GetTransaction(playload.Input.Hash)
		if err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		tx := solanatypes.NewTxFromTxResult(
			parserManager,
			encodedTx.Slot,
			clienttypes.EncodedTransactionWithStatusMeta{
				Transaction: encodedTx.Transaction,
				Meta:        encodedTx.Meta,
			},
		)
		c.JSON(http.StatusOK, types.NewTxResponse(tx))
	})
}
