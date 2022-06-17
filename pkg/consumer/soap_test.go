package consumer

import (
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/continent"
	"github.com/stretchr/testify/assert"
)

func TestSoapRequest(t *testing.T) {
	payload := strings.NewReader(`<?xml version="1.0" encoding="utf-8"?>
	<soap12:Envelope xmlns:soap12="http://www.w3.org/2003/05/soap-envelope">
	<soap12:Body>
		<ListOfCountryNamesByName xmlns="http://www.oorsprong.org/websamples.countryinfo">
		</ListOfCountryNamesByName>
	</soap12:Body>
	</soap12:Envelope>`)

	testCases := []struct {
		Name                  string
		NewRequestFunction    func(method string, url string, body io.Reader) (*http.Request, error)
		ClientDoFunction      func(req *http.Request) (*http.Response, error)
		IoutilReadAllFunction func(r io.Reader) ([]byte, error)
		XMLUnmarshalFunction  func(data []byte, v any) error
		ExpectedError         bool
	}{
		{
			Name:                  "Getting body correct",
			NewRequestFunction:    httpNewRequest,
			ClientDoFunction:      clientDo,
			IoutilReadAllFunction: ioutilReadAll,
			XMLUnmarshalFunction:  xmlUnmarshal,
			ExpectedError:         false,
		}, {
			Name:                  "Error on throwing new request",
			NewRequestFunction:    fakeNewRequest,
			ClientDoFunction:      clientDo,
			IoutilReadAllFunction: ioutilReadAll,
			XMLUnmarshalFunction:  xmlUnmarshal,
			ExpectedError:         true,
		}, {
			Name:                  "Error on doing request",
			NewRequestFunction:    httpNewRequest,
			ClientDoFunction:      fakeClientDo,
			IoutilReadAllFunction: ioutilReadAll,
			XMLUnmarshalFunction:  xmlUnmarshal,
			ExpectedError:         true,
		}, {
			Name:                  "Error reading bytes",
			NewRequestFunction:    httpNewRequest,
			ClientDoFunction:      clientDo,
			IoutilReadAllFunction: fakeIoutilReadAll,
			XMLUnmarshalFunction:  xmlUnmarshal,
			ExpectedError:         true,
		}, {
			Name:                  "Error unmarshal xml",
			NewRequestFunction:    httpNewRequest,
			ClientDoFunction:      clientDo,
			IoutilReadAllFunction: ioutilReadAll,
			XMLUnmarshalFunction:  fakeXMLUnmarshal,
			ExpectedError:         true,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		httpNewRequest = tc.NewRequestFunction
		defer restoreNewRequest(httpNewRequest)

		clientDo = tc.ClientDoFunction
		defer restoreClientDo(clientDo)

		ioutilReadAll = tc.IoutilReadAllFunction
		defer restoreIoutilReadAll(ioutilReadAll)

		xmlUnmarshal = tc.XMLUnmarshalFunction
		defer restoreXMLUnmarshal(xmlUnmarshal)

		rContinent, err := SoapRequest(http.MethodPost, URL, payload, continent.ResponseContinentInfo{})

		if tc.ExpectedError {
			assert.Nil(t, rContinent)
			assert.NotNil(t, err)
		} else {
			r := rContinent.(continent.ResponseContinentInfo)

			assert.NotNil(t, r)
			assert.NoError(t, err)
		}

	}
}
