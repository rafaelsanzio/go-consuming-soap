package handlers

import (
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
)

func HandleAdapter(hf http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		applog.Log.Infof("Requesting - Method: %s, URL %s", r.Method, r.URL)
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		hf(w, r)
	}
}
