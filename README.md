# validator-go
Flexible validator for Go

# installation

```
go get github.com/cgxarrie-go/validator.git
```

# Sample of use

**Declare struct to be validated**
```
type customer struct {
	name    string
	surname string
	age     int
}
```

**Declare the validator and its constructor**
Declare validator type and its constructor along with the funcs to be the validation steps
```

type customerValidator struct {
	validator.BaseValidator
	customer  // include the type t be validated
}

func newCustomerValidator(c customer) validator.Validator {
	v := customerValidator{
		BaseValidator: validator.NewBaseValidator(),
		customer:      c,
	}

	v.AddStep(v.nameShouldNotBeEmpty).
		AddStep(v.surnameShouldNotBeEmpty).
		AddStep(v.ageShouldBeOver17)

	return v
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
```


## Use the validator

```
import (
	"errors"
	"fmt"
	
	validatorgo "github.com/cgxarrie-go/validator.git"
)

func main() {

	cust := customer{
		name:    "John",
		surname: "",
		age:     15,
	}

	validator := newCustomerValidator(cust)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("validator error : %s\n", err)
	}

}

```

** output **
```
validator error : surname is empty
validator error : age should be over 17
```


```
import (
	"errors"
	"fmt"
	
	validatorgo "github.com/cgxarrie/validator-go.git"
)

func main() {

	validable := validable{
		name:    "",
		surname: "",
		age:     15,
	}

	validator := NewValidator(validable)

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty).BreakOnFailure()
	validator.AddStep(ageShouldBeOver17)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("Validation error : %s\n", err)
	}

}

```

** output **
```
validator error : name is empty
validator error : surname is empty
```