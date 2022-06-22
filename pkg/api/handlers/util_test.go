package handlers

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidateCountryCode(t *testing.T) {
	testCases := []struct {
		Name            string
		CountryCode     string
		HttpGetFunction func(url string) (resp *http.Response, err error)
		ExpectedError   bool
	}{
		{
			Name:            "Should validate as true",
			CountryCode:     "BR",
			HttpGetFunction: httpGet,
			ExpectedError:   false,
		}, {
			Name:            "Should throw an error when validate",
			CountryCode:     "ll",
			HttpGetFunction: httpGet,
			ExpectedError:   true,
		}, {
			Name:            "Should throw an error when country code is empty",
			CountryCode:     "",
			HttpGetFunction: httpGet,
			ExpectedError:   true,
		}, {
			Name:            "Should throw an error on http get function",
			CountryCode:     "US",
			HttpGetFunction: fakeHttpGet,
			ExpectedError:   true,
		},
	}

	for _, tc := range testCases {
		t.Logf(tc.Name)

		httpGet = tc.HttpGetFunction
		defer restoreHttpGet(httpGet)

		err := validateCountryCode(tc.CountryCode)

		if tc.ExpectedError {
			assert.NotNil(t, err)
		} else {
			assert.Nil(t, err)
		}
	}
}
