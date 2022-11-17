package govalidator

type ValidationStep struct {
	fn             func(any) error
	breakOnFailure bool
}

//BreakOnFailure forces teh validator to stop when this validation step fails
func (vs *ValidationStep) BreakOnFailure() {
	vs.breakOnFailure = true
}
