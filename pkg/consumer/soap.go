package consumer

import (
	"strings"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/applog"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

var URL = "http://webservices.oorsprong.org/websamples.countryinfo/CountryInfoService.wso"

func SoapRequest(method, url string, payload *strings.Reader, thing interface{}) (interface{}, error) {
	req, err := httpNewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "text/xml; charset=utf-8")

	applog.Log.Infof("Requesting SOAP - Method: %s, URL %s", method, URL)
	applog.Log.Infof("Request Info - Host: %s, Body: %s", req.Host, req.Body)

	res, err := clientDo(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	applog.Log.Infof("Response SOAP - Status: %s", res.Status)

	body, err := ioutilReadAll(res.Body)
	if err != nil {
		return nil, errs.ErrReadingBytes.Throwf(applog.Log, errs.ErrFmt, err)
	}
	defer res.Body.Close()

	applog.Log.Infof("Response Body - Body: %s", string(body))

	err = xmlUnmarshal(body, &thing)
	if err != nil {
		return nil, errs.ErrUnmarshalingXML.Throwf(applog.Log, errs.ErrFmt, err)
	}

	return thing, nil
}
