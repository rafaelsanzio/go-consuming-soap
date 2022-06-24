package handlers

import (
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/consumer"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

func HandleListContinents(w http.ResponseWriter, r *http.Request) {
	continents, err := consumer.GetContinentConsumer().List()
	if err != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, errs.ErrFmt, err)
		errs.HttpInternalServerError(w)
		return
	}

	data, err := jsonMarshal(continents)
	if err != nil {
		_ = errs.ErrMarshalingJson.Throwf(applog.Log, errs.ErrFmt, err)
		errs.HttpInternalServerError(w)
		return
	}

	_, err = write(w, data)
	if err != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, errs.ErrFmt, err)
		errs.HttpInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
