package validator

import (
	"errors"
	"testing"

)

func TestResult_addError(t *testing.T) {

	r := Result{}

	if expected, got := 0, len(r.errors); expected != got {
		t.Errorf("Unexpected error count. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-01"))
	if expected, got := 1, len(r.errors); expected != got {
		t.Errorf("Unexpected error count. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-02"))
	if expected, got := 2, len(r.errors); expected != got {
		t.Errorf("Unexpected error count. Expected: %v but got: %v", expected, got)
	}
}

func TestResult_IsSuccess(t *testing.T) {

	r := Result{}

	if expected, got := true, r.IsSuccess(); expected != got {
		t.Errorf("Unexpected IsSuccess. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-01"))
	if expected, got := false, r.IsSuccess(); expected != got {
		t.Errorf("Unexpected IsSuccess. Expected: %v but got: %v", expected, got)
	}
}

func TestResult_IsFailure(t *testing.T) {

	r := Result{}

	if expected, got := false, r.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-01"))
	if expected, got := true, r.IsFailure(); expected != got {
		t.Errorf("Unexpected IsFailure. Expected: %v but got: %v", expected, got)
	}
}

func TestResult_Errors(t *testing.T) {

	r := Result{}

	if expected, got := 0, len(r.Errors()); expected != got {
		t.Errorf("Unexpected Errors. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-01"))
	if expected, got := 1, len(r.Errors()); expected != got {
		t.Errorf("Unexpected Errors. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-02"))
	if expected, got := 2, len(r.Errors()); expected != got {
		t.Errorf("Unexpected Errors. Expected: %v but got: %v", expected, got)
	}
}

func TestResult_Error(t *testing.T) {

	r := Result{}

	if expected, got := "", r.Error(); expected != got {
		t.Errorf("Unexpected Error. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-01"))
	if expected, got := "error-01", r.Error(); expected != got {
		t.Errorf("Unexpected Error. Expected: %v but got: %v", expected, got)
	}

	r.addError(errors.New("error-02"))
	if expected, got := "error-01,error-02", r.Error(); expected != got {
		t.Errorf("Unexpected Error. Expected: %v but got: %v", expected, got)
	}
}
