package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/suite"
)

type mandateServiceSuite struct{ suite.Suite }

func (ms *mandateServiceSuite) SetupSuite() { setEnv() }

func (ms *mandateServiceSuite) TearDownSuite() { unsetEnv() }

func (ms *mandateServiceSuite) TestMandatesService_Get() {
	type args struct {
		ctx      context.Context
		mandate  string
		customer string
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
			"get mandates works as expected.",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetMandateResponse))
			},
		},
		{
			"get mandate, an error is returned from the server",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get mandate, an error occurs when parsing json",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get mandate, invalid url when building request",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
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

		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/customers/%s/mandates/%s",
					c.args.customer,
					c.args.mandate,
				),
				c.handler,
			)

			res, m, err := tClient.Mandates.Get(c.args.ctx, c.args.customer, c.args.mandate)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&Mandate{}, m)
				ms.Same(c.args.ctx, res.Request.Context())
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ms *mandateServiceSuite) TestMandatesService_Create() {
	type args struct {
		ctx      context.Context
		mandate  Mandate
		customer string
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
			"create mandates works as expected.",
			args{
				context.Background(),
				Mandate{
					Method: PayPal,
				},
				"cst_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "POST")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(testdata.CreateMandateResponse))
			},
		},
		{
			"create mandate, an error is returned from the server",
			args{
				context.Background(),
				Mandate{
					Method: PayPal,
				},
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"create mandate, an error occurs when parsing json",
			args{
				context.Background(),
				Mandate{
					Method: PayPal,
				},
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create mandate, invalid url when building request",
			args{
				context.Background(),
				Mandate{
					Method: PayPal,
				},
				"cst_4qqhO89gsT",
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

		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/customers/%s/mandates",
					c.args.customer,
				),
				c.handler,
			)

			res, m, err := tClient.Mandates.Create(c.args.ctx, c.args.customer, c.args.mandate)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&Mandate{}, m)
				ms.Same(c.args.ctx, res.Request.Context())
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ms *mandateServiceSuite) TestMandatesService_Revoke() {
	type args struct {
		ctx      context.Context
		mandate  string
		customer string
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
			"revoke mandates works as expected.",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "DELETE")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"revoke mandate, an error is returned from the server",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"revoke mandate, invalid url when building request",
			args{
				context.Background(),
				"mdt_h3gAaD5zP",
				"cst_4qqhO89gsT",
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

		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/customers/%s/mandates/%s",
					c.args.customer,
					c.args.mandate,
				),
				c.handler,
			)

			res, err := tClient.Mandates.Revoke(c.args.ctx, c.args.customer, c.args.mandate)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.Same(c.args.ctx, res.Request.Context())
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (ms *mandateServiceSuite) TestMandatesService_List() {
	type args struct {
		ctx      context.Context
		options  *MandatesListOptions
		customer string
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
			"list mandates works as expected.",
			args{
				context.Background(),
				nil,
				"cst_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetMandateResponse))
			},
		},
		{
			"list mandates with options works as expected.",
			args{
				context.Background(),
				&MandatesListOptions{
					Limit: 10,
				},
				"cst_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ms.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ms.T(), r, "GET")
				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetMandateResponse))
			},
		},
		{
			"list mandates, an error is returned from the server",
			args{
				context.Background(),
				nil,
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list mandates, an error occurs when parsing json",
			args{
				context.Background(),
				nil,
				"cst_4qqhO89gsT",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list mandates, invalid url when building request",
			args{
				context.Background(),
				nil,
				"cst_4qqhO89gsT",
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

		ms.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(
				fmt.Sprintf(
					"/v2/customers/%s/mandates",
					c.args.customer,
				),
				c.handler,
			)

			res, m, err := tClient.Mandates.List(c.args.ctx, c.args.customer, c.args.options)
			if c.wantErr {
				ms.NotNil(err)
				ms.EqualError(err, c.err.Error())
			} else {
				ms.Nil(err)
				ms.IsType(&MandatesList{}, m)
				ms.Same(c.args.ctx, res.Request.Context())
				ms.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestMandatesService(t *testing.T) {
	suite.Run(t, new(mandateServiceSuite))
}
