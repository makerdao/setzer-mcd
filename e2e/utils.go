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
	out, err := exec.Command("setzer", params...).Output()
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(out)), nil
}
