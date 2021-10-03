package solana

import (
	"context"
)

type JSONRPCConnectionCommitmentConfig struct {
	commitment CommitmentLevel
}

type JSONRPCConnection struct {
	commitmentConfig JSONRPCConnectionCommitmentConfig
}

func (J *JSONRPCConnection) Commitment(ctx context.Context) (CommitmentLevel, error) {
	panic("implement me")
}

func (J *JSONRPCConnection) GetBalance(ctx context.Context, request GetBalanceRequest) (*GetBalanceResponse, error) {
	panic("implement me")
}
