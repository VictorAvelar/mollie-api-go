package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestPartnerService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		client string
		opts   *GetPartnerClientOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get partner client works as expected.",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPartnerClientResponse))
			},
		},
		{
			"get partner client with options works as expected.",
			args{
				context.Background(),
				"org_1337",
				&GetPartnerClientOptions{
					Embed: "organization",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "embed=organization")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPartnerClientResponse))
			},
		},
		{
			"get partner client, an error is returned from the server",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get partner client, an error occurs when parsing json",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get partner client, invalid url when building request",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/clients/%s", c.args.client), c.handler)

			res, m, err := tClient.Clients.Get(c.args.ctx, c.args.client, c.args.opts)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PartnerClient{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestPartnerService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		client string
		opts   *ListPartnerClientsOptions
	}

	cases := []struct {
		name    string
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"list partner client works as expected.",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPartnerClientResponse))
			},
		},
		{
			"list partner client with options works as expected.",
			args{
				context.Background(),
				"org_1337",
				&ListPartnerClientsOptions{
					Year: 2021,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "year=2021")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetPartnerClientResponse))
			},
		},
		{
			"list partner client, an error is returned from the server",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list partner client, an error occurs when parsing json",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list partner client, invalid url when building request",
			args{
				context.Background(),
				"org_1337",
				nil,
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/clients", c.handler)

			res, m, err := tClient.Clients.List(c.args.ctx, c.args.opts)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PartnerClientList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
