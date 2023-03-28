package validator

import "strings"

// Result Validation result
type Result struct {
	errors []error
}

func (r *Result) addError(err error) {
	r.errors = append(r.errors, err)
}

// IsSuccess returns true when no error has been added to result
func (r Result) IsSuccess() bool {
	return r.errors == nil ||
		len(r.errors) == 0
}

// IsFailure returns true when any error has been added to result
func (r Result) IsFailure() bool {
	return !r.IsSuccess()
}

func (r Result) Errors() []string {
	result := make([]string, len(r.errors))
	for i, err := range r.errors {
		result[i] = err.Error()
	}
	return result
}

func (r Result) Error() string {
	return strings.Join(r.Errors(), ",")
}
