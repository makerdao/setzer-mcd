# Setzer MCD

Query USD price feeds

## Usage

```
Usage: setzer <command> [<args>]
   or: setzer <command> --help

Commands:

   help            Print help about setzer or one of its subcommands
   pairs           List all supported pairs
   price           Show price(s) for a given asset or pair
   sources         Show price sources for a given asset or pair
   test            Test all price feeds
```


## Installation

Dependencies:

* GNU [bc](https://www.gnu.org/software/bc/)
* [curl](https://curl.haxx.se/download.html)
* GNU [datamash](https://www.gnu.org/software/datamash/)
* GNU `date`
* [jshon](http://kmkeen.com/jshon/)
* GNU `timeout`

Install via make:

* `make link` -  link setzer into `/usr/local`
* `make install` -  copy setzer into `/usr/local`
* `make uninstall` -  remove setzer from `/usr/local`

## Configuration

* `SETZER_CACHE` - Cache directory (default: ~/.setzer)
* `SETZER_CACHE_EXPIRY` - Cache expiry (default: 60) seconds
* `SETZER_TIMEOUT` - HTTP request timeout (default: 10) seconds

## wstETH pair requirement

Due to process of pulling details from mainnet for getting price information.
You need to set `ETH_RPC_URL` environemnt variable. By default it will point to `http://127.0.0.1:8545`.

Example of usage: 

```bash
export ETH_RPC_URL="https://mainnet.infura.io/v3/fac98e56ea7e49608825dfc726fab703"
```