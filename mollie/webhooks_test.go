package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestWebhookService_Get(t *testing.T) {
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
			"get webhooks works as expected with access token",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetWebhookExample))
			},
		},
		{
			"get webhooks returns an error when server responds with error",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"get webhook, an error returned when parsing JSON",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			setAccessToken,
			encodingHandler,
		},
		{
			"get webhooks fails when building request with invalid url",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
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
			tMux.HandleFunc("/v2/webhooks/"+c.args.webhook, c.handler)

			res, m, err := tClient.Webhooks.Get(c.args.ctx, c.args.webhook)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Webhook{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestWebhookService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		wh  CreateWebhook
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
			"create webhooks works as expected with access token",
			args{
				context.Background(),
				CreateWebhook{
					URL: "https://example.com/webhook",
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
				_, _ = w.Write([]byte(testdata.CreateWebhookExample))
			},
		},
		{
			"create webhooks returns an error when server responds with error",
			args{
				context.Background(),
				CreateWebhook{
					URL: "https://example.com/webhook",
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"create webhook, an error returned when parsing JSON",
			args{
				context.Background(),
				CreateWebhook{
					URL: "https://example.com/webhook",
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			setAccessToken,
			encodingHandler,
		},
		{
			"create webhooks fails when building request with invalid url",
			args{
				context.Background(),
				CreateWebhook{
					URL: "https://example.com/webhook",
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
			tMux.HandleFunc("/v2/webhooks", c.handler)

			res, m, err := tClient.Webhooks.Create(c.args.ctx, c.args.wh)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Webhook{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestWebhookService_Delete(t *testing.T) {
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
			"delete webhooks works as expected with access token",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "DELETE")

				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"delete webhooks returns an error when server responds with error",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"delete webhooks fails when building request with invalid url",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
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
			tMux.HandleFunc("/v2/webhooks/"+c.args.webhook, c.handler)

			res, err := tClient.Webhooks.Delete(c.args.ctx, c.args.webhook)
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

func TestWebhookService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		params *WebhooksListOptions
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
			"list webhooks works as expected with access token",
			args{
				context.Background(),
				&WebhooksListOptions{},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListWebhooksExample))
			},
		},
		{
			"list webhooks returns an error when server responds with error",
			args{
				context.Background(),
				&WebhooksListOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"list webhooks, an error returned when parsing JSON",
			args{
				context.Background(),
				&WebhooksListOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			setAccessToken,
			encodingHandler,
		},
		{
			"list webhooks fails when building request with invalid url",
			args{
				context.Background(),
				&WebhooksListOptions{},
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
			tMux.HandleFunc("/v2/webhooks", c.handler)

			res, m, err := tClient.Webhooks.List(c.args.ctx, c.args.params)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &WebhookList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestWebhookService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		webhook string
		uw      UpdateWebhook
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
			"update webhooks works as expected with access token",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
				UpdateWebhook{
					URL: "https://example.com/updated-webhook",
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetWebhookExample))
			},
		},
		{
			"update webhooks returns an error when server responds with error",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
				UpdateWebhook{
					URL: "https://example.com/updated-webhook",
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"update webhook, an error returned when parsing JSON",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
				UpdateWebhook{
					URL: "https://example.com/updated-webhook",
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			setAccessToken,
			encodingHandler,
		},
		{
			"update webhooks fails when building request with invalid url",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
				UpdateWebhook{
					URL: "https://example.com/updated-webhook",
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
			tMux.HandleFunc("/v2/webhooks/"+c.args.webhook, c.handler)

			res, m, err := tClient.Webhooks.Update(c.args.ctx, c.args.webhook, c.args.uw)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Webhook{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestWebhookService_Test(t *testing.T) {
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
			"test webhooks works as expected with access token",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")

				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"test webhooks returns an error when server responds with error",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"test webhooks fails when building request with invalid url",
			args{
				context.Background(),
				"hook_B2EyhTH5N4KWUnoYPcgiH",
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
			tMux.HandleFunc("/v2/webhooks/"+c.args.webhook+"/ping", c.handler)

			res, err := tClient.Webhooks.Test(c.args.ctx, c.args.webhook)
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
