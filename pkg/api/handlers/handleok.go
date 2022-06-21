package handlers

import (
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

func HandleOK(w http.ResponseWriter, r *http.Request) {
	dataReturn := OkPayload{
		Health: 1,
		Test:   "Everthing is OK",
	}

	data, err_ := jsonMarshal(dataReturn)
	if err_ != nil {
		_ = errs.ErrMarshalingJson.Throwf(applog.Log, errs.ErrFmt, err_)
		errs.HttpInternalServerError(w)
		return
	}

	_, err_ = write(w, data)

	if err_ != nil {
		_ = errs.ErrResponseWriter.Throwf(applog.Log, errs.ErrFmt, err_)
		errs.HttpInternalServerError(w)
		return
	}

	w.WriteHeader(http.StatusOK)
}
