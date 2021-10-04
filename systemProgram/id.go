// Package systemProgram provides a set of functions for constructing Solana system program
// instructions.
// See instruction definitions here:
// https://github.com/solana-labs/solana/blob/4b2fe9b20d4c895f4d3cb58c2918c72a5b0a5b64/sdk/program/src/system_instruction.rs#L142
package systemProgram

import solana "github.com/BRBussy/solgo"

// ID is the Solana system program ID
var ID = solana.NewPublicKeyFromBase58String("11111111111111111111111111111111")
