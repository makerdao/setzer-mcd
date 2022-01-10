package e2e

import (
	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

func TestPriceETHBTCSuite(t *testing.T) {
	suite.Run(t, new(PriceETHBTCSuite))
}

type PriceETHBTCSuite struct {
	SmockerAPISuite
}

func (s *PriceETHBTCSuite) TestPrice() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithPrice(1)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *PriceETHBTCSuite) TestPrice4Correct2Zero() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithPrice(0)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithPrice(0)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *PriceETHBTCSuite) TestPrice4Correct2Invalid() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithStatusCode(http.StatusNotFound)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *PriceETHBTCSuite) TestPrice3Correct3Invalid() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithStatusCode(http.StatusNotFound)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("0.5000000000", out)
}

func (s *PriceETHBTCSuite) TestPrice1Correct5Invalid() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithStatusCode(http.StatusNotFound)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("0.0000000000", out)
}

func (s *PriceETHBTCSuite) TestPriceMedianCalculation() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithStatusCode(http.StatusNotFound)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("0.0000000000", out)
}
