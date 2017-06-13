package errors

import (
	"log"
	"net/http"
	"testing"
)

func TestValidationErrors(t *testing.T) {
	ms := []string{
		"something went wrong somewhere.",
		"something else went wrong",
	}

	em := "something went wrong somewhere., something else went wrong"

	err := ValidationErrors(ms)

	pe := Parse(err.Error())

	if pe.Detail != em {
		log.Fatalf("Expected '%s' got '%s'", em, pe.Detail)
	}

	if pe.Code != http.StatusBadRequest {
		log.Fatalf("Expected '%d' got '%s'", http.StatusBadRequest, pe.Status)
	}
}

func TestBadRequest(t *testing.T) {
	id := "id"
	details := "details"
	err := BadRequest(id, details)

	pe := Parse(err.Error())

	if pe.Detail != details {
		log.Fatalf("Expected '%s' got '%s'", details, pe.Detail)
	}

	if pe.Code != http.StatusBadRequest {
		log.Fatalf("Expected '%d' got '%s'", http.StatusBadRequest, pe.Status)
	}
}

func TestUnauthorized(t *testing.T) {
	id := "id"
	details := "details"
	err := Unauthorized(id, details)

	pe := Parse(err.Error())

	if pe.Detail != details {
		log.Fatalf("Expected '%s' got '%s'", details, pe.Detail)
	}

	if pe.Code != http.StatusUnauthorized {
		log.Fatalf("Expected '%d' got '%s'", http.StatusUnauthorized, pe.Status)
	}
}

func TestForbidden(t *testing.T) {
	id := "id"
	details := "details"
	err := Forbidden(id, details)

	pe := Parse(err.Error())

	if pe.Detail != details {
		log.Fatalf("Expected '%s' got '%s'", details, pe.Detail)
	}

	if pe.Code != http.StatusForbidden {
		log.Fatalf("Expected '%d' got '%s'", http.StatusForbidden, pe.Status)
	}
}

func TestNotFound(t *testing.T) {
	id := "id"
	details := "details"
	err := NotFound(id, details)

	pe := Parse(err.Error())

	if pe.Detail != details {
		log.Fatalf("Expected '%s' got '%s'", details, pe.Detail)
	}

	if pe.Code != http.StatusNotFound {
		log.Fatalf("Expected '%d' got '%s'", http.StatusNotFound, pe.Status)
	}
}

func TestInternalServerError(t *testing.T) {
	id := "id"
	details := "details"
	err := InternalServerError(id, details)

	pe := Parse(err.Error())

	if pe.Detail != details {
		log.Fatalf("Expected '%s' got '%s'", details, pe.Detail)
	}

	if pe.Code != http.StatusInternalServerError {
		log.Fatalf("Expected '%d' got '%s'", http.StatusInternalServerError, pe.Status)
	}
}
