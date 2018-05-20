package errwrap_test

import (
	"errors"
	"errwrap"
	"net/http"
	"testing"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

func TestCatch(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	tt := []struct {
		ExpectedError error
		Url           string
		Verb          string
		StatusCode    int
		Responder     httpmock.Responder
	}{
		{
			ExpectedError: nil,
			Url:           "https://www.mercadolibre.com.ar",
			Verb:          "GET",
			StatusCode:    200,
			Responder:     httpmock.NewStringResponder(200, ""),
		},
		{
			ExpectedError: errors.New("Not Found"),
			Url:           "https://iooooo.com",
			Verb:          "GET",
			StatusCode:    400,
			Responder:     httpmock.NewStringResponder(400, ""),
		},
	}

	for _, tc := range tt {
		var ew errwrap.ErrorWrapper
		httpmock.RegisterResponder(tc.Verb, tc.Url, tc.Responder)

		var resp *http.Response
		var err error

		switch tc.Verb {
		case "GET":
			resp, err = http.Get(tc.Url)
		case "POST":
			resp, err = http.Post(tc.Url, "appliction/json", nil)
		}

		assert.Equal(t, tc.StatusCode, (ew.Catch(resp, err).(*http.Response)).StatusCode, err)
		httpmock.Reset()
	}
}
