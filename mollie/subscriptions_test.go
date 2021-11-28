package mollie

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/suite"
)

type subscriptionsServiceSuite struct{ suite.Suite }

func (ps *subscriptionsServiceSuite) SetupSuite() { setEnv() }

func (ps *subscriptionsServiceSuite) TearDownSuite() { unsetEnv() }

func (ps *subscriptionsServiceSuite) TestSubscriptionsService_Get() {
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "GET")
				testQuery(ps.T(), r, "testmode=true")

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
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions/%s", c.args.customer, c.args.subscription), c.handler)

			m, err := tClient.Subscriptions.Get(c.args.ctx, c.args.customer, c.args.subscription)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Subscription{}, m)
			}
		})
	}
}

func (ps *subscriptionsServiceSuite) TestSubscriptionsService_Create() {
	type args struct {
		ctx          context.Context
		customer     string
		subscription *Subscription
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
				&Subscription{
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
				testHeader(ps.T(), r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

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
				&Subscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_token_test")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(ps.T(), r, AuthHeader, "Bearer access_token_test")
				testMethod(ps.T(), r, "POST")
				testQuery(ps.T(), r, "testmode=true")

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
				&Subscription{
					Amount: &Amount{
						Currency: "EUR",
						Value:    "100.00",
					},
					Times: 12,
				},
			},
			true,
			fmt.Errorf("response failed with status 500 Internal Server Error\npayload: "),
			noPre,
			errorHandler,
		},
		{
			"create subscription, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_stTC2WHAuS",
				&Subscription{
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
				&Subscription{
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

		ps.T().Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/subscriptions", c.args.customer), c.handler)

			m, err := tClient.Subscriptions.Create(c.args.ctx, c.args.customer, c.args.subscription)
			if c.wantErr {
				ps.NotNil(err)
				ps.EqualError(err, c.err.Error())
			} else {
				ps.Nil(err)
				ps.IsType(&Subscription{}, m)
			}
		})
	}
}

func TestSubscriptionService(t *testing.T) {
	suite.Run(t, new(subscriptionsServiceSuite))
}

func TestSubscriptionsService_Update(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodPatch)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
	})

	s := Subscription{
		Amount: &Amount{
			Currency: "EUR",
			Value:    "11.99",
		},
		Interval: "12 months",
	}

	sub, err := tClient.Subscriptions.Update(context.TODO(), cID, sID, &s)
	if err != nil {
		t.Error(err)
	}

	if sub.ID != sID {
		t.Errorf("unexpected response: got %v, want %v", sub.ID, sID)
	}
}

func TestSubscriptionsService_Delete(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodDelete)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetSubscriptionResponse))
	})

	sub, err := tClient.Subscriptions.Delete(context.TODO(), cID, sID)
	if err != nil {
		t.Error(err)
	}

	if sub.ID != sID {
		t.Errorf("unexpected response: got %v, want %v", sub.ID, sID)
	}
}

func TestSubscriptionsService_List(t *testing.T) {
	setup()
	defer teardown()

	cID := "cst_stTC2WHAuS"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions", cID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
	})

	sub, err := tClient.Subscriptions.List(context.TODO(), cID, &SubscriptionListOptions{Limit: 6})
	if err != nil {
		t.Error(err)
	}

	if sub.Count != 3 {
		t.Errorf("unexpected response: got %v, want 3", sub.Count)
	}
}

func TestSubscriptionsService_All(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc("/v2/subscriptions", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListAllSubscriptionsResponse))
	})

	sub, err := tClient.Subscriptions.All(context.TODO(), &SubscriptionListOptions{Limit: 6})
	if err != nil {
		t.Error(err)
	}

	if sub.Count != 3 {
		t.Errorf("unexpected response: got %v, want 3", sub.Count)
	}
}

func TestSubscriptionsService_GetPayments(t *testing.T) {
	setup()
	defer teardown()
	cID := "cst_stTC2WHAuS"
	sID := "sub_rVKGtNd6s3"
	u := fmt.Sprintf("/v2/customers/%s/subscriptions/%s/payments", cID, sID)
	_ = tClient.WithAuthenticationValue("test_token")
	tMux.HandleFunc(u, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, http.MethodGet)

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
	})
	sub, err := tClient.Subscriptions.GetPayments(context.TODO(), cID, sID, &SubscriptionListOptions{Limit: 6})
	if err != nil {
		t.Error(err)
	}
	if sub.Count != 5 {
		t.Errorf("unexpected response: got %v, want 5", sub.Count)
	}
}

func TestSubscriptionsService_HTTPErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/", errorHandler)
	tests := forceSubscriptionsErrors()

	for _, tt := range tests {
		if !strings.Contains(tt.Error(), "Internal Server Error") {
			t.Error(tt)
		}
	}
}

func TestSubscriptionsService_JsonEncodingErrors(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/v2/", encodingHandler)

	tests := forceSubscriptionsErrors()

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestSubscriptionsService_NewAPIRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u

	tMux.HandleFunc("/v2/", errorHandler)

	tests := forceSubscriptionsErrors()

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}

func forceSubscriptionsErrors() []error {
	_, gerr := tClient.Subscriptions.Get(context.TODO(), "6s7d89a", "sd9af8ds")
	_, cerr := tClient.Subscriptions.Create(context.TODO(), "6s7d89a", &Subscription{Mode: LiveMode})
	_, uerr := tClient.Subscriptions.Update(context.TODO(), "6s7d89a", "0re9w0e", &Subscription{Mode: LiveMode})
	_, derr := tClient.Subscriptions.Delete(context.TODO(), "6s7d89a", "0re9w0e")
	_, lerr := tClient.Subscriptions.List(context.TODO(), "6s7d89a", nil)
	_, aerr := tClient.Subscriptions.All(context.TODO(), nil)
	_, perr := tClient.Subscriptions.GetPayments(context.TODO(), "6s7d89a", "0re9w0e", nil)

	return []error{gerr, cerr, uerr, derr, lerr, aerr, perr}
}
