package validator

// Validator .
type Validator interface {
	Validate() Result
}

// BaseValidator describres a validation process
type BaseValidator struct {
	breakOnFailure bool
	steps          []validationStep
}

// ValidationStep represents a step in the validation process
type validationStep struct {
	fn             func() error
	breakOnFailure bool
}

// NewBaseValidator instantiates a base validator
// All validation steps will be run
func NewBaseValidator() BaseValidator {
	return BaseValidator{
		breakOnFailure: false,
		steps:          []validationStep{},
	}
}

// BreakOnFailure sets the validator to stop after first failure found.
func (v *BaseValidator) BreakOnFailure() *BaseValidator {
	v.breakOnFailure = true
	return v
}

// AddStep adds a validation step
func (v *BaseValidator) AddStep(step func() error) *BaseValidator {
	vs := validationStep{
		fn:             step,
		breakOnFailure: false,
	}
	v.steps = append(v.steps, vs)

	return v
}

// AddStepWithBreakOnFailure adds a validation step that stops the validation
// process when failed
func (v *BaseValidator) AddStepWithBreakOnFailure(step func() error) *BaseValidator {
	vs := validationStep{
		fn:             step,
		breakOnFailure: true,
	}
	v.steps = append(v.steps, vs)

	return v
}

// Validate runs the validation process
func (v BaseValidator) Validate() Result {

	resp := Result{}

	for _, step := range v.steps {
		err := step.fn()
		if err == nil {
			continue
		}

		resp.addError(err)
		if v.breakOnFailure || step.breakOnFailure {
			return resp
		}
	}

	return resp
}
