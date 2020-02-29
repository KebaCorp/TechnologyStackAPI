package apiserver

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAPIServer_handleDashboard(t *testing.T) {
	s := New(NewConfig())
	rec := httptest.NewRecorder()

	req, err := http.NewRequest(http.MethodGet, "/api/v1/dashboard", nil)
	if err != nil {
		t.Fatal(err)
	}

	s.handleDashboard().ServeHTTP(rec, req)

	if status := rec.Code; status != http.StatusOK {
		t.Errorf("Handler returned wrong status code: got %v want %v", status, http.StatusOK)
	}
}
