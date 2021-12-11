package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type partnersServiceSuite struct{ suite.Suite }

func (os *partnersServiceSuite) SetupSuite() { setEnv() }

func (os *partnersServiceSuite) TearDownSuite() { unsetEnv() }

func (os *partnersServiceSuite) TestPartnerService_Get() {
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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "testmode=true")

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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "embed=organization&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/clients/%s", c.args.client), c.handler)

			res, m, err := tClient.Partners.Get(c.args.ctx, c.args.client, c.args.opts)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&PartnerClient{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (os *partnersServiceSuite) TestPartnerService_List() {
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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "testmode=true")

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
				testHeader(os.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(os.T(), r, "GET")
				testQuery(os.T(), r, "testmode=true&year=2021")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		os.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/clients", c.handler)

			res, m, err := tClient.Partners.List(c.args.ctx, c.args.opts)
			if c.wantErr {
				os.NotNil(err)
				os.EqualError(err, c.err.Error())
			} else {
				os.Nil(err)
				os.IsType(&PartnerClientList{}, m)
				os.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestPartnersService(t *testing.T) {
	suite.Run(t, new(partnersServiceSuite))
}
