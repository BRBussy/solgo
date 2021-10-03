package tests

import (
	"context"
	solana "github.com/BRBussy/solgo"
	"github.com/stretchr/testify/suite"
	"testing"
)

func TestJSONRPCConnectionTestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping JSONRPCConnectionTestSuite in short mode.")
	}
	suite.Run(t, new(JSONRPCConnectionTestSuite))
}

type JSONRPCConnectionTestSuite struct {
	suite.Suite
	jsonrpcConnection *solana.JSONRPCConnection
}

// SetupSuite will run before the suite
func (suite *JSONRPCConnectionTestSuite) SetupSuite() {
	suite.jsonrpcConnection = solana.NewJSONRPCConnection()
}

func (suite *JSONRPCConnectionTestSuite) TestGetBalance() {
	_, err := suite.jsonrpcConnection.GetBalance(
		context.Background(),
		solana.GetBalanceRequest{
			PublicKey: nil,
		},
	)
	suite.Require().Nil(err)
}
