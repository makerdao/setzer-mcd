package e2e

import (
	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"
	"github.com/stretchr/testify/suite"
	"net/http"
	"testing"
)

func TestExchangesE2ESuite(t *testing.T) {
	suite.Run(t, new(SetzerExchangesE2ETest))
}

type SetzerExchangesE2ETest struct {
	SmockerAPISuite
}

func (s *SetzerExchangesE2ETest) TestBalancer() {
	ex := origin.
		NewExchange("balancer").
		WithSymbol("BAL/USD").
		WithPrice(1).
		WithCustom("contract", "0xba100000625a3754423978a60c9317c58a424e3d")

	err := infestor.NewMocksBuilder().Reset().Add(ex).Deploy(s.api)
	s.Require().NoError(err)

	out, err := callSetzer("x-price", "balancer", "balusd")

	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	ex = origin.
		NewExchange("balancer").
		WithSymbol("BAL/USD").
		WithPrice(1).
		WithCustom("contract", "0xba100000625a3754423978a60c9317c58a424e3d").
		WithStatusCode(http.StatusNotFound)

	err = infestor.NewMocksBuilder().Reset().Add(ex).Deploy(s.api)
	s.Require().NoError(err)

	out, err = callSetzer("x-price", "balancer", "balusd")

	s.Require().NoError(err)
	s.Require().Equal("0.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestBinance() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("binance").WithSymbol("ETH/BTC").WithPrice(1)).
		Add(origin.NewExchange("binance").WithSymbol("AAVE/BTC").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "binance", "ethbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "binance", "aave:btc")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestBitfinex() {
	// NOTE: For symbols of 4 chars you have to write `SYMBOL:` otherwise API request to smocker will fail.
	// Example: AVAX/USD should be written in mock as `AVAX:/USD`
	err := infestor.NewMocksBuilder().
		Reset().
		// Add(origin.NewExchange("bitfinex").WithSymbol("AVAX:/USD").WithPrice(1)).
		Add(origin.NewExchange("bitfinex").WithSymbol("MKR/ETH").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	// out, err := callSetzer("x-price", "bitfinex", "avax%3A:usd")
	// s.Require().NoError(err)
	// s.Require().Equal("1.0000000000", out)

	out, err := callSetzer("x-price", "bitfinex", "mkr:eth")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestBitstamp() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("bitstamp").WithSymbol("ETH/USD").WithPrice(1)).
		Add(origin.NewExchange("bitstamp").WithSymbol("BTC/USD").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "bitstamp", "ethusd")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "bitstamp", "btcusd")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestBithumb() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("bitthumb").WithSymbol("PAXG/USDT").WithPrice(1)).
		Add(origin.NewExchange("bitthumb").WithSymbol("SOL/BTC").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "bitthumb", "paxg:usdt")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "bitthumb", "sol:btc")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestBittrex() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("bittrex").WithSymbol("BAT/BTC").WithPrice(1)).
		Add(origin.NewExchange("bittrex").WithSymbol("BTC/USD").WithPrice(2)).
		Add(origin.NewExchange("bittrex").WithSymbol("BTC/GNT").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "bittrex", "batbtc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "bittrex", "btcusd")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "bittrex", "btcgnt")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestCoinbase() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("coinbase").WithSymbol("BAL/USD").WithPrice(1)).
		Add(origin.NewExchange("coinbase").WithSymbol("BTC/USD").WithPrice(2)).
		Add(origin.NewExchange("coinbase").WithSymbol("COMP/USD").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "coinbase", "bal:usd")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "coinbase", "btcusd")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "coinbase", "comp:usd")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestCryptocompare() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("cryptocompare").WithSymbol("POLY/USD").WithPrice(1)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "cryptocompare", "poly:usd")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestFTX() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("ftx").WithSymbol("ETH/USD").WithPrice(1)).
		Add(origin.NewExchange("ftx").WithSymbol("LINK/USD").WithPrice(2)).
		Add(origin.NewExchange("ftx").WithSymbol("MATIC/USD").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "ftx", "ethusd")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "ftx", "link:usd")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "ftx", "matic:usd")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestGateIO() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("gateio").WithSymbol("AVAX/USDT").WithPrice(1)).
		Add(origin.NewExchange("gateio").WithSymbol("LRC/USDT").WithPrice(2)).
		Add(origin.NewExchange("gateio").WithSymbol("SOL/USDT").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "gateio", "avax:usdt")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "gateio", "lrc:usdt")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "gateio", "sol:usdt")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestGemini() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("gemini").WithSymbol("AAVE/USD").WithPrice(1)).
		Add(origin.NewExchange("gemini").WithSymbol("BTC/USD").WithPrice(2)).
		Add(origin.NewExchange("gemini").WithSymbol("MATIC/USD").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "gemini", "aave:usd")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "gemini", "btcusd")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "gemini", "matic:usd")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestHitBTC() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("hitbtc").WithSymbol("MKR/BTC").WithPrice(1)).
		Add(origin.NewExchange("hitbtc").WithSymbol("XRP/BTC").WithPrice(2)).
		Add(origin.NewExchange("hitbtc").WithSymbol("XTZ/USD").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "hitbtc", "mkr:btc")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "hitbtc", "xrp:btc")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "hitbtc", "xtz:usd")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestHuobi() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("huobi").WithSymbol("AAVE/USDT").WithPrice(1)).
		Add(origin.NewExchange("huobi").WithSymbol("BAL/USDT").WithPrice(2)).
		Add(origin.NewExchange("huobi").WithSymbol("DOT/USDT").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "huobi", "aave:usdt")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "huobi", "bal:usdt")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "huobi", "dot:usdt")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestKraken() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("kraken").WithSymbol("XXBT/ZUSD").WithPrice(1)).
		Add(origin.NewExchange("kraken").WithSymbol("COMP/USD").WithPrice(2)).
		Add(origin.NewExchange("kraken").WithSymbol("DOT/USD").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "kraken", "xxbt:zusd")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "kraken", "comp:usd")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "kraken", "dot:usd")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestKukoin() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("kucoin").WithSymbol("COMP/USDT").WithPrice(2)).
		Add(origin.NewExchange("kucoin").WithSymbol("DOT/USDT").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "kucoin", "comp:usdt")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "kucoin", "dot:usdt")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestKyber() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("kyber").WithSymbol("DGX/ETH").WithPrice(1)).
		Add(origin.NewExchange("kyber").WithSymbol("KNC/ETH").WithPrice(2)).
		Add(origin.NewExchange("kyber").WithSymbol("MKR/ETH").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "kyber", "dgx:eth")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "kyber", "knc:eth")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "kyber", "mkr:eth")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestOkex() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("okex").WithSymbol("AAVE/USDT").WithPrice(1)).
		Add(origin.NewExchange("okex").WithSymbol("BAL/USDT").WithPrice(2)).
		Add(origin.NewExchange("okex").WithSymbol("MKR/BTC").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "okex", "aave:usdt")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "okex", "bal:usdt")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "okex", "mkr:btc")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestPoloniex() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("poloniex").WithSymbol("AAVE/USDT").WithPrice(1)).
		Add(origin.NewExchange("poloniex").WithSymbol("BAL/USDT").WithPrice(2)).
		Add(origin.NewExchange("poloniex").WithSymbol("MKR/BTC").WithPrice(3)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "poloniex", "aave:usdt")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "poloniex", "bal:usdt")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)

	out, err = callSetzer("x-price", "poloniex", "mkr:btc")
	s.Require().NoError(err)
	s.Require().Equal("3.0000000000", out)
}

func (s *SetzerExchangesE2ETest) TestUpbit() {
	err := infestor.NewMocksBuilder().
		Reset().
		Add(origin.NewExchange("upbit").WithSymbol("BAT/KRW").WithPrice(1)).
		Add(origin.NewExchange("upbit").WithSymbol("MANA/KRW").WithPrice(2)).
		Deploy(s.api)

	s.Require().NoError(err)

	out, err := callSetzer("x-price", "upbit", "batkrw")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, err = callSetzer("x-price", "upbit", "mana:krw")
	s.Require().NoError(err)
	s.Require().Equal("2.0000000000", out)
}
