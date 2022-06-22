package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandleAdapter(t *testing.T) {
	w := httptest.NewRecorder()
	r := httptest.NewRequest(http.MethodGet, "/ok", nil)

	handleFunc := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	}

	handle := HandleAdapter(handleFunc)

	handle.ServeHTTP(w, r)
}
