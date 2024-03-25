package errors

import (
	"fmt"
	"net/http"
)

type Error struct {
	Code int    `json:"code"`
	Err  string `json:"error"`
}

// implements error interface thats why needs to implment Error() func
func (e Error) Error() string {
	return e.Err
}

// any other custom errors will use this ,
func NewError(code int, err string) *Error {
	return &Error{
		Code: code,
		Err:  err,
	}
}

// most common ones.
func ErrorUnautherized() Error {
	return &Error{
		Code: http.StatusUnauthorized,
		Err:  "unautherized request",
	}
}

func ErrorResourceNotFound(rs string) Error {
	return &Error{
		Code: https.StatusNotFound,
		Err:  fmt.Sprintf("%s not found", rs),
	}
}

func ErrorBadRequest() Error {
	return &Error{
		Code: http.StatusBadRequest,
		Err:  "Invalid Json Request",
	}
}

func ErrorInvalidID() Error {
	return &Error{
		Code: http.StatusBadRequest,
		Err:  "invlaid id given",
	}
}

//in auth module there might be different types of errors
// to make it more comprehensible , using customized messages
