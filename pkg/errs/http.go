package errs

import (
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
)

func HttpUnprocessableEntity(w http.ResponseWriter, message string) {
	w.WriteHeader(http.StatusUnprocessableEntity)
	_, err := w.Write([]byte(message))
	if err != nil {
		_ = ErrResponseWriter.Throwf(applog.Log, "error writing response body, err: [%v]", err)
	}
}

func HttpInternalServerError(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
}

func HttpNotFound(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
}
