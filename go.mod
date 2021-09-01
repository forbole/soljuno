module github.com/forbole/soljuno

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/desmos-labs/juno v0.0.0-20210812114304-074d555149ed
	github.com/go-co-op/gocron v0.3.3
	github.com/gorilla/mux v1.8.0
	github.com/lib/pq v1.9.0
	github.com/mr-tron/base58 v1.2.0
	github.com/pelletier/go-toml v1.8.1
	github.com/prometheus/client_golang v1.11.0
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/ybbus/jsonrpc/v2 v2.1.6
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
