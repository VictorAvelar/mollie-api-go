package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestWebhookEventsService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		webhook string
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
			"get webhook events works as expected with access token",
			args{
				context.Background(),
				"event_GvJ8WHrp5isUdRub9CJyH",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetWebhookEventExample))
			},
		},
		{
			"get webhook events returns an error when server responds with error",
			args{
				context.Background(),
				"event_GvJ8WHrp5isUdRub9CJyH",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"get webhook events, an error returned when parsing JSON",
			args{
				context.Background(),
				"event_GvJ8WHrp5isUdRub9CJyH",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			setAccessToken,
			encodingHandler,
		},
		{
			"get webhook events fails when building request with invalid url",
			args{
				context.Background(),
				"event_GvJ8WHrp5isUdRub9CJyH",
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
			tMux.HandleFunc("/v2/events/"+c.args.webhook, c.handler)

			res, m, err := tClient.WebhookEvents.Get(c.args.ctx, c.args.webhook)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &WebhookEvent{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
