# validator-go
Flexible validator for Go

# installation

```
go get github.com/cgxarrie/validator-go.git
```

# Sample of use

**Declare validable struct**
```
type validable struct {
	name    string
	surname string
	age     int
}
```

**Declare the funcs to represent the validation steps**

```
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
```

## Validator that runs all validation steps

```
import (
	"errors"
	"fmt"
	
	validatorgo "github.com/cgxarrie/validator-go.git"
)

func main() {

	validable := validable{
		name:    "John",
		surname: "",
		age:     15,
	}

	validator := NewValidator(validable)

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty)
	validator.AddStep(ageShouldBeOver17)

	result := validator.Validate()

	for _, err := range result.Errors() {
		fmt.Printf("Validation error : %s\n", err)
	}

}

```

** output **
```
validator error : surname is empty
validator error : age should be over 17
```

## Validator that breaks on first validation step failure

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

	validator := NewValidator(validable).BreakOnFailure()

	validator.AddStep(nameShouldNotBeEmpty)
	validator.AddStep(surnameShouldNotBeEmpty)
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
```


## Validator with Brek on failure declared at step level

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
	validator.AddStep(surnameShouldNotBeEmpty).BreakOnFailure
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