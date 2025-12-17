package mollie

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestSalesInvoicesService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		req CreateSalesInvoice
	}

	recipient := SalesInvoiceRecipient{
		Type:  ConsumerSalesInvoiceRecipientType,
		Email: "sample@example.com",
		Address: Address{
			StreetAndNumber: "Keizersgracht 313",
			PostalCode:      "1016 EE",
			City:            "Amsterdam",
			Country:         "NL",
		},
		Locale: Dutch,
	}

	lines := []SalesInvoiceLineItem{
		{
			Description: "Product A",
			Quantity:    2,
			UnitPrice: Amount{
				Currency: "EUR",
				Value:    "50.00",
			},
			VATRate: "21.00",
		},
	}

	issuedCreateHandler := func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
		testMethod(t, r, "POST")

		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		var payload map[string]any
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		status, _ := payload["status"].(string)
		paymentDetails, hasPaymentDetails := payload["paymentDetails"].(map[string]any)

		if hasPaymentDetails {
			if source, ok := paymentDetails["source"].(string); !ok || source == "" {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(`{"status":422,"title":"Unprocessable Entity","detail":"The 'source' field is required when providing paymentDetails.","field":"paymentDetails.source"}`))
				return
			}

			if status == string(IssuedSalesInvoiceStatus) {
				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusUnprocessableEntity)
				_, _ = w.Write([]byte(`{"status":422,"title":"Unprocessable Entity","detail":"The 'paymentDetails' field is not allowed when status is 'issued'.","field":"paymentDetails"}`))
				return
			}
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateSalesInvoicesResponse))
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
			name: "create sales invoice successfully",
			args: args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              DraftSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			wantErr: false,
			pre:     noPre,
			handler: func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(testdata.CreateSalesInvoicesResponse))
			},
		},
		{
			"create sales invoice works as expected with access tokens",
			args{
				context.Background(),
				CreateSalesInvoice{
					Status:              DraftSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(testdata.CreateSalesInvoicesResponse))
			},
		},
		{
			"create sales invoices error handler",
			args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              DraftSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create sales invoice, an error occurs when parsing json",
			args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              DraftSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create sales invoice, invalid url when building request",
			args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              DraftSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
		{
			name: "create sales invoice, payment details list in response",
			args: args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              DraftSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			wantErr: false,
			err:     nil,
			pre:     noPre,
			handler: func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "POST")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
					return
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(`{
  "resource": "sales-invoice",
  "id": "invoice_4Y0eZitmBnQ6IDoMqZQKh",
  "status": "paid",
  "paymentDetails": [
    { "source": "manual", "sourceReference": "ref_1" },
    { "source": "payment", "sourceReference": "tr_2" }
  ]
}`))
			},
		},
		{
			name: "create issued sales invoice, without payment details",
			args: args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              IssuedSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
				},
			},
			wantErr: false,
			err:     nil,
			pre:     noPre,
			handler: issuedCreateHandler,
		},
		{
			name: "create issued sales invoice, payment details without source returns error",
			args: args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              IssuedSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
					PaymentDetails:      &SalesInvoicePaymentDetails{},
				},
			},
			wantErr: true,
			err: &BaseError{
				Status: http.StatusUnprocessableEntity,
				Title:  "Unprocessable Entity",
				Detail: "The 'source' field is required when providing paymentDetails.",
				Field:  "paymentDetails.source",
			},
			pre:     noPre,
			handler: issuedCreateHandler,
		},
		{
			name: "create issued sales invoice, payment details with source returns error",
			args: args{
				ctx: context.Background(),
				req: CreateSalesInvoice{
					Status:              IssuedSalesInvoiceStatus,
					RecipientIdentifier: "customer_123456789",
					Recipient:           recipient,
					Lines:               lines,
					PaymentDetails: &SalesInvoicePaymentDetails{
						Source: ManualSalesInvoiceSource,
					},
				},
			},
			wantErr: true,
			err: &BaseError{
				Status: http.StatusUnprocessableEntity,
				Title:  "Unprocessable Entity",
				Detail: "The 'paymentDetails' field is not allowed when status is 'issued'.",
				Field:  "paymentDetails",
			},
			pre:     noPre,
			handler: issuedCreateHandler,
		},
	}

	for _, c := range cases {
		setup()
		defer teardown()

		t.Run(c.name, func(t *testing.T) {
			c.pre()
			tMux.HandleFunc("/v2/sales-invoices", c.handler)

			res, m, err := tClient.SalesInvoices.Create(c.args.ctx, c.args.req)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SalesInvoice{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSalesInvoicesService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		id  string
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
			name: "get sales invoice successfully",
			args: args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			wantErr: false,
			pre:     noPre,
			handler: func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetSalesInvoicesResponse))
			},
		},
		{
			"get sales invoice works as expected with access tokens",
			args{
				context.Background(),
				"invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetSalesInvoicesResponse))
			},
		},
		{
			"get sales invoice error handler",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get sales invoice, an error occurs when parsing json",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get sales invoice, invalid url when building request",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
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
			tMux.HandleFunc("/v2/sales-invoices/"+c.args.id, c.handler)

			res, m, err := tClient.SalesInvoices.Get(c.args.ctx, c.args.id)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SalesInvoice{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSalesInvoicesService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx  context.Context
		opts ListSalesInvoicesOptions
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
			name: "list sales invoices successfully",
			args: args{
				ctx: context.Background(),
				opts: ListSalesInvoicesOptions{
					Limit: 5,
				},
			},
			wantErr: false,
			pre:     noPre,
			handler: func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListSalesInvoicesResponse))
			},
		},
		{
			"list sales invoices works as expected with access tokens",
			args{
				context.Background(),
				ListSalesInvoicesOptions{
					Limit: 5,
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "GET")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.ListSalesInvoicesResponse))
			},
		},
		{
			"list sales invoices error handler",
			args{
				ctx:  context.Background(),
				opts: ListSalesInvoicesOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"list sales invoices, an error occurs when parsing json",
			args{
				ctx:  context.Background(),
				opts: ListSalesInvoicesOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list sales invoices, invalid url when building request",
			args{
				ctx:  context.Background(),
				opts: ListSalesInvoicesOptions{},
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
			tMux.HandleFunc("/v2/sales-invoices", c.handler)

			res, m, err := tClient.SalesInvoices.List(c.args.ctx, &c.args.opts)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SalesInvoiceList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestSalesInvoicesService_Delete(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		id  string
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
			name: "delete sales invoice successfully",
			args: args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			wantErr: false,
			pre:     noPre,
			handler: func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "DELETE")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"delete sales invoice works as expected with access tokens",
			args{
				context.Background(),
				"invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "DELETE")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"delete sales invoice error handler",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"delete sales invoice, invalid url when building request",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
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
			tMux.HandleFunc("/v2/sales-invoices/"+c.args.id, c.handler)

			res, err := tClient.SalesInvoices.Delete(c.args.ctx, c.args.id)
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

func TestSalesInvoicesService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx context.Context
		id  string
		req UpdateSalesInvoice
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
			name: "update sales invoice successfully",
			args: args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
				req: UpdateSalesInvoice{
					Status: PaidSalesInvoiceStatus,
				},
			},
			wantErr: false,
			pre:     noPre,
			handler: func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
				testMethod(t, r, "PATCH")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetSalesInvoicesResponse))
			},
		},
		{
			"update sales invoice works as expected with access tokens",
			args{
				context.Background(),
				"invoice_4Y0eZitmBnQ6IDoMqZQKh",
				UpdateSalesInvoice{
					Status: PaidSalesInvoiceStatus,
				},
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_token_test")
				testMethod(t, r, "PATCH")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
				_, _ = w.Write([]byte(testdata.GetSalesInvoicesResponse))
			},
		},
		{
			"update sales invoice error handler",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
				req: UpdateSalesInvoice{
					Status: PaidSalesInvoiceStatus,
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"update sales invoice, an error occurs when parsing json",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
				req: UpdateSalesInvoice{
					Status: PaidSalesInvoiceStatus,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update sales invoice, invalid url when building request",
			args{
				ctx: context.Background(),
				id:  "invoice_4Y0eZitmBnQ6IDoMqZQKh",
				req: UpdateSalesInvoice{
					Status: PaidSalesInvoiceStatus,
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
			tMux.HandleFunc("/v2/sales-invoices/"+c.args.id, c.handler)

			res, m, err := tClient.SalesInvoices.Update(c.args.ctx, c.args.id, c.args.req)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &SalesInvoice{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
