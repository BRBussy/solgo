package solana

// Context is extra contextual information for RPC responses
type Context struct {
	Slot uint64 `json:"slot"`
}
