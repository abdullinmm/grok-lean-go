package validator

// Validator is an interface that defines the Validate method.
type Validator interface {
	Validate() error
}
