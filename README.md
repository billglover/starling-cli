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

All required configuration options can be specified on the command line. You may find it useful to create a configuration file in your home directory, e.g. `~/.starling.yaml`. An example file is shown below.

```yaml
# the default operating environment
env: live

# the user access token
token: zsbrdTn4GsXadKAuXdKRSJjww0CcmDCVmoNufV2ubqViZ4ynKUuF88uTspLYoCRu

# the default number of items displayed when requesting a list
limit: 5

# the default currency
currency: GBP
```

## Usage

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
