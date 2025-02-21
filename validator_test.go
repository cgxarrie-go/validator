package validator_test

import (
	"errors"
	"testing"

	"github.com/cgxarrie-go/validator"
	"github.com/stretchr/testify/assert"
)

type dummyType struct {
}

func Test_Validator_WhenAllStepsPass_ShouldReturnSuccess(t *testing.T) {
	// Arrange

	req := dummyType{}

	step1 := func(src dummyType) error { return nil }
	step2 := func(src dummyType) error { return nil }
	step3 := func(src dummyType) error { return nil }

	v := validator.New[dummyType]()
	v.AddStep(step1)
	v.AddStep(step2)
	v.AddStep(step3)

	// Act
	result := v.Validate(req)

	// Assert
	assert.True(t, result.IsSuccess())
}

func Test_Validator_WhenNoStepsDefined_ShouldReturnFailure(t *testing.T) {
	// Arrange

	req := dummyType{}

	v := validator.New[dummyType]()

	// Act
	result := v.Validate(req)

	// Assert
	assert.True(t, result.IsFailure())
	errors := result.GetFailures()
	assert.Len(t, errors, 1)
	assert.ErrorContains(t, errors[0], "No validation steps defined")
}

func Test_Validator_WhenBreakOnFailureIsSet_ShouldReturnFailureOnFirstError(t *testing.T) {

	// Arrange
	req := dummyType{}

	step1 := func(src dummyType) error { return nil }
	step2 := func(src dummyType) error { return errors.New("Error-2") }
	step3 := func(src dummyType) error { return errors.New("Error-3") }

	v := validator.New[dummyType]().BreakOnFailure()
	v.AddStep(step1)
	v.AddStep(step2)
	v.AddStep(step3)

	// Act
	result := v.Validate(req)

	// Assert
	assert.True(t, result.IsFailure())
	errors := result.GetFailures()
	assert.Len(t, errors, 1)
	assert.ErrorContains(t, errors[0], "Error-2")
}

func Test_Validator_WhenStepBreakOnFailureIsSet_ShouldReturnFailureWithAllErrorsUntilTheStep(t *testing.T) {

	// Arrange
	req := dummyType{}

	step1 := func(src dummyType) error { return nil }
	step2 := func(src dummyType) error { return errors.New("Error-2") }
	step3 := func(src dummyType) error { return errors.New("Error-3") }
	step4 := func(src dummyType) error { return errors.New("Error-4") }

	v := validator.New[dummyType]()
	v.AddStep(step1)
	v.AddStep(step2)
	v.AddStep(step3).BreakOnFailure()
	v.AddStep(step4)

	// Act
	result := v.Validate(req)

	// Assert
	assert.True(t, result.IsFailure())
	failures := result.GetFailures()
	assert.Len(t, failures, 2)
	assert.ErrorContains(t, failures[0], "Error-2")
	assert.ErrorContains(t, failures[1], "Error-3")
}

func Test_Validator_WhenBreakOnFailureIsNotSet_ShouldReturnFailureWithAllErrors(t *testing.T) {

	// Arrange
	req := dummyType{}

	step1 := func(src dummyType) error { return nil }
	step2 := func(src dummyType) error { return errors.New("Error-2") }
	step3 := func(src dummyType) error { return errors.New("Error-3") }
	step4 := func(src dummyType) error { return errors.New("Error-4") }

	v := validator.New[dummyType]()
	v.AddStep(step1)
	v.AddStep(step2)
	v.AddStep(step3)
	v.AddStep(step4)

	// Act
	result := v.Validate(req)

	// Assert
	assert.True(t, result.IsFailure())
	errors := result.GetFailures()
	assert.Len(t, errors, 3)
	assert.ErrorContains(t, errors[0], "Error-2")
	assert.ErrorContains(t, errors[1], "Error-3")
	assert.ErrorContains(t, errors[2], "Error-4")
}
