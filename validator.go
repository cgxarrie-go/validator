package validatorgo

// Validator
type Validator struct {
	validable      any
	breakOnFailure bool
	steps          []func(any) error
}

func (v *Validator) AddStep(step func(any) error) {
	v.steps = append(v.steps, step)
}

func (v Validator) Validate() Result {

	resp := Result{}

	for _, step := range v.steps {
		err := step(v.validable)
		if err == nil {
			continue
		}

		resp.addError(err)
		if v.breakOnFailure {
			return resp
		}
	}

	return resp
}
