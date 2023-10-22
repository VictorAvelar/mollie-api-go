package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/suite"
)

type organizationsServiceSuite struct{ suite.Suite }

func (os *organizationsServiceSuite) SetupSuite() { setEnv() }

func (os *organizationsServiceSuite) TearDownSuite() { unsetEnv() }

func (os *organizationsServiceSuite) TestOrganizationsService_Get() {
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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")

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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/organizations/%s", c.args.organization), c.handler)

			res, m, err := tClient.Organizations.Get(c.args.ctx, c.args.organization)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Organization{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (os *organizationsServiceSuite) TestOrganizationsService_GetCurrent() {
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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")

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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/organizations/me", c.handler)

			res, m, err := tClient.Organizations.GetCurrent(c.args.ctx)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&Organization{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (os *organizationsServiceSuite) TestOrganizationsService_GetPartnerStatus() {
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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")

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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/organizations/me/partner", c.handler)

			res, m, err := tClient.Organizations.GetPartnerStatus(c.args.ctx)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&OrganizationPartnerStatus{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestOrganizationsService(t *testing.T) {
	suite.Run(t, new(organizationsServiceSuite))
}
