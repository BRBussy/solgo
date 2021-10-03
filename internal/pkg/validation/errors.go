package validation

type ErrValidationFailed struct {
	Reason string
}

func (e ErrValidationFailed) Error() string {
	return "validation failed: " + e.Reason
}
