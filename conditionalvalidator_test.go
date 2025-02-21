package validator

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

type testRequest struct {
	value int
}

func Test_component_Description(t *testing.T) {

	defVal := New[testRequest]()
	defVal.AddStep(func(req testRequest) error {
		return errors.New("error-in-default")
	})

	tests := []struct {
		name             string
		condVal          int
		defaultValidator Validator[testRequest]
		wantMsg          string
	}{
		{
			name:             "Condition 1",
			condVal:          1,
			defaultValidator: nil,
			wantMsg:          "error-in-validator-1",
		},
		{
			name:             "Condition 2",
			condVal:          2,
			defaultValidator: nil,
			wantMsg:          "error-in-validator-2",
		},
		{
			name:             "Condition not found and no default validator",
			condVal:          3,
			defaultValidator: nil,
			wantMsg:          "No validator found for condition",
		},
		{
			name:             "Condition not found but default validator exists",
			condVal:          3,
			defaultValidator: defVal,
			wantMsg:          "error-in-default",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Arrange
			v1 := New[testRequest]()
			v1.AddStep(func(req testRequest) error {
				return errors.New("error-in-validator-1")
			})

			v2 := New[testRequest]()
			v2.AddStep(func(req testRequest) error {
				return errors.New("error-in-validator-2")
			})

			req := testRequest{value: test.condVal}

			condVal := NewConditional[int, testRequest]()
			condVal.WithCondition(func(req testRequest) int {
				return req.value
			})
			condVal.WithValidator(1, v1)
			condVal.WithValidator(2, v2)
			condVal.WithDefaultValidator(test.defaultValidator)

			// Act
			result := condVal.Validate(req)

			// Assert
			assert.True(t, result.IsFailure())
			failures := result.GetFailureMessages()
			assert.Len(t, failures, 1)
			assert.Equal(t, test.wantMsg, failures[0])

		})
	}
}
