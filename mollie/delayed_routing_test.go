package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestDelayedRoutingService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
		dr      CreateDelayedRouting
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
			"creates a delayed routing successfully",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
				dr: CreateDelayedRouting{
					Description: "Delayed routing description",
					Amount: Amount{
						Currency: "EUR",
						Value:    "10.00",
					},
					Destination: DelayedRoutingDestination{
						Type:           "organization",
						OrganizationID: "org_123456789",
					},
				},
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(testdata.GetDelayedRoutingExample))
			},
		},
		{
			"creates a delayed routing works as expected with access tokens",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
				dr: CreateDelayedRouting{
					Description: "Delayed routing description",
					Amount: Amount{
						Currency: "EUR",
						Value:    "10.00",
					},
					Destination: DelayedRoutingDestination{
						Type:           "organization",
						OrganizationID: "org_123456789",
					},
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(testdata.GetDelayedRoutingExample))
			},
		},
		{
			"create delayed routing fails with error in handler",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
				dr: CreateDelayedRouting{
					Description: "Delayed routing description",
					Amount: Amount{
						Currency: "EUR",
						Value:    "10.00",
					},
					Destination: DelayedRoutingDestination{
						Type:           "organization",
						OrganizationID: "org_123456789",
					},
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create delayed routing, an error occurs when parsing json",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
				dr: CreateDelayedRouting{
					Description: "Delayed routing description",
					Amount: Amount{
						Currency: "EUR",
						Value:    "10.00",
					},
					Destination: DelayedRoutingDestination{
						Type:           "organization",
						OrganizationID: "org_123456789",
					},
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create delayed routing, invalid url when building request",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
				dr: CreateDelayedRouting{
					Description: "Delayed routing description",
					Amount: Amount{
						Currency: "EUR",
						Value:    "10.00",
					},
					Destination: DelayedRoutingDestination{
						Type:           "organization",
						OrganizationID: "org_123456789",
					},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/routes", c.args.payment), c.handler)

			res, m, err := tClient.DelayedRouting.Create(c.args.ctx, c.args.payment, c.args.dr)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Route{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestDelayedRoutingService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		payment string
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
			"lists delayed routings successfully",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListDelayedRoutingsExample))
			},
		},
		{
			"lists delayed routings works as expected with access tokens",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListDelayedRoutingsExample))
			},
		},
		{
			"list delayed routings fails with error in handler",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"list delayed routings, an error occurs when parsing json",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list delayed routings, invalid url when building request",
			args{
				ctx:     context.Background(),
				payment: "tr_123456789",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/payments/%s/routes", c.args.payment), c.handler)

			res, m, err := tClient.DelayedRouting.List(c.args.ctx, c.args.payment)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentRoutesList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
