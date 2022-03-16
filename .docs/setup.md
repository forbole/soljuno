# Setup
Setting up Soljuno is pretty straightforward. It requires three things to be done:
1. Install Soljuno.
1. Initialize the configuration.
2. Start the parser.

## Installing Juno
In order to install Soljuno you are required to have [Go 1.15+](https://golang.org/dl/) installed on your machine. Once you have it, the first thing to do is to clone the GitHub repository. To do this you can run

```shell
$ git clone https://github.com/forbole/soljuno.git
```

Then, you need to install the binary. To do this, run

```shell
$ make install
```

This will put the `soljuno` binary inside your `$GOPATH/bin` folder. You should now be able to run `soljuno` to make sure it's installed:

```shell
$ soljuno
A Solana chain data aggregator. It improves the chain's data accessibility
by providing an indexed database exposing aggregated resources and models such as blocks, validators, transactions, and etc. 
soljuno is meant to run with a GraphQL layer on top so that it even further eases the ability for developers and
downstream clients to answer queries such as "What is the average gas cost of a block?" while also allowing
them to compose more aggregate and complex queries.

Usage:
  soljuno [command]

Available Commands:
  db               Database subcommands
  help             Help about any command
  import-snapshot  Import a snapshot at specific slot
  import-tokenlist Import a tokenlist to the token unit table
  init             Initializes the configuration files
  parse            Start parsing the blockchain data
  proxy-start      Start actions proxy server for Hasura Actions
  version          Print the version information

Flags:
  -h, --help          help for soljuno
      --home string   Set the home folder of the application, where all files will be stored (default "/home/kilem/.soljuno")

Use "soljuno [command] --help" for more information about a command.
```

## Initializing the configuration
In order to correctly parse and store the data based on your requirements, Soljuno allows you to customize its behavior via a TOML file called `config.toml`. In order to create the first instance of the `config.toml` file you can run

```shell
$ soljuno init
```

This will create such file inside the `~/.soljuno` folder.  
Note that if you want to change the folder used by Soljuno you can do this using the `--home` flag:

```shell
$ soljuno init --home /path/to/my/folder
```

Once the file is created, you are required to edit it and change the different values. To do this you can run

```shell
$ nano ~/.soljuno/config.toml
```

For a better understanding of what each section and field refers to, please read the [config reference](config.md).

## Running Soljuno
Once the configuration file has been setup, you can run Soljuno using the following command:

```shell
$ soljuno parse
```

If you are using a custom folder for the configuration file, please specify it using the `--home` flag:


```shell
$ soljuno parse --home /path/to/my/config/folder
```

We highly suggest you running Soljuno as a system service so that it can be restarted automatically in the case it stops. To do this you can run:

```shell
$ sudo tee /etc/systemd/system/soljuno.service > /dev/null <<EOF
[Unit]
Description=SolJuno parser
After=network-online.target

[Service]
User=$USER
ExecStart=$GOPATH/bin/soljuno parse
Restart=always
RestartSec=3
LimitNOFILE=4096

[Install]
WantedBy=multi-user.target
EOF
```

Then you need to enable and start the service:

```shell
$ sudo systemctl enable soljuno
$ sudo systemctl start soljuno
```