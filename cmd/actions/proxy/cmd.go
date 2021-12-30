package proxy

import (
	proxy "github.com/forbole/soljuno/actions-proxy"
	cmdtypes "github.com/forbole/soljuno/cmd/types"
	"github.com/spf13/cobra"
)

const (
	flagPort = "port"
)

func StartProxyCmd(cmdCfg *cmdtypes.Config) *cobra.Command {
	return &cobra.Command{
		Use:     "proxy-start",
		Short:   "Start actions proxy server for Hasura Actions",
		PreRunE: cmdtypes.ReadConfig(cmdCfg),
		RunE: func(cmd *cobra.Command, args []string) error {
			context, err := GetProxyContext(cmdCfg)
			if err != nil {
				return err
			}
			return StartProxyServer(context)
		},
	}
}

func StartProxyServer(context *Context) error {
	proxy.StartProxyServer(context.Proxy, context.Port)
	return nil
}
