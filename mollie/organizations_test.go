package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestOrganizationsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx          context.Context
		organization string
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
				"org_12345678",
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
				_, _ = w.Write([]byte(testdata.GetOrganizationResponse))
			},
		},
		{
			"get organization, an error is returned from the server",
			args{
				context.Background(),
				"org_12345678",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get organization, an error occurs when parsing json",
			args{
				context.Background(),
				"org_12345678",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get organization, invalid url when building request",
			args{
				context.Background(),
				"org_12345678",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/organizations/%s", c.args.organization), c.handler)

			res, m, err := tClient.Organizations.Get(c.args.ctx, c.args.organization)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Organization{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrganizationsService_GetCurrent(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
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
			"get current organization works as expected.",
			args{
				context.Background(),
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
				_, _ = w.Write([]byte(testdata.GetOrganizationResponse))
			},
		},
		{
			"get current organization, an error is returned from the server",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get current organization, an error occurs when parsing json",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get current organization, invalid url when building request",
			args{
				context.Background(),
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
			tMux.HandleFunc("/v2/organizations/me", c.handler)

			res, m, err := tClient.Organizations.GetCurrent(c.args.ctx)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Organization{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestOrganizationsService_GetPartnerStatus(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
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
			"get current partner status works as expected.",
			args{
				context.Background(),
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
				_, _ = w.Write([]byte(testdata.GetPartnerStatusResponse))
			},
		},
		{
			"get current partner status, an error is returned from the server",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get current partner status, an error occurs when parsing json",
			args{
				context.Background(),
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get current partner status, invalid url when building request",
			args{
				context.Background(),
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
			tMux.HandleFunc("/v2/organizations/me/partner", c.handler)

			res, m, err := tClient.Organizations.GetPartnerStatus(c.args.ctx)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &OrganizationPartnerStatus{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
