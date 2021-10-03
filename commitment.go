package solana

type Commitment string

var (
	ProcessedCommitment    = "processed"
	ConfirmedCommitment    = "confirmed"
	FinalizedCommitment    = "finalized"
	RecentCommitment       = "recent"
	SingleCommitment       = "single"
	SingleGossipCommitment = "singleGossip"
	RootCommitment         = "root"
	MaxCommitment          = "max"
)
