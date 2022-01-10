package e2e

import (
	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"os"
	"testing"
)

// SETZER_MIN_MEDIAN

func TestMedianSuite(t *testing.T) {
	suite.Run(t, new(MedianSuite))
}

type MedianSuite struct {
	SmockerAPISuite
	minMedian string
}

func (s *MedianSuite) SetupSuite() {
	s.Setup()
	s.minMedian = os.Getenv("SETZER_MIN_MEDIAN")
}

func (s *MedianSuite) TearDownTest() {
	//err := os.Setenv("SETZER_MIN_MEDIAN", s.minMedian)
	//s.Require().NoError(err)
}

func (s *MedianSuite) TestBaseMedian() {
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

func (s *MedianSuite) Test50x50Median() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, _, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.5000000000", out)
}

func (s *MedianSuite) TestMinMedian() {
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

	err = os.Setenv("SETZER_MIN_MEDIAN", "4")
	s.Require().NoError(err)

	_, exitCode, err := callSetzer("price", "ethbtc")
	s.Require().Error(err)
	s.Require().Equal("1", exitCode)
}
