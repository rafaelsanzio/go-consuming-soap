package handlers

import (
	"net/http"
	"strings"

	"github.com/gorilla/mux"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/consumer"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

func HandleGetCountryInfo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	countryCode := strings.ToUpper(vars["code"])

	err := validateCountryCode(countryCode)
	if err != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, errs.ErrFmt, err)
		errs.HttpNotFound(w)
		return
	}

	countryInfo, err := consumer.GetCountryConsumer().Info(countryCode)
	if err != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, errs.ErrFmt, err)
		errs.HttpInternalServerError(w)
		return
	}

	data, err := jsonMarshal(countryInfo)
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
