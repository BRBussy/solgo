package tests

import (
	"context"
	solana "github.com/BRBussy/solgo"
	"github.com/BRBussy/solgo/systemProgram"
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

func (suite *JSONRPCConnectionTestSuite) TestGetAccountInfo() {
	getEncodedAccountInfoResponse, err := suite.jsonrpcConnection.GetAccountInfo(
		context.Background(),
		solana.GetAccountInfoRequest{
			PublicKey:       solana.NewPublicKeyFromBase58String("7ivguYMpnUBMboByJbKc7z31fJMg2pXYQ4nNPziWLchZ"),
			CommitmentLevel: solana.ProcessedCommitmentLevel,
		},
	)
	suite.Require().NotNil(getEncodedAccountInfoResponse)
	suite.Require().Nil(err)

	getJSONParsedAccountInfoResponse, err := suite.jsonrpcConnection.GetAccountInfo(
		context.Background(),
		solana.GetAccountInfoRequest{
			PublicKey:       solana.NewPublicKeyFromBase58String("DQLhiiGkoqRVtuBM8qczvrYdS29oWfnZcUzQJE16gZ2y"),
			CommitmentLevel: solana.ProcessedCommitmentLevel,
			Encoding:        solana.JSONParsedEncoding,
		},
	)
	suite.Require().NotNil(getJSONParsedAccountInfoResponse)
	suite.Require().Nil(err)
}

func (suite *JSONRPCConnectionTestSuite) TestGetBalance() {
	getBalanceResponse, err := suite.jsonrpcConnection.GetBalance(
		context.Background(),
		solana.GetBalanceRequest{
			PublicKey:       solana.NewPublicKeyFromBase58String("7ivguYMpnUBMboByJbKc7z31fJMg2pXYQ4nNPziWLchZ"),
			CommitmentLevel: solana.ProcessedCommitmentLevel,
		},
	)
	suite.Require().NotNil(getBalanceResponse)
	suite.Require().Nil(err)
}

func (suite *JSONRPCConnectionTestSuite) TestGetRecentBlockHash() {
	getRecentBlockHashResponse, err := suite.jsonrpcConnection.GetRecentBlockHash(
		context.Background(),
		solana.GetRecentBlockHashRequest{
			CommitmentLevel: solana.ProcessedCommitmentLevel,
		},
	)
	suite.Require().NotNil(getRecentBlockHashResponse)
	suite.Require().Nil(err)
}

func (suite *JSONRPCConnectionTestSuite) TestSendTransaction() {
	// create key pair for the account that will pay the opening Lamports for
	// the account that is going to be created
	fromKP := solana.MustNewRandomKeypair()

	// create key pair for the new account
	newAccKP := solana.MustNewRandomKeypair()

	// create a transaction
	tx := solana.NewTransaction()

	// build & add required operations
	createAccInstructions, err := systemProgram.CreateAccount(
		systemProgram.CreateAccountParams{
			FromPubkey:       fromKP.PublicKey,
			NewAccountPubkey: newAccKP.PublicKey,
			Lamports:         10000,
			Space:            0,
			ProgramID:        systemProgram.ID,
		},
	)
	suite.Require().Nilf(err, "error calling systemProgram.CreateAccount")
	suite.Require().Nil(
		tx.AddInstructions(createAccInstructions...),
		"error calling AddInstructions",
	)

	//base64, err := tx.ToBase58()
	//suite.Require().Nilf(err, "error marshalling transaction")

	//fmt.Println(base64)
}
