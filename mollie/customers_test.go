package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestCustomerService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx      context.Context
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
			"get mollie customers works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
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

				_, _ = w.Write([]byte(testdata.GetCustomerResponse))
			},
		},
		{
			"get mollie customers, an error is returned from the server",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get mollie customers, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get mollie customers, invalid url when building request",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s", c.args.customer), c.handler)
			c.pre()
			res, cc, err := tClient.Customers.Get(c.args.ctx, c.args.customer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Customer{}, cc)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCustomersService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()
	type args struct {
		ctx      context.Context
		customer CreateCustomer
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"create mollie customers works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				CreateCustomer{Locale: German},
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

				_, _ = w.Write([]byte(testdata.CreateCustomerResponse))
			},
		},
		{
			"create mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				CreateCustomer{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create mollie customers, an error occurs when parsing json",
			http.StatusInternalServerError,
			args{
				context.Background(),
				CreateCustomer{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				CreateCustomer{},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc("/v2/customers", c.handler)
			c.pre()
			res, cc, err := tClient.Customers.Create(c.args.ctx, c.args.customer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Customer{}, cc)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCustomersService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx        context.Context
		customerID string
		customer   UpdateCustomer
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"update mollie customers works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				UpdateCustomer{Locale: French},
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

				_, _ = w.Write([]byte(testdata.UpdateCustomerResponse))
			},
		},
		{
			"update mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				UpdateCustomer{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"update mollie customers, an error occurs when parsing json",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				UpdateCustomer{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				UpdateCustomer{},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s", c.args.customerID), c.handler)
			c.pre()
			res, cc, err := tClient.Customers.Update(c.args.ctx, c.args.customerID, c.args.customer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Customer{}, cc)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCustomersService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()
	type args struct {
		ctx     context.Context
		options *ListCustomersOptions
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"list mollie customers with options works as expected",
			http.StatusAccepted,
			args{
				context.Background(),
				&ListCustomersOptions{
					SequenceType: OneOffSequence,
				},
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

				_, _ = w.Write([]byte(testdata.ListCustomersResponse))
			},
		},
		{
			"list mollie customers works as expected",
			http.StatusAccepted,
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

				_, _ = w.Write([]byte(testdata.ListCustomersResponse))
			},
		},
		{
			"list mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
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
			"list mollie customers, an error occurs when parsing json",
			http.StatusInternalServerError,
			args{
				context.Background(),
				&ListCustomersOptions{
					SequenceType: OneOffSequence,
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				&ListCustomersOptions{
					SequenceType: OneOffSequence,
				},
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc("/v2/customers", c.handler)
			c.pre()
			res, cc, err := tClient.Customers.List(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &CustomersList{}, cc)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCustomersService_Delete(t *testing.T) {
	setEnv()
	defer unsetEnv()
	type args struct {
		ctx      context.Context
		customer string
	}

	cases := []struct {
		name    string
		status  int
		args    args
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"delete mollie customers with options works as expected",
			http.StatusNoContent,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
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
			"delete mollie customers, an error is returned from the server",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"delete mollie customers, invalid url when building request",
			http.StatusInternalServerError,
			args{
				context.Background(),
				"cst_kEn1PlbGa",
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s", c.args.customer), c.handler)
			c.pre()
			res, err := tClient.Customers.Delete(c.args.ctx, c.args.customer)
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

func TestCustomerService_GetPayments(t *testing.T) {
	setEnv()
	defer unsetEnv()
	type args struct {
		ctx      context.Context
		customer string
		options  *ListCustomersOptions
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
			"get mollie customers payments works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
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
			"get mollie customers payments with options works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				&ListCustomersOptions{Limit: 100},
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
			"get mollie customers payments, an error is returned from the server",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				&ListCustomersOptions{SequenceType: RecurringSequence},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get mollie customers payments, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get mollie customers payments, invalid url when building request",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				nil,
			},
			true,
			errBadBaseURL,
			crashSrv,
			errorHandler,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/payments", c.args.customer), c.handler)
			c.pre()
			res, cc, err := tClient.Customers.GetPayments(c.args.ctx, c.args.customer, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentList{}, cc)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestCustomerService_CreatePayment(t *testing.T) {
	setEnv()
	defer unsetEnv()
	type args struct {
		ctx      context.Context
		customer string
		payment  CreatePayment
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
			"create mollie customers payments works as expected",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				CreatePayment{
					CreatePaymentAccessTokenFields: CreatePaymentAccessTokenFields{
						Testmode: true,
					},
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

				_, _ = w.Write([]byte(testdata.ListPaymentsResponse))
			},
		},
		{
			"create mollie customers payments, an error is returned from the server",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				CreatePayment{
					CreatePaymentAccessTokenFields: CreatePaymentAccessTokenFields{
						Testmode: true,
					},
				},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create mollie customers payments, an error occurs when parsing json",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				CreatePayment{
					CreatePaymentAccessTokenFields: CreatePaymentAccessTokenFields{
						Testmode: true,
					},
				},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create mollie customers payments, invalid url when building request",
			args{
				context.Background(),
				"cst_kEn1PlbGa",
				CreatePayment{
					CreatePaymentAccessTokenFields: CreatePaymentAccessTokenFields{
						Testmode: true,
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
		t.Run(c.name, func(t *testing.T) {
			setup()
			defer teardown()
			tMux.HandleFunc(fmt.Sprintf("/v2/customers/%s/payments", c.args.customer), c.handler)
			c.pre()
			res, cc, err := tClient.Customers.CreatePayment(c.args.ctx, c.args.customer, c.args.payment)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.Nil(t, err)
				assert.IsType(t, &Payment{}, cc)
				assert.EqualValues(t, c.args.ctx, res.Request.Context())
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}
