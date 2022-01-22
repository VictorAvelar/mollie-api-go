package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type refundsServiceTest struct{ suite.Suite }

func (rs *refundsServiceTest) SetupSuite() { setEnv() }

func (rs *refundsServiceTest) TearDownSuite() { unsetEnv() }

func (rs *refundsServiceTest) TestRefundsService_Get() {
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
				testHeader(rs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(rs.T(), r, "GET")
				testQuery(rs.T(), r, "embed=profile&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		rs.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds/%s", c.args.payment, c.args.refund), c.handler)

			res, m, err := tClient.Refunds.Get(c.args.ctx, c.args.payment, c.args.refund, c.args.options)
			if c.wantErr {
				rs.NotNil(err)
				rs.EqualError(err, c.err.Error())
			} else {
				rs.Nil(err)
				rs.IsType(&Refund{}, m)
				rs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (rs *refundsServiceTest) TestRefundsService_Create() {
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
				testHeader(rs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(rs.T(), r, "POST")
				testQuery(rs.T(), r, "embed=profile&testmode=true")

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
			setAccesstoken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(rs.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(rs.T(), r, "POST")
				testQuery(rs.T(), r, "testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		rs.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds", c.args.payment), c.handler)

			res, m, err := tClient.Refunds.Create(c.args.ctx, c.args.payment, c.args.refund, c.args.options)
			if c.wantErr {
				rs.NotNil(err)
				rs.EqualError(err, c.err.Error())
			} else {
				rs.Nil(err)
				rs.IsType(&Refund{}, m)
				rs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (rs *refundsServiceTest) TestRefundsService_Cancel() {
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
				testHeader(rs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(rs.T(), r, "DELETE")
				testQuery(rs.T(), r, "testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		rs.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds/%s", c.args.payment, c.args.refund), c.handler)

			res, err := tClient.Refunds.Cancel(c.args.ctx, c.args.payment, c.args.refund)
			if c.wantErr {
				rs.NotNil(err)
				rs.EqualError(err, c.err.Error())
			} else {
				rs.Nil(err)
				rs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (rs *refundsServiceTest) TestRefundsService_List() {
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
				testHeader(rs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(rs.T(), r, "GET")
				testQuery(rs.T(), r, "limit=10&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		rs.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/refunds", c.handler)

			res, m, err := tClient.Refunds.ListRefund(c.args.ctx, c.args.options)
			if c.wantErr {
				rs.NotNil(err)
				rs.EqualError(err, c.err.Error())
			} else {
				rs.Nil(err)
				rs.IsType(&RefundList{}, m)
				rs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func (rs *refundsServiceTest) TestRefundsService_ListPaynents() {
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
				testHeader(rs.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(rs.T(), r, "GET")
				testQuery(rs.T(), r, "limit=10&testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		rs.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/refunds", c.args.payment), c.handler)

			res, m, err := tClient.Refunds.ListRefundPayment(c.args.ctx, c.args.payment, c.args.options)
			if c.wantErr {
				rs.NotNil(err)
				rs.EqualError(err, c.err.Error())
			} else {
				rs.Nil(err)
				rs.IsType(&RefundList{}, m)
				rs.IsType(&http.Response{}, res.Response)
			}
		})
	}
}

func TestRefundsService(t *testing.T) {
	suite.Run(t, new(refundsServiceTest))
}
