package consumer

import (
	"io"
	"net/http"
	"testing"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/continent"

	"github.com/stretchr/testify/assert"
)

func TestContinentConsumerList(t *testing.T) {

	testCases := []struct {
		Name                 string
		NewRequestFunction   func(method string, url string, body io.Reader) (*http.Request, error)
		XMLUnmarshalFunction func(data []byte, v interface{}) error
		ExpectedContinents   []continent.Continent
	}{
		{
			Name:                 "Should return all the continents",
			NewRequestFunction:   httpNewRequest,
			XMLUnmarshalFunction: xmlUnmarshal,
			ExpectedContinents: []continent.Continent{
				{Code: "AF", Name: "Africa"},
				{Code: "AN", Name: "Antarctica"},
				{Code: "AS", Name: "Asia"},
				{Code: "EU", Name: "Europe"},
				{Code: "OC", Name: "Ocenania"},
				{Code: "AM", Name: "The Americas"},
			},
		}, {
			Name:                 "Throw an error on Soap Request function",
			NewRequestFunction:   fakeNewRequest,
			XMLUnmarshalFunction: xmlUnmarshal,
			ExpectedContinents:   nil,
		}, {
			Name:                 "Throw an error on XML Unmarshal function",
			NewRequestFunction:   httpNewRequest,
			XMLUnmarshalFunction: fakeXMLUnmarshal,
			ExpectedContinents:   nil,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		httpNewRequest = tc.NewRequestFunction
		defer restoreNewRequest(httpNewRequest)

		xmlUnmarshal = tc.XMLUnmarshalFunction
		defer restoreXMLUnmarshal(xmlUnmarshal)

		continents, _ := GetContinentConsumer().List()
		assert.Equal(t, tc.ExpectedContinents, continents)
	}
}
