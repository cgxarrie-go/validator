package validatorgo

import (
	"errors"
	"strings"
	"testing"
)

func TestValidator_AddStep(t *testing.T) {

	type validable struct {
		a int64
		b int64
	}

	v := NewValidator(validable{a: 10, b: 20})

	if expected, got := 0, len(v.steps); expected != got {
		t.Errorf("Unexpected steps count. Expected: %v but got: %v", expected, got)
	}

	v.AddStep(func(a any) error {
		return nil
	})

	if expected, got := 1, len(v.steps); expected != got {
		t.Errorf("Unexpected steps count. Expected: %v but got: %v", expected, got)
	}

	v.AddStep(func(a any) error {
		return nil
	})

	if expected, got := 2, len(v.steps); expected != got {
		t.Errorf("Unexpected steps count. Expected: %v but got: %v", expected, got)
	}
}

func TestValidator_Validate_FailedStepShouldReturnFailure(t *testing.T) {

	type validable struct {
		a int64
		b int64
	}

	v := NewValidator(validable{a: 10, b: 20})
	v.AddStep(func(a any) error {
		return errors.New("error-step-01")
	})

	result := v.Validate()

	if expected, got := true, result.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}
}

func TestValidator_Validate_SuccessStepShouldReturnSuccess(t *testing.T) {

	type validable struct {
		a int64
		b int64
	}

	v := NewValidator(validable{a: 10, b: 20})
	v.AddStep(func(a any) error {
		return nil
	})

	result := v.Validate()

	if expected, got := true, result.IsSuccess(); expected != got {
		t.Errorf("Unexpected IsSuccess. Expected: %v but got: %v", expected, got)
	}
}

func TestValidator_Validate_BreakOnFailureFalseShouldReturnAllErrors(t *testing.T) {

	type validable struct {
		a int64
		b int64
	}

	v := NewValidator(validable{a: 10, b: 20})
	v.AddStep(func(a any) error {
		return errors.New("error-step-01")
	})

	v.AddStep(func(a any) error {
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

func TestValidator_Validate_BreakOnFailureTrueShouldReturnFirstError(t *testing.T) {

	type validable struct {
		a int64
		b int64
	}

	v := NewValidator(validable{a: 10, b: 20}).BreakOnFailure()
	v.AddStep(func(a any) error {
		return errors.New("error-step-01")
	})

	v.AddStep(func(a any) error {
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

func TestValidator_Validate_ValidationStepBreakOnFailureTrueShouldReturnErrorsUntilStep(t *testing.T) {

	type validable struct {
		a int64
		b int64
	}

	v := NewValidator(validable{a: 10, b: 20})
	v.AddStep(func(a any) error {
		return errors.New("error-step-01")
	})

	v.AddStep(func(a any) error {
		return errors.New("error-step-02")
	}).BreakOnFailure()

	v.AddStep(func(a any) error {
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
