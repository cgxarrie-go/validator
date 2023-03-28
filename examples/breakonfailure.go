package examples

import (
	"fmt"

	"github.com/cgxarrie-go/validator"
)

func newBreakOnFailure(c customer) validator.Validator {
	v := customerValidator{
		BaseValidator: validator.NewBaseValidator(),
		customer:      c,
	}
	v.BreakOnFailure()

	v.AddStep(v.nameShouldNotBeEmpty).
		AddStep(v.surnameShouldNotBeEmpty).
		AddStep(v.ageShouldBeOver17)

	return v
}

func runBreakOnFailure() {

	cust := customer{
		name:    "",
		surname: "",
		age:     15,
	}

	validator := newBreakOnFailure(cust)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validatorBreakOnFailire error : %s\n", err)
	}

}
