package types

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDefaultConfigParser(t *testing.T) {
	data := `
[chain]
  modules = [
    "pruning"
  ]

[rpc]
  client_name = "soljuno"
  address = "http://localhost:8899"

[logging]
  format = "text"
  level = "debug"

[parsing]
  workers = 5
  listen_new_blocks = true
  parse_old_blocks = true
  start_height = 1

[database]
  host = "localhost"
  name = "juno"
  password = "password"
  port = 5432
  schema = "public"
  ssl_mode = ""
  user = "user"

[pruning]
  keep_recent = 100
  keep_every = 5
  interval = 10
`

	cfg, err := DefaultConfigParser([]byte(data))
	require.NoError(t, err)

	require.Equal(t, []string{"pruning"}, cfg.GetChainConfig().GetModules())

	require.Equal(t, "soljuno", cfg.GetRPCConfig().GetClientName())
	require.Equal(t, "http://localhost:8899", cfg.GetRPCConfig().GetAddress())

}
