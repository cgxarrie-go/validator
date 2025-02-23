# Validator

- [Installation](#installation)
- [Result](#result) 
- [Validation Step](#validation-step)
- [Validator](#validator)
- [Conditional Validator](#conditional-validator)

## Installation
```Bash
go get github.com/cgxarrie-go/validator@latest
```

## How to use the Validator

### Result

Offers a versatile response object to return as validation result with the ability to return a collection of failures

#### Methods

- AddFailureMessage(s string): Adds a failure with the given string
- AddFailure(err error): Adds a failure for the given error
- GetFailureMessages(): Return a slice containing all failures stringified
- GetFailures(): Returns a slice containing all the failures
- IsFailure(): Returns true if the number of failures is greater than 0. Otherwise it returns false
- IsSuccess(): Returns true if the number of failures is 0. Otherwise it returns false

#### Example

```Go
import (
    "github.com/cgxarrie-go/validator"
)

func main() {
    objectToValidate := any object instance 

    result := validate(objectToValidate)
    // Any of the following conditions can be checked

    if result.IsFailure() {
        // actions
        fmt.Println(result.Error()) // prints all the errors concatenated by semicolon (;)

        for _, v := range result.Getrrors()) {
            fmt.Println(v) // prints each error individually
            }
    }

    if result.IsSuccess(){
        // actions
    }
}

func validate(objToValidate anyType) (result validator.result) {

    if (objectToValidate does not meet condition 1) {
        result.AddFailureMessage("condition 1 not met")
    }

    if (objectToValidate does not meet condition 2) {
        f := errors.New("condition 2 not met")
        result.AddFailure(f)
    }

    return result

}

```

### Validation Step
Functon to be added to  a validator as a validation step

#### Methods
- BreakOnFailure(): Forces the validator to stop processing validations steps if the current validation fails

```Go
import (
    "github.com/cgxarrie-go/validator"
)

type dummyType struct {
    Field1 string
    Field2 int
    Field3 bool
    ...
}

func main() {
    condition1 := func(src dummyType) error { 
        // validation logic
        return error if fails, otherwise
        return nil 
        }

    vldtr := validator.NewValidator[dummyType]()
    vldtr.AddStep(condition1).
        BreakOnFailure() // optional - stops processing if the step fails

    req := Instance_Of_DummyType
    result := vldtr.Validate(req)
}

```

### Validator

Offers a versatile response object to return as validation result with the ability to return a collection of error messages

Validation is build by adding validation steps to a validator instance

vldtr.BreakOnFailure() configures the validator to stop processing steps when the first failure is encountered
ValidatorStep.BreakOnFailure() configures the step to stop processing if the step fails

#### Methods

- BreakOnFailure(): If added to a validator, the validator will stop processing steps whenever there is a failure
- AddStep(fn): Adds a step to the validator
- AddValidator(validator): Adds a validator a a step to the current validator.The result of the validation will include all the added steps and the steps of the added validator
-Validate(request): Runs the validation steps and returns the result

#### Example

```Go
import (
    "github.com/cgxarrie-go/validator"
)

type dummyType struct {
    Field1 string
    Field2 int
    Field3 bool
    ...
}

func main() {
    condition1 := func(src dummyType) error { 
        // validation logic
        return error if fails, otherwise
        return nil 
        }

        condition2 := func(src dummyType) error {
        // validation logic
        return error if fails, otherwise
        return nil
    }

    validator1 := validator.NewValidator[dummyType]()
    validator1.AddStep(condition1)

    validator2 := validator.NewValidator[dummyType]().
        BreakOnFailure()
    validator2.AddValidator(validator1)
    validator2.AddStep(condition2)

    req := Instance_Of_DummyType
    result := validator2.Validate(req)
}


```

### Conditional Validator
This is a validator composed by many validators, each of which is executed based on a condition

#### Methods
- WithCondition(fn): Adds a condition to the validator. The condition is a function returning the condition type
- WithValidator(value, validator): States the validator to be run for a specific condition value
- WithDefaultValidator(validator): States the validator to be run if no condition is met. If no default validator is set, and thhere is no validator for the condition value, the conditional validator will return a failure result
- Validate(request): Evaluates the condition and runs the validation for the corresponding validator

#### Example

```Go
import (
    "github.com/cgxarrie-go/validator"
)

type dummyType struct {
    Field1 string
    Field2 int
    Field3 bool
    ...
}

func main() {

    defaultvldtr := validator.NewValidator[dummyType]()
    validator1 := validator.NewValidator[dummyType]()
    validator2 := validator.NewValidator[dummyType]()
    validator3 := validator.NewValidator[dummyType]()

    condition1 := func(req dummyType) int { 
        return req.Field2 
        }

    vldtr := validator.NewConditionalValidator[int, dummyType]()
    vldtr.WithCondition(condition1)
    vldtr.WithDefaultValidator(defaultValidator)
    vldtr.WithValidator(1, validator1)
    vldtr.WithValidator(2, validator2)
    vldtr.WithValidator(3, validator3)


    req := Instance_Of_DummyType
    result := vldtr.Validate(req)
}

```
