package solana

type FeeCalculator struct {
	feeScheduleBlockHash string
	lamportsPerSignature int64
}

func (f *FeeCalculator) LamportsPerSignature() int64 {
	return f.lamportsPerSignature
}

func (f *FeeCalculator) FeeScheduleBlockHash() string {
	return f.feeScheduleBlockHash
}

func (f *FeeCalculator) CalculateTransactionFee(transaction Transaction) int64 {
	return f.lamportsPerSignature * int64(len(transaction.signatures))
}
