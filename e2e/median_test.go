package e2e

import (
	"net/http"
	"os"
	"testing"

	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"
	"github.com/stretchr/testify/suite"
)

func TestMedianE2ESuite(t *testing.T) {
	suite.Run(t, new(MedianE2ESuite))
}

type MedianE2ESuite struct {
	SmockerAPISuite
	minMedian   string
	cacheExpiry string
}

func (s *MedianE2ESuite) SetupSuite() {
	s.Setup()
	s.minMedian = os.Getenv("SETZER_MIN_MEDIAN")
	s.cacheExpiry = os.Getenv("SETZER_CACHE_EXPIRY")
}

func (s *MedianE2ESuite) TearDownTest() {
	err := os.Setenv("SETZER_MIN_MEDIAN", s.minMedian)
	s.Require().NoError(err)

	err = os.Setenv("SETZER_CACHE_EXPIRY", s.cacheExpiry)
	s.Require().NoError(err)
}

func (s *MedianE2ESuite) TestBaseMedian() {
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

func (s *MedianE2ESuite) Test50x50Median() {
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

func (s *MedianE2ESuite) TestMedianWithInvalidSources() {
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

	err = os.Setenv("SETZER_MIN_MEDIAN", "3")
	s.Require().NoError(err)

	out, _, err := callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *MedianE2ESuite) TestMedianWithMoreInvalidSourcesThanRequired() {
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
	s.Require().Equal(1, exitCode)
}

func (s *MedianE2ESuite) TestCacheShouldReturnValidValue() {
	// Setting up cache expiration timeout
	err := os.Setenv("SETZER_CACHE_EXPIRY", "10")
	s.Require().NoError(err)

	err = infestor.NewMocksBuilder().
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

	// Should ignore API calls and just get median from cache
	err = infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithStatusCode(http.StatusNotFound)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithStatusCode(http.StatusNotFound)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, _, err = callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *MedianE2ESuite) TestDisabledCacheValidValue() {
	// Setting up cache expiration timeout
	err := os.Setenv("SETZER_CACHE_EXPIRY", "-1")
	s.Require().NoError(err)

	err = infestor.NewMocksBuilder().
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

	// Should ignore API calls and just get median from cache
	err = infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("bitfinex").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("coinbase").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("huobi").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("poloniex").WithSymbol("ETH/BTC").WithPrice(2)).
		Add(origin.NewExchange("kraken").WithSymbol("XETH/XXBT").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, _, err = callSetzer("price", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)
}
