package controllers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type ControllersTestSuite struct {
	suite.Suite
	controllers Controllers
}

func TestSuiteControllers(t *testing.T) {
	suite.Run(t, new(ControllersTestSuite))
}

func (suite *ControllersTestSuite) SetupSuite() {
	fmt.Printf("Starting testing server")
}

func (suite *ControllersTestSuite) TearDownSuite() {
	fmt.Printf("Shutting down testing server")
}