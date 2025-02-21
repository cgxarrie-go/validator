package validator

type conditionalValidator[TCond any, TRequest any] struct {
	validators       map[any]Validator[TRequest]
	defaultValidator Validator[TRequest]
	condition        func(TRequest) TCond
}

func NewConditional[TCond any, TRequest any]() *conditionalValidator[TCond, TRequest] {
	return &conditionalValidator[TCond, TRequest]{
		validators: make(map[any]Validator[TRequest]),
	}
}

func (v *conditionalValidator[TCond, TRequest]) WithCondition(condition func(TRequest) TCond) *conditionalValidator[TCond, TRequest] {
	v.condition = condition
	return v
}

func (v *conditionalValidator[TCond, TRequest]) WithValidator(condition TCond, validator Validator[TRequest]) *conditionalValidator[TCond, TRequest] {
	v.validators[condition] = validator
	return v
}

func (v *conditionalValidator[TCond, TRequest]) WithDefaultValidator(validator Validator[TRequest]) *conditionalValidator[TCond, TRequest] {
	v.defaultValidator = validator
	return v
}

func (v *conditionalValidator[TCond, TRequest]) Validate(req TRequest) Result {
	condition := v.condition(req)
	validator, ok := v.validators[condition]
	if !ok {
		if v.defaultValidator != nil {
			return v.defaultValidator.Validate(req)
		}

		result := Result{}
		result.AddFailureMessage("No validator found for condition")
		return result
	}

	return validator.Validate(req)
}
