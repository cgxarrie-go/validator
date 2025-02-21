// package validator provides a generic validation framework for validating
// data of any type.
//
// The validator type is a generic type that allows for the definition of
// multiple validation steps that can be applied to a given data type. Each
// validation step is represented by a validationStep struct, which contains
// a validation function and a flag indicating whether to break on failure.
//
// The validator type provides methods for adding validation steps, setting
// the break-on-failure flag, and performing validation on a given data
// instance.
//
// Types:
//   - validator[T any]: A generic validator type that holds validation steps
//     and a flag indicating whether to break on failure.
//
// Functions:
//   - (v validator[T]) Validate(src T) Result: Validates the given data instance
//     using the defined validation steps and returns a Result containing any
//     validation errors.
//   - New[T any]() *validator[T]: Creates a new validator instance with no
//     validation steps and the break-on-failure flag set to false.
//   - (v *validator[T]) BreakOnFailure() *validator[T]: Sets the break-on-failure
//     flag to true and returns the validator instance.
//   - (v *validator[T]) AddStep() *validationStep[T]: Adds a new validation step
//     to the validator and returns a pointer to the newly added validation step.
//   - returnError(customError, defaultError error) error: Returns the custom error
//     if it is not nil, otherwise returns the default error.
package validator

type Validator[T any] interface {
	Validate(src T) Result
}

type validator[T any] struct {
	breakOnFailure bool
	validators     []validationStep[T]
}

func (v validator[T]) Validate(src T) Result {
	result := Result{}

	if len(v.validators) == 0 {
		result.AddFailureMessage("No validation steps defined")
		return result
	}

	for _, step := range v.validators {
		err := step.validator(src)

		if err != nil {

			if res, ok := err.(Result); ok {
				for _, v := range res.GetFailures() {
					result.AddFailure(v)
				}
			} else {
				result.AddFailure(err)
			}

			if step.breakOnFailure || v.breakOnFailure {
				return result
			}
		}
	}

	return result
}

// New creates a new validator instance with no validation steps and the
// break-on-failure flag set to false.
func New[T any]() *validator[T] {
	return &validator[T]{
		breakOnFailure: false,
		validators:     make([]validationStep[T], 0),
	}

}

func (v *validator[T]) BreakOnFailure() *validator[T] {
	v.breakOnFailure = true
	return v
}

func (v *validator[T]) AddStep(steps ...func(req T) error) *validationStep[T] {

	if steps == nil {
		steps = []func(req T) error{func(T) error { return nil }}
	}

	for _, step := range steps {
		validationStep := validationStep[T]{
			breakOnFailure: false,
			validator:      step,
		}

		v.validators = append(v.validators, validationStep)

	}

	return &v.validators[len(v.validators)-1]
}

func (v *validator[T]) AddValidator(validator Validator[T]) {

	step := func(req T) error {
		result := validator.Validate(req)
		if result.IsFailure() {
			return result
		}
		return nil
	}

	v.AddStep(step)
}
