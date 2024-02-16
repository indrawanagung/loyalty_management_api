package exception

type BadRequestError struct {
	err string
}

func (e BadRequestError) Error() string {
	return e.err
}

func NewBadRequestError(error string) BadRequestError {
	return BadRequestError{err: error}
}
