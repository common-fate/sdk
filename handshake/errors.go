package handshake

import (
	"bytes"
	"fmt"
)

const (
	// unknown
	ErrUnknownError = 1111

	// internal
	ErrInternalError = 1100
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

// Error implements the error interface
func (se *ServerError) Error() string {
	buf := &bytes.Buffer{}
	buf.WriteString(se.Message)

	fmt.Fprintf(buf, " (errno %v)", se.Num)

	return buf.String()
}
