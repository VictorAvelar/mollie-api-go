package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestRefundsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		refund  string
		options *RefundOptions
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
			"get refund works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
				&RefundOptions{
					Embed: "profile",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "embed=profile")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetRefundResponse))
			},
		},
		{
			"get refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get refund, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds/%s", c.args.payment, c.args.refund), c.handler)

			res, m, err := tClient.Refunds.Get(c.args.ctx, c.args.payment, c.args.refund, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Refund{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		refund  Refund
		options *RefundOptions
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
			"create refund works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Refund{
					Amount: &Amount{
						Currency: "USD",
						Value:    "100.00",
					},
				},
				&RefundOptions{
					Embed: "profile",
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")
				testQuery(t, r, "embed=profile")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetRefundResponse))
			},
		},
		{
			"create refund with access token works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Refund{
					Amount: &Amount{
						Currency: "USD",
						Value:    "100.00",
					},
				},
				nil,
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetRefundResponse))
			},
		},
		{
			"create refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Refund{
					Amount: &Amount{
						Currency: "USD",
						Value:    "100.00",
					},
				},
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"create refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Refund{
					Amount: &Amount{
						Currency: "USD",
						Value:    "100.00",
					},
				},
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create refund, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				Refund{
					Amount: &Amount{
						Currency: "USD",
						Value:    "100.00",
					},
				},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds", c.args.payment), c.handler)

			res, m, err := tClient.Refunds.Create(c.args.ctx, c.args.payment, c.args.refund, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Refund{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_Cancel(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		refund  string
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
			"get settlement works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
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
				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"get settlement, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get settlement, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds/%s", c.args.payment, c.args.refund), c.handler)

			res, err := tClient.Refunds.Cancel(c.args.ctx, c.args.payment, c.args.refund)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		options *ListRefundOptions
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
			"list refund works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				&ListRefundOptions{
					Limit: 10,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=10")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetRefundListResponse))
			},
		},
		{
			"list refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"list refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list refund, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
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
			tMux.HandleFunc("/v2/refunds", c.handler)

			res, m, err := tClient.Refunds.ListRefund(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &RefundList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService_ListPaynents(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		refund  string
		options *ListRefundOptions
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
			"get refund works as expected.",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
				&ListRefundOptions{
					Limit: 10,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=10")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetRefundListResponse))
			},
		},
		{
			"get refund, an error is returned from the server",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request."),
			noPre,
			errorHandler,
		},
		{
			"get refund, an error occurs when parsing json",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get refund, invalid url when building request",
			args{
				context.Background(),
				"tr_WDqYK6vllg",
				"re_4qqhO89gsT",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds", c.args.payment), c.handler)

			res, m, err := tClient.Refunds.ListRefundPayment(c.args.ctx, c.args.payment, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &RefundList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
