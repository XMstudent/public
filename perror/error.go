package errors

func New(errMsg string) error {
	return &errors{errMsg}
}

// errorString is a trivial implementation of error.
type errors struct {
	errMsg string
}

func (e *errors) Error() string {
	return e.errMsg
}
