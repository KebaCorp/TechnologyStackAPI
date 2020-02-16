package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAPIServer_handleTechnologies(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/technologies", nil)
	s.handleTechnologies().ServeHTTP(rec, req)
	assert.Equal(t, rec.Body.String(), "Hello")
}
