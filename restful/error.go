package restful

import "fmt"

type HttpLogicError struct {
	ErrorCode int    `json:"code"`
	Message   string `json:"message"`
}

func (e *HttpLogicError) Error() string {
	return fmt.Sprintf("Error Code: %d, Message: %s", e.ErrorCode, e.Message)
}

type ErrorOption func(*HttpLogicError)

func WithErrorCode(code int) ErrorOption {
	return func(hle *HttpLogicError) {
		hle.ErrorCode = code
	}
}

func WithMessage(message string) ErrorOption {
	return func(hle *HttpLogicError) {
		hle.Message = message
	}
}

func NewHttpLogicError(opts ...ErrorOption) *HttpLogicError {
	err := &HttpLogicError{
		ErrorCode: 100,
		Message:   "Failed",
	}

	for _, opt := range opts {
		opt(err)
	}
	return err
}
