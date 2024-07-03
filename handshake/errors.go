package handshake

import (
	"bytes"
	"fmt"
)

const (
	// 1000 - 1199 server errors
	// unknown
	ErrUnknownError = 1000

	// internal
	ErrInternalError = 1100

	// 1200 -  grant errors
	ErrGrantInactiveError = 1200
)

// ServerError is a known error type
type ServerError struct {
	Num     int
	Message string
}

func InternalServerError() ServerError {
	return ServerError{
		Num:     ErrInternalError,
		Message: "internal server error",
	}
}
func GrantInactiveError() ServerError {
	return ServerError{
		Num:     ErrGrantInactiveError,
		Message: "The grant is inactive",
	}
}

// Error implements the error interface
func (se *ServerError) Error() string {
	buf := &bytes.Buffer{}
	buf.WriteString(se.Message)

	fmt.Fprintf(buf, " (errno %v)", se.Num)

	return buf.String()
}
