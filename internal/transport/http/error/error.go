package error

type HTTPError struct {
	Err string `json:"error"`
}

func NewError(message string) HTTPError {
	return HTTPError{Err: message}
}
