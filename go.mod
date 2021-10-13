module github.com/forbole/soljuno

go 1.16

require (
	github.com/cosmos/cosmos-sdk v0.42.9
	github.com/go-co-op/gocron v1.9.0
	github.com/go-redis/redis v6.15.5+incompatible // indirect
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.4
	github.com/lib/pq v1.9.0
	github.com/mr-tron/base58 v1.2.0
	github.com/pelletier/go-toml v1.8.1
	github.com/prometheus/client_golang v1.11.0
	github.com/prometheus/common v0.30.0 // indirect
	github.com/prometheus/procfs v0.7.1 // indirect
	github.com/rs/zerolog v1.21.0
	github.com/spf13/cobra v1.1.3
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.11
	github.com/ybbus/jsonrpc/v2 v2.1.6
	golang.org/x/sys v0.0.0-20210630005230-0f9fa26af87c // indirect
	google.golang.org/protobuf v1.27.1 // indirect
	gopkg.in/yaml.v2 v2.4.0
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
