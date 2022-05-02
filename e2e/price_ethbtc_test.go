package e2e

import (
	"net/http"
	"testing"

	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"
	"github.com/stretchr/testify/suite"
)

func TestPriceETHBTCE2ESuite(t *testing.T) {
	suite.Run(t, new(PriceETHBTCE2ESuite))
}

type PriceETHBTCE2ESuite struct {
	SmockerAPISuite
}

func (s *PriceETHBTCE2ESuite) TestPrice() {
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

	out, _, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *PriceETHBTCE2ESuite) TestPrice4Correct2Zero() {
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

	out, _, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *PriceETHBTCE2ESuite) TestPrice4Correct2Invalid() {
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

	out, _, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *PriceETHBTCE2ESuite) TestPrice3Correct3Invalid() {
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

	// SETZER_MIN_MEDIAN=4 and setzer should fail
	_, exitCode, err := callSetzer("price", "ethbtc")
	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}

func (s *PriceETHBTCE2ESuite) TestPriceMedianCalculationNotEnoughMinSources() {
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

	_, exitCode, err := callSetzer("price", "ethbtc")
	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}
