package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

var jsonMarshal = json.Marshal

func fakeMarshal(v interface{}) ([]byte, error) {
	return []byte{}, errs.ErrMarshalingJson
}

func restoreMarshal(replace func(v interface{}) ([]byte, error)) {
	jsonMarshal = replace
}

var write = http.ResponseWriter.Write

func fakeWrite(http.ResponseWriter, []byte) (int, error) {
	return 0, errs.ErrResponseWriter
}

func restoreWrite(replace func(http.ResponseWriter, []byte) (int, error)) {
	write = replace
}

var atoi = strconv.Atoi

func fakeAtoi(s string) (int, error) {
	return 0, errs.ErrConvertingStringToInt
}

func restoreAtoi(replace func(s string) (int, error)) {
	atoi = replace
}

var httpGet = http.Get

func fakeHttpGet(url string) (resp *http.Response, err error) {
	return nil, errs.ErrConvertingStringToInt
}

func restoreHttpGet(replace func(url string) (resp *http.Response, err error)) {
	httpGet = replace
}

var rateLimitAllow = rateLimit.Allow

func fakeRateLimitAllow() bool {
	return false
}

func restoreRateLimitAllow(replace func() bool) {
	rateLimitAllow = replace
}

func validateCountryCode(countryCode string) error {
	if countryCode == "" || len(countryCode) != 2 {
		return errs.ErrGettingParam.Throwf(applog.Log, errs.ErrFmt, "param with bad format")
	}

	res, err := httpGet(fmt.Sprintf("https://restcountries.com/v3.1/alpha/%s", countryCode))
	if err != nil {
		return errs.ErrNewRequest.Throwf(applog.Log, errs.ErrFmt, err)
	}

	if res.StatusCode != http.StatusOK {
		return errs.ErrCountryNotFound.Throwf(applog.Log, errs.ErrFmt, fmt.Sprintf("country code: %s", countryCode))
	}

	return nil
}
