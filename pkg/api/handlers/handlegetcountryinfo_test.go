package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/consumer"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/country"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"
	"github.com/stretchr/testify/assert"
)

func mockCountryInfoFunc(countryCode string) (*country.CountryInfo, error) {
	countryInfo := &country.CountryInfo{
		Capital: "Brasilia",
		Currency: country.Currency{
			Code: "BRL",
			Name: "Brazil Real",
		},
		Flag:      "http://www.oorsprong.org/WebSamples.CountryInfo/Flags/Brazil.jpg",
		PhoneCode: 55,
	}

	return countryInfo, nil
}

func mockCountryInfoThrowFunc(countryCode string) (*country.CountryInfo, error) {
	return nil, errs.ErrRepoMockAction
}

func TestHandleGetCountryInfo(t *testing.T) {
	testCases := []struct {
		Name                        string
		CountryCode                 string
		CountryConsumerInfoFunction func(countryCode string) (*country.CountryInfo, error)
		MarshalFunction             func(v interface{}) ([]byte, error)
		WriteFunction               func(http.ResponseWriter, []byte) (int, error)
		ExpectedCountryInfo         *country.CountryInfo
		ExpectedStatusCode          int
	}{
		{
			Name:                        "Success handle get country info",
			CountryCode:                 "BR",
			CountryConsumerInfoFunction: mockCountryInfoFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			ExpectedCountryInfo: &country.CountryInfo{
				Capital: "Brasilia",
				Currency: country.Currency{
					Code: "BRL",
					Name: "Brazil Real",
				},
				Flag:      "http://www.oorsprong.org/WebSamples.CountryInfo/Flags/Brazil.jpg",
				PhoneCode: 55,
			},
			ExpectedStatusCode: 200,
		}, {
			Name:                        "Throwing error list continents function",
			CountryCode:                 "BR",
			CountryConsumerInfoFunction: mockCountryInfoThrowFunc,
			MarshalFunction:             fakeMarshal,
			WriteFunction:               write,
			ExpectedStatusCode:          500,
		}, {
			Name:                        "Throwing error marshal function",
			CountryCode:                 "BR",
			CountryConsumerInfoFunction: mockCountryInfoFunc,
			MarshalFunction:             fakeMarshal,
			WriteFunction:               write,
			ExpectedStatusCode:          500,
		}, {
			Name:                        "Throwing error on write function",
			CountryCode:                 "BR",
			CountryConsumerInfoFunction: mockCountryInfoFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               fakeWrite,
			ExpectedStatusCode:          500,
		}, {
			Name:                        "Throwing error validating country code",
			CountryCode:                 "",
			CountryConsumerInfoFunction: mockCountryInfoFunc,
			MarshalFunction:             jsonMarshal,
			WriteFunction:               write,
			ExpectedStatusCode:          404,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		consumer.SetCountryConsumer(consumer.MockCountryConsumer{
			InfoFunc: tc.CountryConsumerInfoFunction,
		})
		defer consumer.SetCountryConsumer(nil)

		jsonMarshal = tc.MarshalFunction
		defer restoreMarshal(jsonMarshal)

		write = tc.WriteFunction
		defer restoreWrite(write)

		req, err := http.NewRequest(http.MethodGet, "/countries/{code}", nil)
		req = mux.SetURLVars(req, map[string]string{"code": tc.CountryCode})
		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleGetCountryInfo(res, req)

		assert.Equal(t, tc.ExpectedStatusCode, res.Code)

		if res.Code == 200 {
			countryInfo := &country.CountryInfo{}
			err = json.Unmarshal(res.Body.Bytes(), &countryInfo)
			assert.NoError(t, err)

			assert.Equal(t, tc.ExpectedCountryInfo, countryInfo)
		}
	}
}
