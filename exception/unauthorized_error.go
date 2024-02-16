package exception

type UnauthorizedError struct {
	err string
}

func (e UnauthorizedError) Error() string {
	return e.err
}

func NewUnauthorizedError(error string) UnauthorizedError {
	return UnauthorizedError{err: error}
}
