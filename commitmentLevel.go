package solana

type CommitmentLevel string

var (
	ProcessedCommitmentLevel    CommitmentLevel = "processed"
	ConfirmedCommitmentLevel    CommitmentLevel = "confirmed"
	FinalizedCommitmentLevel    CommitmentLevel = "finalized"
	RecentCommitmentLevel       CommitmentLevel = "recent"
	SingleCommitmentLevel       CommitmentLevel = "single"
	SingleGossipCommitmentLevel CommitmentLevel = "singleGossip"
	RootCommitmentLevel         CommitmentLevel = "root"
	MaxCommitmentLevel          CommitmentLevel = "max"
)
