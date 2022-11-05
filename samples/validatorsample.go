package main

import (
	"errors"
	"fmt"

	validatorgo "github.com/cgxarrie/validator-go.git"
)

type validable struct {
	name    string
	surname string
	age     int
}

func main() {

	validator()
	validatorBreakOnFailire()
	validatorStepBreakOnFailire()

}

func validator() {
	validable := validable{
		name:    "John",
		surname: "",
		age:     15,
	}

	validator := validatorgo.NewValidator(validable)

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty)
	validator.AddStep(ageShouldBeOver17)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validator error : %s\n", err)
	}
}

func validatorBreakOnFailire() {
	validable := validable{
		name:    "",
		surname: "",
		age:     15,
	}

	validator := validatorgo.NewValidator(validable).BreakOnFailure()

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty)
	validator.AddStep(ageShouldBeOver17)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validatorBreakOnFailire error : %s\n", err)
	}
}

func validatorStepBreakOnFailire() {
	validable := validable{
		name:    "",
		surname: "",
		age:     15,
	}

	validator := validatorgo.NewValidator(validable)

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty).BreakOnFailure()
	validator.AddStep(ageShouldBeOver17)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validatorStepBreakOnFailire error : %s\n", err)
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
