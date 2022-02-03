package e2e

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

func TestBaseBehaviourE2ESuite(t *testing.T) {
	suite.Run(t, new(BaseBehaviourE2ESuite))
}

type BaseBehaviourE2ESuite struct {
	SmockerAPISuite
}

func (s *BaseBehaviourE2ESuite) TestVersionCommand() {
	out, _, err := callSetzer("version")

	s.Require().NoError(err)
	s.Require().Contains(out, "setzer-mcd")
}

func (s *BaseBehaviourE2ESuite) TestFormatCommand() {
	out, _, err := callSetzer("--format")
	s.Require().NoError(err)
	s.Require().Equal("0.0000000000", out)

	out, _, err = callSetzer("--format", "1")
	s.Require().NoError(err)
	s.Require().Equal("1.0000000000", out)

	out, _, err = callSetzer("--format", "1.5")
	s.Require().NoError(err)
	s.Require().Equal("1.5000000000", out)

	out, _, err = callSetzer("--format", "1.5000000005")
	s.Require().NoError(err)
	s.Require().Equal("1.5000000005", out)

	out, _, err = callSetzer("--format", "1.50000000056")
	s.Require().NoError(err)
	s.Require().Equal("1.5000000006", out)

	// not number should fail
	_, exitCode, err := callSetzer("--format", "abc")
	s.Require().Error(err)
	s.Require().Equal(1, exitCode)
}

func (s *BaseBehaviourE2ESuite) TestPairsCommand() {
	out, _, err := callSetzer("pairs")
	s.Require().NoError(err)
	s.Require().Contains(out, "BTCUSD\n")
}

func (s *BaseBehaviourE2ESuite) TestHelpCommand() {
	out, _, err := callSetzer("help")
	s.Require().NoError(err)
	s.Require().Contains(out, "Usage: setzer")
}

func (s *BaseBehaviourE2ESuite) TestSourcesCommand() {
	// No pair arg for `sources` command
	_, exitCode, err := callSetzer("sources")
	s.Require().Error(err)
	s.Require().Equal(1, exitCode)

	out, _, err := callSetzer("sources", "ETHBTC")
	s.Require().NoError(err)
	s.Require().Contains(out, "binance\n")
}
