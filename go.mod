module github.com/forbole/soljuno

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-co-op/gocron v1.9.0
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.10.6
	github.com/mkevac/debugcharts v0.0.0-20191222103121-ae1c48aa8615
	github.com/mr-tron/base58 v1.2.0
	github.com/panjf2000/ants/v2 v2.4.6
	github.com/pelletier/go-toml v1.9.5
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/zerolog v1.28.0
	github.com/spf13/cobra v1.5.0
	github.com/spf13/viper v1.13.0
	github.com/stretchr/testify v1.8.0
	github.com/tendermint/tendermint v0.35.9
	github.com/ybbus/jsonrpc/v2 v2.1.7
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
