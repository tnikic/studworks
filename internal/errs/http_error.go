package errs

import "fmt"

type HttpError struct {
	Code    int
	Message string
	Err     error
}

func NewHttpError(code int, message string, err error) *HttpError {
	return &HttpError{
		Code:    code,
		Message: message,
		Err:     err,
	}
}

func (e *HttpError) Error() string {
	return fmt.Sprintf("code=%d, message=%s, err=%v", e.Code, e.Message, e.Err)
}
