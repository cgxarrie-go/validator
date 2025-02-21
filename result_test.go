package validator

import (
	"errors"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestResult_AddMessage(t *testing.T) {
	tests := []struct {
		name     string
		errorMsg string
		expected string
	}{
		{name: "empty string", errorMsg: "", expected: ""},
		{name: "white spaces", errorMsg: "   ", expected: ""},
		{name: "valid message", errorMsg: "error-message", expected: "error-message"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Result{}
			e.AddFailureMessage(tt.errorMsg)

			if tt.expected == "" {
				assert.Len(t, e.GetFailures(), 0)
				return
			}

			assert.Len(t, e.GetFailures(), 1)
			assert.Equal(t, tt.expected, e.GetFailureMessages()[0])
		})
	}
}

type valueForAddParameterIsNotValid struct {
	name string
	age  int
}

func TestResult_IsSuccess(t *testing.T) {
	tests := []struct {
		name      string
		errorMsgs []string
		expected  bool
	}{
		{name: "without-errors", errorMsgs: []string{}, expected: true},
		{name: "with-errors", errorMsgs: []string{"error 1", "error 2"}, expected: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Result{}
			for _, errorMsg := range tt.errorMsgs {
				e.AddFailureMessage(errorMsg)
			}

			assert.Equal(t, tt.expected, e.IsSuccess())
		})
	}
}

func TestResult_IsFailure(t *testing.T) {
	tests := []struct {
		name      string
		errorMsgs []string
		expected  bool
	}{
		{name: "without-errors", errorMsgs: []string{}, expected: false},
		{name: "with-errors", errorMsgs: []string{"error 1", "error 2"}, expected: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Result{}
			for _, errorMsg := range tt.errorMsgs {
				err := errors.New(errorMsg)
				e.AddFailure(err)
			}

			assert.Equal(t, tt.expected, e.IsFailure())
		})
	}
}

func TestResult_Errors(t *testing.T) {
	tests := []struct {
		name      string
		errorMsgs []string
		expected  []string
	}{
		{name: "without-errors", errorMsgs: []string{}, expected: []string{}},
		{name: "with-errors", errorMsgs: []string{"error 1", "error 2"}, expected: []string{"error 1", "error 2"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := Result{}
			for _, errorMsg := range tt.errorMsgs {
				err := errors.New(errorMsg)
				e.AddFailure(err)
			}

			assert.Equal(t, tt.expected, e.GetFailureMessages())

			failureMsgs := e.GetFailureMessages()
			sort.Strings(tt.expected)
			sort.Strings(failureMsgs)

			for i := 0; i < len(tt.expected); i++ {
				assert.Equal(t, tt.expected[i], failureMsgs[i])
			}
		})
	}
}
