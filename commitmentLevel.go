package solana

// CommitmentLevel is the level of commitment desired when querying state
type CommitmentLevel string

var (
	// ProcessedCommitmentLevel indicates to query the most recent
	// block which has reached 1 confirmation by the connected node
	ProcessedCommitmentLevel CommitmentLevel = "processed"

	// ConfirmedCommitmentLevel indicates to query the most recent
	// block which has reached 1 confirmation by the cluster
	ConfirmedCommitmentLevel CommitmentLevel = "confirmed"

	// FinalizedCommitmentLevel indicates to query the most recent
	// block which has been finalized by the cluster
	FinalizedCommitmentLevel CommitmentLevel = "finalized"

	RecentCommitmentLevel       CommitmentLevel = "recent"       // Deprecated as of v1.5.5
	SingleCommitmentLevel       CommitmentLevel = "single"       // Deprecated as of v1.5.5
	SingleGossipCommitmentLevel CommitmentLevel = "singleGossip" // Deprecated as of v1.5.5
	RootCommitmentLevel         CommitmentLevel = "root"         // Deprecated as of v1.5.5
	MaxCommitmentLevel          CommitmentLevel = "max"          // Deprecated as of v1.5.5
)
