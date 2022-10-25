package validatorgo

import (
	"errors"
	"fmt"
	"testing"
)

type validable struct {
	name    string
	surname string
	age     int
}

func TestValidator_ValidatorSample(t *testing.T) {

	validable := validable{
		name:    "John",
		surname: "",
		age:     15,
	}

	validator := NewValidator(validable, false)

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty)
	validator.AddStep(ageShouldBeOver17)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("Validation error : %s\n", err)
	}

}

func nameShouldNotBeEmpty(a any) error {
	v := a.(validable)
	if v.name == "" {
		return errors.New("name is empty")
	}
	return nil
}

func surnameShouldNotBeEmpty(a any) error {
	v := a.(validable)
	if v.surname == "" {
		return errors.New("surname is empty")
	}
	return nil
}

func ageShouldBeOver17(a any) error {
	v := a.(validable)
	if v.age <= 17 {
		return errors.New("age should be over 17")
	}
	return nil
}
