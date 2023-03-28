package examples

import (
	"fmt"

	"github.com/cgxarrie-go/validator"
)

func newStepBreakOnFailure(c customer) validator.Validator {
	v := customerValidator{
		BaseValidator: validator.NewBaseValidator(),
		customer:      c,
	}

	v.AddStep(v.nameShouldNotBeEmpty).
		AddStepWithBreakOnFailure(v.surnameShouldNotBeEmpty).
		AddStep(v.ageShouldBeOver17)

	return v
}

func runStepBreakOnFailure() {
	cust := customer{
		name:    "",
		surname: "",
		age:     15,
	}

	validator := newStepBreakOnFailure(cust)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validatorStepBreakOnFailire error : %s\n", err)
	}
}
