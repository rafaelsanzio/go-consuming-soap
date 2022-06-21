package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHandleOk(t *testing.T) {
	goodReq := httptest.NewRequest(http.MethodGet, "/ok", nil)

	testCases := []struct {
		Name               string
		Request            *http.Request
		MarshalFunction    func(v interface{}) ([]byte, error)
		WriteFunction      func(http.ResponseWriter, []byte) (int, error)
		ExpectedStatusCode int
	}{
		{
			Name:               "Should return 200 if successful",
			Request:            goodReq,
			MarshalFunction:    jsonMarshal,
			WriteFunction:      write,
			ExpectedStatusCode: 200,
		}, {
			Name:               "Error marshaling json",
			Request:            goodReq,
			MarshalFunction:    fakeMarshal,
			WriteFunction:      write,
			ExpectedStatusCode: 500,
		},
		{
			Name:               "Error writing to response writer",
			Request:            goodReq,
			MarshalFunction:    jsonMarshal,
			WriteFunction:      fakeWrite,
			ExpectedStatusCode: 500,
		},
	}

	for _, tc := range testCases {
		t.Log(tc.Name)

		jsonMarshal = tc.MarshalFunction
		defer restoreMarshal(jsonMarshal)

		write = tc.WriteFunction
		defer restoreWrite(write)

		w := httptest.NewRecorder()

		HandleOK(w, tc.Request)
		assert.Equal(t, tc.ExpectedStatusCode, w.Code)
	}
}
