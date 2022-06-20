package consumer

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
)

var httpNewRequest = http.NewRequest

func fakeNewRequest(method string, url string, body io.Reader) (*http.Request, error) {
	return nil, errs.ErrRequest
}

func restoreNewRequest(replace func(method string, url string, body io.Reader) (*http.Request, error)) {
	httpNewRequest = replace
}

var client = &http.Client{}
var clientDo = client.Do

func fakeClientDo(req *http.Request) (*http.Response, error) {
	return nil, errs.ErrRequest
}

func restoreClientDo(replace func(req *http.Request) (*http.Response, error)) {
	clientDo = replace
}

var ioutilReadAll = ioutil.ReadAll

func fakeIoutilReadAll(r io.Reader) ([]byte, error) {
	return nil, errs.ErrRequest
}

func restoreIoutilReadAll(replace func(r io.Reader) ([]byte, error)) {
	ioutilReadAll = replace
}

var xmlUnmarshal = xml.Unmarshal

func fakeXMLUnmarshal(data []byte, v any) error {
	return errs.ErrUnmarshalingXML
}

func restoreXMLUnmarshal(replace func(data []byte, v any) error) {
	xmlUnmarshal = replace
}
