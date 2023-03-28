package validator

import (
	"errors"
	"strings"
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestBaseValidator_New_ShouldReturnNoBreakOnFailure(t *testing.T) {

	// Arrange

	// Act
	v := NewBaseValidator()

	// Assert
	assert.Equal(t, false, v.breakOnFailure)
	assert.Equal(t, 0, len(v.steps))
}

func TestBaseValidator_NewBreakOnFailure_ShouldReturnBreakOnFailure(t *testing.T) {

	// Arrange

	// Act
	v := NewBaseValidator()
	v.BreakOnFailure()

	// Assert
	assert.Equal(t, true, v.breakOnFailure)
	assert.Equal(t, 0, len(v.steps))
}

func TestBaseValidator_AddStep(t *testing.T) {

	v := BaseValidator{}

	if expected, got := 0, len(v.steps); expected != got {
		t.Errorf("Unexpected steps count. Expected: %v but got: %v", expected, got)
	}

	v.AddStep(func() error {
		return nil
	})

	if expected, got := 1, len(v.steps); expected != got {
		t.Errorf("Unexpected steps count. Expected: %v but got: %v", expected, got)
	}

	v.AddStep(func() error {
		return nil
	})

	if expected, got := 2, len(v.steps); expected != got {
		t.Errorf("Unexpected steps count. Expected: %v but got: %v", expected, got)
	}
}

func TestBaseValidator_Validate_FailedStepShouldReturnFailure(t *testing.T) {

	v := BaseValidator{}
	v.AddStep(func() error {
		return errors.New("error-step-01")
	})

	result := v.Validate()

	if expected, got := true, result.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}
}

func TestBaseValidator_Validate_SuccessStepShouldReturnSuccess(t *testing.T) {

	v := BaseValidator{}
	v.AddStep(func() error {
		return nil
	})

	result := v.Validate()

	if expected, got := true, result.IsSuccess(); expected != got {
		t.Errorf("Unexpected IsSuccess. Expected: %v but got: %v", expected, got)
	}
}

func TestBaseValidator_Validate_BreakOnFailureFalseShouldReturnAllErrors(t *testing.T) {

	v := BaseValidator{}
	v.AddStep(func() error {
		return errors.New("error-step-01")
	})

	v.AddStep(func() error {
		return errors.New("error-step-02")
	})

	result := v.Validate()

	if expected, got := true, result.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}

	if expected, got := 2, len(result.Errors()); expected != got {
		t.Errorf("Unexpected Errors count. Expected: %v but got: %v", expected, got)
	}

	if expected, got := true, strings.Contains(result.Error(), "error-step-01"); expected != got {
		t.Errorf("Expected error-step-01 not found. Expected: %v but got: %v", expected, got)
	}

	if expected, got := true, strings.Contains(result.Error(), "error-step-02"); expected != got {
		t.Errorf("Expected error-step-02 not found. Expected: %v but got: %v", expected, got)
	}
}

func TestBaseValidator_Validate_BreakOnFailureTrueShouldReturnFirstError(t *testing.T) {

	v := NewBaseValidator()
	v.BreakOnFailure()

	v.AddStep(func() error {
		return errors.New("error-step-01")
	})

	v.AddStep(func() error {
		return errors.New("error-step-02")
	})

	result := v.Validate()

	if expected, got := true, result.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}

	if expected, got := 1, len(result.Errors()); expected != got {
		t.Errorf("Unexpected Errors count. Expected: %v but got: %v", expected, got)
	}

	if expected, got := true, strings.Contains(result.Error(), "error-step-01"); expected != got {
		t.Errorf("Expected error-step-01 not found. Expected: %v but got: %v", expected, got)
	}

	if expected, got := false, strings.Contains(result.Error(), "error-step-02"); expected != got {
		t.Errorf("Unexpected error-step-02 found. Expected: %v but got: %v", expected, got)
	}
}

func TestBaseValidator_Validate_ValidationStepBreakOnFailureTrueShouldReturnErrorsUntilStep(t *testing.T) {

	v := BaseValidator{}
	v.AddStep(func() error {
		return errors.New("error-step-01")
	})

	v.AddStepWithBreakOnFailure(func() error {
		return errors.New("error-step-02")
	})

	v.AddStep(func() error {
		return errors.New("error-step-03")
	})

	result := v.Validate()

	if expected, got := true, result.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}

	if expected, got := 2, len(result.Errors()); expected != got {
		t.Errorf("Unexpected Errors count. Expected: %v but got: %v", expected, got)
	}

	if expected, got := true, strings.Contains(result.Error(), "error-step-01"); expected != got {
		t.Errorf("Expected error-step-01 not found. Expected: %v but got: %v", expected, got)
	}

	if expected, got := true, strings.Contains(result.Error(), "error-step-02"); expected != got {
		t.Errorf("Unexpected error-step-02 found. Expected: %v but got: %v", expected, got)
	}

	if expected, got := false, strings.Contains(result.Error(), "error-step-03"); expected != got {
		t.Errorf("Unexpected error-step-02 found. Expected: %v but got: %v", expected, got)
	}

}
