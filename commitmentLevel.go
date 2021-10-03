package solana

type CommitmentLevel string

var (
	ProcessedCommitmentLevel    = "processed"
	ConfirmedCommitmentLevel    = "confirmed"
	FinalizedCommitmentLevel    = "finalized"
	RecentCommitmentLevel       = "recent"
	SingleCommitmentLevel       = "single"
	SingleGossipCommitmentLevel = "singleGossip"
	RootCommitmentLevel         = "root"
	MaxCommitmentLevel          = "max"
)
