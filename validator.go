package govalidator

// Validator
type Validator struct {
	validable      any
	breakOnFailure bool
	steps          []ValidationStep
}

func NewValidator(validable any) *Validator {
	v := Validator{
		validable:      validable,
		breakOnFailure: false,
		steps:          []ValidationStep{},
	}
	return &v
}

//BreakOnFailure forces teh validator to stop on the first validation step failure
func (v *Validator) BreakOnFailure() *Validator {
	v.breakOnFailure = true
	return v
}

func (v *Validator) AddStep(step func(any) error) *ValidationStep {
	vs := ValidationStep{
		fn:             step,
		breakOnFailure: false,
	}
	v.steps = append(v.steps, vs)

	return &v.steps[len(v.steps)-1]
}

func (v Validator) Validate() Result {

	resp := Result{}

	for _, step := range v.steps {
		err := step.fn(v.validable)
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
