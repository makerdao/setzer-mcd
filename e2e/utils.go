package e2e

import (
	"context"
	"fmt"
	"github.com/chronicleprotocol/infestor/smocker"
	"github.com/stretchr/testify/suite"
	"os"
	"os/exec"
	"strings"
)

type SmockerAPISuite struct {
	suite.Suite
	api smocker.API
	url string
}

func (s *SmockerAPISuite) Setup() {
	smockerHost, exist := os.LookupEnv("SMOCKER_HOST")
	s.Require().True(exist, "SMOCKER_HOST env variable have to be set")

	s.api = smocker.API{
		Host: smockerHost,
		Port: 8081,
	}

	s.url = fmt.Sprintf("%s:8080", smockerHost)
}

func (s *SmockerAPISuite) Reset() {
	err := s.api.Reset(context.Background())
	s.Require().NoError(err)
}

func (s *SmockerAPISuite) SetupSuite() {
	s.Setup()
}

func (s *SmockerAPISuite) SetupTest() {
	s.Reset()
}

func callSetzer(params ...string) (string, string, error) {
	cmd := exec.Command("setzer", params...)
	cmd.Env = os.Environ()

	out, err := cmd.Output()

	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			return "", s, fmt.Errorf("setzer exited with exit code: %s", s)
		}
	}

	return strings.TrimSpace(string(out)), "0", nil
}
