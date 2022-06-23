package handlers

import (
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/consumer"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

func HandleListCountries(w http.ResponseWriter, r *http.Request) {
	var limit, offset int

	limitStr := r.URL.Query().Get("limit")
	offsetStr := r.URL.Query().Get("offset")

	countries, err := consumer.GetCountryConsumer().List()
	if err != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, errs.ErrFmt, err)
		errs.HttpInternalServerError(w)
		return
	}

	if limitStr != "" && offsetStr != "" {
		limit, err = atoi(limitStr)
		if err != nil {
			_ = errs.ErrConvertingStringToInt.Throwf(applog.Log, errs.ErrFmt, err)
			errs.HttpInternalServerError(w)
			return
		}

		offset, err = atoi(offsetStr)
		if err != nil {
			_ = errs.ErrConvertingStringToInt.Throwf(applog.Log, errs.ErrFmt, err)
			errs.HttpInternalServerError(w)
			return
		}

		countries = countries[offset : limit+offset]
	}

	if limitStr != "" && offsetStr == "" {
		limit, err = atoi(limitStr)
		if err != nil {
			_ = errs.ErrConvertingStringToInt.Throwf(applog.Log, errs.ErrFmt, err)
			errs.HttpInternalServerError(w)
			return
		}

		countries = countries[:limit]
	}

	if limitStr == "" && offsetStr != "" {
		offset, err = atoi(offsetStr)
		if err != nil {
			_ = errs.ErrConvertingStringToInt.Throwf(applog.Log, errs.ErrFmt, err)
			errs.HttpInternalServerError(w)
			return
		}

		countries = countries[offset:]
	}

	data, err := jsonMarshal(countries)
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
