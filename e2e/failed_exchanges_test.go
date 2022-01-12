package e2e

import (
	"net/http"
	"testing"

	"github.com/chronicleprotocol/infestor"
	"github.com/chronicleprotocol/infestor/origin"
	"github.com/stretchr/testify/suite"
)

func TestFailedExchangesE2ESuite(t *testing.T) {
	suite.Run(t, new(FailedExchangesE2ESuite))
}

type FailedExchangesE2ESuite struct {
	SmockerAPISuite
}

func (s *FailedExchangesE2ESuite) TestPriceWithoutArgs() {
	_, exitCode, err := callSetzer("price")

	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}

func (s *FailedExchangesE2ESuite) TestUnknownPairFailes() {
	_, exitCode, err := callSetzer("price", "balancer", "foo:bar")

	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}

func (s *FailedExchangesE2ESuite) TestFailedBalancer() {
	ex := origin.
		NewExchange("balancer").
		WithSymbol("BAL/USD").
		WithCustom("contract", "0xba100000625a3754423978a60c9317c58a424e3d").
		WithStatusCode(http.StatusBadRequest)

	err := infestor.NewMocksBuilder().Reset().Add(ex).Deploy(s.api)
	s.Require().NoError(err)

	_, exitCode, err := callSetzer("price", "balancer", "balusd")
	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}

func (s *FailedExchangesE2ESuite) TestXPriceUnknownExchange() {
	_, exitCode, err := callSetzer("price", "foobar", "ethbtc")
	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}
