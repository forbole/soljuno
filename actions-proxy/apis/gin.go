package apis

import (
	"fmt"
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

	group.POST("/tx_meta", func(c *gin.Context) {
		var playload types.TxByAddressPayload
		if err := c.BindJSON(&playload); err != nil {
			c.JSON(http.StatusBadRequest, err)
			return
		}
		metas, err := proxy.GetSignaturesForAddress(
			playload.Input.Address,
			clienttypes.GetSignaturesForAddressConfig{
				Limit:  playload.Input.Config.Limit,
				Before: playload.Input.Config.Before,
				Until:  playload.Input.Config.Until,
			},
		)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		c.JSON(http.StatusOK, types.NewTxMetasResponse(metas))
	})

	group.POST("/tx", func(c *gin.Context) {
		var playload types.TxPayload
		if err := c.BindJSON(&playload); err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		encodedTx, err := proxy.GetTransaction(playload.Input.Hash)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
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

	group.POST("/account_info", func(c *gin.Context) {
		var playload types.AccountInfoPayload
		if err := c.BindJSON(&playload); err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		addr := playload.Input.Address
		info, err := proxy.GetAccountInfo(addr)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		if info.Value == nil {
			c.JSON(http.StatusBadRequest, types.NewError(fmt.Errorf("%s does not exist", addr)))
			return
		}
		res, err := types.NewAccountInfoResponse(info)
		if err != nil {
			c.JSON(http.StatusBadRequest, types.NewError(err))
			return
		}
		c.JSON(http.StatusOK, res)
	})
}
