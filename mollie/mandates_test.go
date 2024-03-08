package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestMandatesService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
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

		t.Run(c.name, func(t *testing.T) {
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Mandate{}, m)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestMandatesService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		mandate  CreateMandate
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
				CreateMandate{
					Method: PayPal,
				},
				"cst_4qqhO89gsT",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")
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
				CreateMandate{
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
				CreateMandate{
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
				CreateMandate{
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

		t.Run(c.name, func(t *testing.T) {
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Mandate{}, m)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestMandatesService_Revoke(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "DELETE")
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

		t.Run(c.name, func(t *testing.T) {
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestMandatesService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
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
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
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

		t.Run(c.name, func(t *testing.T) {
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
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &MandatesList{}, m)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
