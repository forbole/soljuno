module github.com/forbole/soljuno

go 1.16

require (
	github.com/gin-gonic/gin v1.7.7
	github.com/go-co-op/gocron v1.9.0
	github.com/gorilla/mux v1.8.0
	github.com/jmoiron/sqlx v1.3.5
	github.com/lib/pq v1.9.0
	github.com/magiconair/properties v1.8.5 // indirect
	github.com/mitchellh/mapstructure v1.3.3 // indirect
	github.com/mkevac/debugcharts v0.0.0-20191222103121-ae1c48aa8615
	github.com/mr-tron/base58 v1.2.0
	github.com/panjf2000/ants/v2 v2.4.6
	github.com/pelletier/go-toml v1.8.1
	github.com/prometheus/client_golang v1.13.0
	github.com/rs/zerolog v1.21.0
	github.com/spf13/afero v1.3.4 // indirect
	github.com/spf13/cast v1.3.1 // indirect
	github.com/spf13/cobra v1.1.3
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/viper v1.7.1
	github.com/stretchr/testify v1.8.0
	github.com/tendermint/tendermint v0.34.11
	github.com/ybbus/jsonrpc/v2 v2.1.6
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad // indirect
	gopkg.in/ini.v1 v1.61.0 // indirect
	gopkg.in/yaml.v2 v2.4.0
	gopkg.in/yaml.v3 v3.0.1
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
