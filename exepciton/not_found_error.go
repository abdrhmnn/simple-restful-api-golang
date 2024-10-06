package exepciton

// mengikuti kontrak interface `error` bawaan golang
type NotFoundError struct {
	Error string
}

func NewNotFoundError(error string) NotFoundError {
	return NotFoundError{Error: error}
}
