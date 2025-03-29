package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestTerminalsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		id  string
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		want    string
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get terminal correctly",
			args{
				context.Background(),
				"term_7MgL4wea46qkRcoTZjWEH",
			},
			false,
			nil,
			testdata.GetTerminalResponse,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetTerminalResponse))
			},
		},
		{
			"get terminal, an error is returned from the server",
			args{
				context.Background(),
				"term_7MgL4wea46qkRcoTZjWEH",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			"",
			noPre,
			errorHandler,
		},
		{
			"get terminal, an error occurs when parsing json",
			args{
				context.Background(),
				"term_7MgL4wea46qkRcoTZjWEH",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			"",
			noPre,
			encodingHandler,
		},
		{
			"get terminal, invalid url when building request",
			args{
				context.Background(),
				"term_7MgL4wea46qkRcoTZjWEH",
			},
			true,
			errBadBaseURL,
			"",
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/terminals/%s", c.args.id), c.handler)

			res, m, err := tClient.Terminals.Get(c.args.ctx, c.args.id)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Terminal{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestTerminalsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *ListTerminalsOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		want    string
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"list terminals correctly",
			args{
				context.Background(),
				&ListTerminalsOptions{},
			},
			false,
			nil,
			testdata.GetTerminalResponse,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListTerminalsResponse))
			},
		},
		{
			"list terminals correctly with an access token",
			args{
				context.Background(),
				&ListTerminalsOptions{},
			},
			false,
			nil,
			testdata.GetTerminalResponse,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.ListTerminalsResponse))
			},
		},
		{
			"get terminals list, an error is returned from the server",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			"",
			noPre,
			errorHandler,
		},
		{
			"get terminals list, an error occurs when parsing json",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			"",
			noPre,
			encodingHandler,
		},
		{
			"get terminals list, invalid url when building request",
			args{
				context.Background(),
				nil,
			},
			true,
			errBadBaseURL,
			"",
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/terminals", c.handler)

			res, m, err := tClient.Terminals.List(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &TerminalList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
