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

func (s *SmockerAPISuite) SetupSuite() {
	smockerHost, exist := os.LookupEnv("SMOCKER_HOST")
	s.Require().True(exist, "SMOCKER_HOST env variable have to be set")

	s.api = smocker.API{
		Host: smockerHost,
		Port: 8081,
	}

	s.url = fmt.Sprintf("%s:8080", smockerHost)
}

func (s *SmockerAPISuite) SetupTest() {
	err := s.api.Reset(context.Background())
	s.Require().NoError(err)
}

func callSetzer(params ...string) (string, error) {
	cmd := exec.Command("setzer", params...)
	out, err := cmd.Output()

	if werr, ok := err.(*exec.ExitError); ok {
		if s := werr.Error(); s != "0" {
			return "", fmt.Errorf("setzer exited with exit code: %s", s)
		}
	}

	return strings.TrimSpace(string(out)), nil
}
