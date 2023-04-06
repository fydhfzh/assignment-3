package errs

import "net/http"

type ErrMessage interface {
	Message() string
	Status() int
	Error() string
}

type ErrorMessage struct {
	ErrMessage string `json:"message"`
	ErrStatus  int    `json:"status"`
	ErrError   string `json:"error"`
}

func (e *ErrorMessage) Message() string {
	return e.ErrMessage
}

func (e *ErrorMessage) Status() int {
	return e.ErrStatus
}

func (e *ErrorMessage) Error() string {
	return e.ErrError
}

func NewBadRequest(message string) *ErrorMessage {
	return &ErrorMessage{
		ErrMessage: message,
		ErrStatus:  http.StatusBadRequest,
		ErrError: "BAD_REQUEST",
	}
}

func NewInternalServerError(message string) *ErrorMessage {
	return &ErrorMessage{
		ErrMessage: message,
		ErrStatus: http.StatusInternalServerError,
		ErrError: "INTERNAL_SERVER_ERROR",
	}
}

func NewUnprocessableEntityError(message string) *ErrorMessage {
	return &ErrorMessage{
		ErrMessage: message,
		ErrStatus: http.StatusUnprocessableEntity,
		ErrError: "INVALID_REQUEST_BODY",
	}
}

func NewNotFoundError(message string) *ErrorMessage {
	return &ErrorMessage{
		ErrMessage: message,
		ErrStatus: http.StatusNotFound,
		ErrError: "NOT_FOUND",
	}
}

func NewUnauthorizedError(message string) *ErrorMessage {
	return &ErrorMessage{
		ErrMessage: message,
		ErrStatus: http.StatusUnauthorized,
		ErrError: "NOT_AUTHORIZED",
	}
}

func NewUnauthenticatedError(message string) *ErrorMessage {
	return &ErrorMessage{
		ErrMessage: message,
		ErrStatus: http.StatusForbidden,
		ErrError: "NOT_AUTHENTICATED",
	}
}
