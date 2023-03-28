package examples

import (
	"fmt"

	"github.com/cgxarrie-go/validator"
)

func newRunAllSteps(c customer) validator.Validator {
	v := customerValidator{
		BaseValidator: validator.NewBaseValidator(),
		customer:      c,
	}

	v.AddStep(v.nameShouldNotBeEmpty).
		AddStep(v.surnameShouldNotBeEmpty).
		AddStep(v.ageShouldBeOver17)

	return v
}

func runAllSteps() {
	cust := customer{
		name:    "John",
		surname: "",
		age:     15,
	}

	validator := newRunAllSteps(cust)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validator error : %s\n", err)
	}

}
