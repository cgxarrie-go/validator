package validator

type validationStep[T any] struct {
	breakOnFailure bool
	validator      func(T) error
	err            error
	defaultErr     error
}

func (v *validationStep[T]) BreakOnFailure() *validationStep[T] {
	v.breakOnFailure = true
	return v
}
