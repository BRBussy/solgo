package solana

// Context is extra contextual information for RPC responses
type Context struct {
	Slot int `json:"slot"`
}
