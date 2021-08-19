package solana

import "fmt"

type Network string

const (
	MainnetBeta  Network = "Mainnet Beta"
	Testnet      Network = "Testnet"
	Devnet       Network = "Devnet"
	LocalTestnet Network = "LocalTestnet"
)

func (n Network) String() string {
	return string(n)
}

// ToRPCURL returns the the rpc url of the relevant public
// Solana foundation nodes for MainnetBeta, Testnet and Devnet.
func (n Network) ToRPCURL() (string, error) {
	switch n {
	case MainnetBeta:
		return "https://api.mainnet-beta.solana.com", nil

	case Testnet:
		return "https://api.testnet.solana.com", nil

	case Devnet:
		return "https://api.devnet.solana.com", nil

	case LocalTestnet:
		return "http://localhost:8899", nil
	}

	return "", fmt.Errorf("%s: %w", n, ErrUnexpectedNetwork)
}
