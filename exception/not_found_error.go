package exception

type NotFoundError struct {
	err string
}

func (e NotFoundError) Error() string {
	return e.err
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{err: error}
}
