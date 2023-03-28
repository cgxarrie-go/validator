package examples

import (
	"errors"

	"github.com/cgxarrie-go/validator"
)

type customer struct {
	name    string
	surname string
	age     int
}

type customerValidator struct {
	validator.BaseValidator
	customer
}

func (v customerValidator) nameShouldNotBeEmpty() error {
	if v.customer.name == "" {
		return errors.New("name is empty")
	}
	return nil
}

func (v customerValidator) surnameShouldNotBeEmpty() error {
	if v.customer.surname == "" {
		return errors.New("surname is empty")
	}
	return nil
}

func (v customerValidator) ageShouldBeOver17() error {
	if v.customer.age <= 17 {
		return errors.New("age should be over 17")
	}
	return nil
}
