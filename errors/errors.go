// These are extensions or wrappers to the go-micro package. The wrappers are in
// place just incase we decide to drop micro or their error package.package main

package errors

import (
	"net/http"
	"strings"

	me "github.com/micro/go-micro/errors"
)

// ValidationErrors returns the go micro error.
func ValidationErrors(messages []string) error {
	return &me.Error{
		Id:     "",
		Code:   http.StatusBadRequest,
		Detail: strings.Join(messages, ", "),
		Status: http.StatusText(http.StatusBadRequest),
	}
}

// Parse is a wrapper of go-micro's error.
func Parse(err string) *me.Error {
	return me.Parse(err)
}

// BadRequest is a wrapper of go-micro's error.
func BadRequest(id, detail string) error {
	return me.BadRequest(id, detail)
}

// Unauthorized is a wrapper of go-micro's error.
func Unauthorized(id, detail string) error {
	return me.Unauthorized(id, detail)
}

// Forbidden is a wrapper of go-micro's error.
func Forbidden(id, detail string) error {
	return me.Forbidden(id, detail)
}

// NotFound is a wrapper of go-micro's error.
func NotFound(id, detail string) error {
	return me.NotFound(id, detail)
}

// InternalServerError is a wrapper of go-micro's error.
func InternalServerError(id, detail string) error {
	return me.InternalServerError(id, detail)
}
