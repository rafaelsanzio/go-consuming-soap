package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rafaelsanzio/go-consuming-soap/pkg/consumer"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/continent"
	"github.com/rafaelsanzio/go-consuming-soap/pkg/errs"

	"github.com/stretchr/testify/assert"
)

func mockListContinentsFunc() ([]continent.Continent, error) {
	continents := []continent.Continent{
		{Code: "AF", Name: "Africa"},
		{Code: "AN", Name: "Antarctica"},
		{Code: "AS", Name: "Asia"},
		{Code: "EU", Name: "Europe"},
		{Code: "OC", Name: "Ocenania"},
		{Code: "AM", Name: "The Americas"},
	}

	return continents, nil
}

func mockListContinentsThrowFunc() ([]continent.Continent, error) {
	return nil, errs.ErrRepoMockAction
}

func TestHandleListContinents(t *testing.T) {
	testCases := []struct {
		Name                          string
		ContinentConsumerListFunction func() ([]continent.Continent, error)
		MarshalFunction               func(v interface{}) ([]byte, error)
		WriteFunction                 func(http.ResponseWriter, []byte) (int, error)
		ExpectedContinents            []continent.Continent
		ExpectedStatusCode            int
	}{
		{
			Name:                          "Success handle list continents",
			ContinentConsumerListFunction: mockListContinentsFunc,
			MarshalFunction:               jsonMarshal,
			WriteFunction:                 write,
			ExpectedContinents: []continent.Continent{
				{Code: "AF", Name: "Africa"},
				{Code: "AN", Name: "Antarctica"},
				{Code: "AS", Name: "Asia"},
				{Code: "EU", Name: "Europe"},
				{Code: "OC", Name: "Ocenania"},
				{Code: "AM", Name: "The Americas"},
			},
			ExpectedStatusCode: 200,
		}, {
			Name:                          "Throwing error list continents function",
			ContinentConsumerListFunction: mockListContinentsThrowFunc,
			MarshalFunction:               fakeMarshal,
			WriteFunction:                 write,
			ExpectedStatusCode:            500,
		}, {
			Name:                          "Throwing error marshal function",
			ContinentConsumerListFunction: mockListContinentsFunc,
			MarshalFunction:               fakeMarshal,
			WriteFunction:                 write,
			ExpectedStatusCode:            500,
		}, {
			Name:                          "Throwing error on write function",
			ContinentConsumerListFunction: mockListContinentsFunc,
			MarshalFunction:               jsonMarshal,
			WriteFunction:                 fakeWrite,
			ExpectedStatusCode:            500,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		consumer.SetContinentConsumer(consumer.MockContinentConsumer{
			ListFunc: tc.ContinentConsumerListFunction,
		})
		defer consumer.SetContinentConsumer(nil)

		jsonMarshal = tc.MarshalFunction
		defer restoreMarshal(jsonMarshal)

		write = tc.WriteFunction
		defer restoreWrite(write)

		req, err := http.NewRequest(http.MethodGet, "/continents", nil)
		assert.NoError(t, err)
		res := httptest.NewRecorder()

		HandleListContinents(res, req)

		assert.Equal(t, tc.ExpectedStatusCode, res.Code)

		if res.Code == 200 {
			continents := []continent.Continent{}
			err = json.Unmarshal(res.Body.Bytes(), &continents)
			assert.NoError(t, err)

			assert.Equal(t, tc.ExpectedContinents, continents)
		}
	}
}
