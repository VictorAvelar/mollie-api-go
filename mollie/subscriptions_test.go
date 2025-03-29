package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestSubscriptionsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx          context.Context
		customer     string
		subscription string
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
			"get subscription works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
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
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"get subscription, an error is returned from the server",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get subscription, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get subscription, invalid url when building request",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions/%s", c.args.customer, c.args.subscription), c.handler)

			res, m, err := tClient.Subscriptions.Get(c.args.ctx, c.args.customer, c.args.subscription)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Subscription{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSubscriptionsService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx          context.Context
		customer     string
		subscription CreateSubscription
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
			"create subscription works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				CreateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
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
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"create subscription with access tokens works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				CreateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
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
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"create subscription, an error is returned from the server",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				CreateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create subscription, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				CreateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create subscription, invalid url when building request",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				CreateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions", c.args.customer), c.handler)

			res, m, err := tClient.Subscriptions.Create(c.args.ctx, c.args.customer, c.args.subscription)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Subscription{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSubscriptionsService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx          context.Context
		customer     string
		sid          string
		subscription UpdateSubscription
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
			"update subscription works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				UpdateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "PATCH")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"update subscription with access tokens works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				UpdateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"update subscription, an error is returned from the server",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				UpdateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"update subscription, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				UpdateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update subscription, invalid url when building request",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				UpdateSubscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions/%s", c.args.customer, c.args.sid), c.handler)

			res, m, err := tClient.Subscriptions.Update(c.args.ctx, c.args.customer, c.args.sid, c.args.subscription)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Subscription{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSubscriptionsService_Cancel(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		customer string
		sid      string
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
			"delete subscription works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
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
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"delete subscription with access tokens works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "DELETE")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
			},
		},
		{
			"delete subscription, an error is returned from the server",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"delete subscription, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"delete subscription, invalid url when building request",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions/%s", c.args.customer, c.args.sid), c.handler)

			res, m, err := tClient.Subscriptions.Cancel(c.args.ctx, c.args.customer, c.args.sid)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Subscription{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSubscriptionsService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
		customer string
		options  *ListSubscriptionsOptions
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
			"list subscriptions works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
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
				_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
			},
		},
		{
			"list subscriptions with options works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				&ListSubscriptionsOptions{
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
				_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
			},
		},
		{
			"list subscriptions, an error is returned from the server",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"list subscriptions, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list subscriptions, invalid url when building request",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions", c.args.customer), c.handler)

			res, m, err := tClient.Subscriptions.List(c.args.ctx, c.args.customer, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SubscriptionsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSubscriptionsService_All(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *ListSubscriptionsOptions
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
			"list subscriptions works as expected.",
			args{
				context.Background(),
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
				_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
			},
		},
		{
			"list subscriptions with options works as expected.",
			args{
				context.Background(),
				&ListSubscriptionsOptions{
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
				_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
			},
		},
		{
			"list subscriptions, an error is returned from the server",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"list subscriptions, an error occurs when parsing json",
			args{
				context.Background(),
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list subscriptions, invalid url when building request",
			args{
				context.Background(),
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
			tMux.HandleFunc("/v2/subscriptions", c.handler)

			res, m, err := tClient.Subscriptions.All(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SubscriptionsList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSubscriptionsService_GetPayments(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx          context.Context
		customer     string
		subscription string
		options      *ListSubscriptionsOptions
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
			"get subscription payments works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
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
				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"get subscription payments with options works as expected.",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				&ListSubscriptionsOptions{
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
				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"get subscription payments, an error is returned from the server",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get subscription payments, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get subscription payments, invalid url when building request",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				"sub_rVKGtNd6s3",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions/%s/payments", c.args.customer, c.args.subscription), c.handler)

			res, m, err := tClient.Subscriptions.ListPayments(c.args.ctx, c.args.customer, c.args.subscription, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
