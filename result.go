package validator

import (
	"errors"
	"strings"
)

// Result represent the result of a validation process
type Result struct {
	failures []error
}

// Error implements the error interface. Return all the error messages in the Result joined by semicolon (;)
func (e Result) Error() string {
	if e.IsSuccess() {
		return ""
	}

	return strings.Join(e.GetFailureMessages(), ";")
}

// AddFailureMessage adds a validation failure message to the Result
// msg is the message to be added
// if the message to be added is empty or whitespaces, nothing is added
func (e *Result) AddFailureMessage(msg string) {
	if strings.TrimSpace(msg) == "" {
		return
	}

	err := errors.New(msg)
	e.AddFailure(err)
}

// AddFailure adds a validation failure to the Result
// msg is the message to be added
// if the message to be added is empty or whitespaces, nothing is added
func (e *Result) AddFailure(failure error) {
	if failure == nil {
		return
	}
	e.failures = append(e.failures, failure)
}

// IsSuccess returns true when no error has been added to the result. Otherwise, it return false
func (e Result) IsSuccess() bool {
	return len(e.failures) == 0
}

// IsFailure returns true when any error has been added to the result. Otherwise, it return false
func (e Result) IsFailure() bool {
	return !e.IsSuccess()
}

// GetFailures returns a list of all errors in the result
// If no errors are found return and empty slice
func (e Result) GetFailures() []error {
	return e.failures
}

// GetFailureMessages returns a list of all errors in the result
// If no errors are found return and empty slice
func (e Result) GetFailureMessages() []string {
	s := make([]string, len(e.failures))

	for i, v := range e.failures {
		s[i] = v.Error()
	}
	return s
}
