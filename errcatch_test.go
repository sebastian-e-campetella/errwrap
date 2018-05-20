package errwrap_test

import (
	"errors"
	"errwrap"
	"fmt"
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
		Method        string
	}{
		{
			ExpectedError: nil,
			Url:           "https://www.mercadolibre.com.ar",
			Verb:          "GET",
			StatusCode:    200,
			Responder:     httpmock.NewStringResponder(200, ""),
			Method:        "Catch",
		},
		{
			ExpectedError: errors.New("Not Found"),
			Url:           "https://iooooo.com",
			Verb:          "GET",
			StatusCode:    400,
			Responder:     httpmock.NewStringResponder(400, ""),
			Method:        "CatchWrapper",
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

		switch tc.Verb {
		case "Catch":
			ew.Catch(resp, err)
		case "CatchWrapper":
			ew.CatchWrapper(ew, func() {
				fmt.Println("hola")
			})
		}

		if ew.Any != nil {
			assert.Equal(t, tc.StatusCode, (ew.Any.(*http.Response)).StatusCode, err)
		}
		httpmock.Reset()
	}
}
