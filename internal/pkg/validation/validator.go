package validation

// Validator is an interface for validation
type Validator interface {
	Validate(interface{}) error
}
