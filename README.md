# starling-cli

[![Build Status](https://travis-ci.com/billglover/starling-cli.svg?branch=master)](https://travis-ci.com/billglover/starling-cli)

A command line client for Starling Bank

## Installation

### Homebrew

```plain
brew install billglover/tap/starling-cli
```

### Go

```plain
go get -u github.com/billglover/starling-cli
```

## Configuration

All required configuration options can be specified on the command line bu you may find it useful to create a configuration file in your home directory, e.g. `~/.starling.yaml`. An example file is shown below.

```yaml
# the default operating environment
env: live

# your personal access token
token: zsbrdTn4GsXadKAuXdKRSJjww0CcmDCVmoNufV2ubqViZ4ynKUuF88uTspLYoCRu

# the default number of items displayed when requesting a list
limit: 5

# the default currency
currency: GBP
```

Values provided on the command line will override values provided in the configuration file.

## Usage

You will need your Starling Bank personal access token to authenticate your requests. This can be found on the [Starling Developer portal](https://developer.starlingbank.com) under the [Personal Access](https://developer.starlingbank.com/personal/list) tab. Remember this token grants access to your bank account and should be protected accordingly.

Listing transactions:

```plain
starling-cli list txns
#   Created                  Amount     Narrative
000 2018-06-09T14:43:15.990Z       8.00 demo 4
001 2018-06-09T14:42:45.214Z     -10.00 demo 4
002 2018-06-09T08:24:26.442Z     -32.26 External Payment
003 2018-06-09T08:24:24.899Z     -22.63 External Payment
004 2018-06-09T08:24:24.835Z     -18.56 Mastercard
5 of 100 transactions
```

Displaying your balance:

```plain
starling-cli show balance
Amount:       15807.02
Available:    15807.02
Cleared:      15807.02
Overdraft:        0.00
Pending:          0.00
Effective:    15807.02
Currency:          GBP
```
Getting help:

```plain
$ starling-cli
This is a basic command line interface for personal Starling
Bank accounts. It allows you to perform basic banking from
the command line. For example:

        starling-cli list txns

The Starling API is still under active development and until it
stabilises there may be some instability in the functionality
provided

Usage:
  starling-cli [command]

Available Commands:
  create      Create things like goals, payees, etc.
  delete      Delete things like goals, payees, etc.
  help        Help about any command
  list        Display a list of items based on sub-command
  show        Display a table of information based on sub-command
  transfer    Transfer money to/from a savings goal

Flags:
      --config string     config file (default is $HOME/.starling.yaml)
      --currency string   the currency you are using (default "GBP")
      --env string        the environment you want to use: live, sandbox (default "sandbox")
  -h, --help              help for starling-cli
      --token string      API access token
      --uuid              display UUID for objects

Use "starling-cli [command] --help" for more information about a command.
```
