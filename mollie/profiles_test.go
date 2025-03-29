package mollie

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v4/testdata"
	"github.com/stretchr/testify/assert"
)

func TestProfilesService_Get(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
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
			"get profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"get profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", c.args.profile), c.handler)

			res, m, err := tClient.Profiles.Get(c.args.ctx, c.args.profile)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Profile{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_GetCurrent(t *testing.T) {
	setEnv()
	defer unsetEnv()

	cases := []struct {
		name    string
		wantErr bool
		err     error
		pre     func()
		handler http.HandlerFunc
	}{
		{
			"get current profile works as expected.",
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"get current profile, an error is returned from the server",
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"get current profile, an error occurs when parsing json",
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"get current profile, invalid url when building request",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", "me"), c.handler)

			res, m, err := tClient.Profiles.Current(context.Background())
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Profile{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_List(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		options *ListProfilesOptions
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
			"list profiles works as expected.",
			args{
				context.Background(),
				&ListProfilesOptions{},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfilesListResponse))
			},
		},
		{
			"list profiles with options, works as expected.",
			args{
				context.Background(),
				&ListProfilesOptions{
					Limit: 100,
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "GET")
				testQuery(t, r, "limit=100&testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfilesListResponse))
			},
		},
		{
			"list profiles, an error is returned from the server",
			args{
				context.Background(),
				&ListProfilesOptions{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"list profiles, an error occurs when parsing json",
			args{
				context.Background(),
				&ListProfilesOptions{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"list profiles, invalid url when building request",
			args{
				context.Background(),
				&ListProfilesOptions{},
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
			tMux.HandleFunc("/v2/profiles", c.handler)

			res, m, err := tClient.Profiles.List(c.args.ctx, c.args.options)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &ProfilesList{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_Create(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile CreateOrUpdateProfile
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
			"create profile works as expected.",
			args{
				context.Background(),
				CreateOrUpdateProfile{
					Name: "testing name",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "POST")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"create profile, an error is returned from the server",
			args{
				context.Background(),
				CreateOrUpdateProfile{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"create profile, an error occurs when parsing json",
			args{
				context.Background(),
				CreateOrUpdateProfile{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"create profile, invalid url when building request",
			args{
				context.Background(),
				CreateOrUpdateProfile{},
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
			tMux.HandleFunc("/v2/profiles", c.handler)

			res, m, err := tClient.Profiles.Create(c.args.ctx, c.args.profile)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Profile{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_Update(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx       context.Context
		profileID string
		profile   CreateOrUpdateProfile
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
			"update profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				CreateOrUpdateProfile{
					Name: "testing name",
				},
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "PATCH")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"update profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				CreateOrUpdateProfile{},
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"update profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				CreateOrUpdateProfile{},
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"update profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				CreateOrUpdateProfile{},
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", c.args.profileID), c.handler)

			res, m, err := tClient.Profiles.Update(c.args.ctx, c.args.profileID, c.args.profile)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &Profile{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_Delete(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
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
			"delete profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "DELETE")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.GetProfileResponse))
			},
		},
		{
			"delete profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"delete profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s", c.args.profile), c.handler)

			res, err := tClient.Profiles.Delete(c.args.ctx, c.args.profile)
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

func TestProfilesService_EnablePaymentMethod(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
		method  PaymentMethod
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
			"enable payment method for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "POST")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.EnablePaymentMethodResponse))
			},
		},
		{
			"enable payment method for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"enable payment method for profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable payment method for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/%s", c.args.profile, c.args.method), c.handler)

			res, m, err := tClient.Profiles.EnablePaymentMethod(c.args.ctx, c.args.profile, c.args.method)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &PaymentMethodDetails{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_DisablePaymentMethod(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
		method  PaymentMethod
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
			"disable payment method for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "DELETE")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"disable payment method for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"disable payment method for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PayPal,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/%s", c.args.profile, c.args.method), c.handler)

			res, err := tClient.Profiles.DisablePaymentMethod(c.args.ctx, c.args.profile, c.args.method)
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

func TestProfilesService_EnableGiftCardIssuer(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
		issuer  GiftCardIssuer
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
			"enable gifcard issuer for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "POST")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.EnablePaymentMethodResponse))
			},
		},
		{
			"enable gifcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"enable gifcard issuer for profile, an error occurs when parsing json",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable gifcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", c.args.profile, c.args.issuer), c.handler)

			res, m, err := tClient.Profiles.EnableGiftCardIssuer(c.args.ctx, c.args.profile, c.args.issuer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &GiftCardEnabled{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_DisableGiftCardIssuer(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
		issuer  GiftCardIssuer
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
			"disable giftcard issuer for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "DELETE")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"disable giftcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"disable giftcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", c.args.profile, c.args.issuer), c.handler)

			res, err := tClient.Profiles.DisableGiftCardIssuer(c.args.ctx, c.args.profile, c.args.issuer)
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

func TestProfilesService_EnableGiftCardIssuerForCurrent(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		issuer GiftCardIssuer
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
			"enable gifcard issuer for profile works as expected.",
			args{
				context.Background(),
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "POST")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}
				_, _ = w.Write([]byte(testdata.EnableGiftCardIssuerResponse))
			},
		},
		{
			"enable gifcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"enable gifcard issuer for profile, an error occurs when parsing json",
			args{
				context.Background(),
				Good4fun,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable gifcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/me/methods/giftcard/issuers/%s", c.args.issuer), c.handler)

			res, m, err := tClient.Profiles.EnableGiftCardIssuerForCurrent(c.args.ctx, c.args.issuer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &GiftCardEnabled{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfilesService_DisableGiftCardIssuerForCurrent(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		issuer GiftCardIssuer
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
			"disable giftcard issuer for profile works as expected.",
			args{
				context.Background(),
				Good4fun,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testHeader(t, r, AuthHeader, "Bearer access_X12b31ggg23")
				testMethod(t, r, "DELETE")
				testQuery(t, r, "testmode=true")

				if _, ok := r.Header[AuthHeader]; !ok {
					w.WriteHeader(http.StatusUnauthorized)
				}

				w.WriteHeader(http.StatusOK)
			},
		},
		{
			"disable giftcard issuer for profile, an error is returned from the server",
			args{
				context.Background(),
				Good4fun,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"disable giftcard issuer for profile, invalid url when building request",
			args{
				context.Background(),
				Good4fun,
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
			tMux.HandleFunc(fmt.Sprintf("/v2/profiles/me/methods/giftcard/issuers/%s", c.args.issuer), c.handler)

			res, err := tClient.Profiles.DisableGiftCardIssuerForCurrent(c.args.ctx, c.args.issuer)
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

func TestProfilesService_EnableVoucherIssuer(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
		issuer  VoucherIssuer
		vi      *EnableVoucherIssuer
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
			"enable voucher issuer for profile works as expected using an access key.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
				nil,
			},
			false,
			nil,
			setAccessToken,
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				_, _ = w.Write([]byte(testdata.EnableVoucherIssuerResponse))
			},
		},
		{
			"enable voucher issuer for profile an error is returned from the server.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
				nil,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			setAccessToken,
			errorHandler,
		},
		{
			"enable voucher issuer for profile an error occurs when parsing json.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
				nil,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			setAccessToken,
			encodingHandler,
		},
		{
			"enable voucher issuer for profile invalid url when building request.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
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

			if tClient.HasAccessToken() {
				tMux.HandleFunc(
					fmt.Sprintf("/v2/profiles/%s/methods/voucher/issuers/%s",
						c.args.profile,
						c.args.issuer,
					),
					c.handler,
				)
			} else {
				tMux.HandleFunc(
					fmt.Sprintf("/v2/profiles/me/methods/voucher/issuers/%s",
						c.args.issuer,
					),
					c.handler,
				)
			}

			res, m, err := tClient.Profiles.EnableVoucherIssuer(c.args.ctx, c.args.profile, c.args.issuer, c.args.vi)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &VoucherIssuerEnabled{}, m)
				assert.IsType(t, &http.Response{}, res.Response)
			}
		})
	}
}

func TestProfileService_DisableVoucherIssuer(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx     context.Context
		profile string
		issuer  VoucherIssuer
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
			"disable voucher issuer for profile works as expected.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
			},
			false,
			nil,
			func() {
				tClient.WithAuthenticationValue("access_X12b31ggg23")
			},
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"disable voucher issuer for profile an error is returned from the server.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"disable voucher issuer for profile invalid url when building request.",
			args{
				context.Background(),
				"pfl_v9hTwCvYqw",
				PluxeeEcoVoucher,
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
				fmt.Sprintf("/v2/profiles/%s/methods/voucher/issuers/%s",
					c.args.profile,
					c.args.issuer,
				),
				c.handler,
			)

			res, err := tClient.Profiles.DisableVoucherIssuer(c.args.ctx, c.args.profile, c.args.issuer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &http.Response{}, res.Response)
				assert.Equal(t, http.StatusNoContent, res.StatusCode)
			}
		})
	}
}

func TestProfileService_EnableVoucherIssuerForCurrent(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		issuer VoucherIssuer
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
			"enable voucher issuer for current profile works as expected.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "POST")
				w.WriteHeader(http.StatusCreated)
				_, _ = w.Write([]byte(testdata.EnableVoucherIssuerResponse))
			},
		},
		{
			"enable voucher issuer for current profile an error is returned from the server.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"enable voucher issuer for current profile an error occurs when parsing json.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
			},
			true,
			fmt.Errorf("invalid character 'h' looking for beginning of object key string"),
			noPre,
			encodingHandler,
		},
		{
			"enable voucher issuer for current profile invalid url when building request.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
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
				fmt.Sprintf("/v2/profiles/me/methods/voucher/issuers/%s",
					c.args.issuer,
				),
				c.handler,
			)

			res, vi, err := tClient.Profiles.EnableVoucherIssuerForCurrent(c.args.ctx, c.args.issuer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &http.Response{}, res.Response)
				assert.IsType(t, &VoucherIssuerEnabled{}, vi)
				assert.Equal(t, http.StatusCreated, res.StatusCode)
			}
		})
	}
}

func TestProfileService_DisableVoucherIssuerForCurrent(t *testing.T) {
	setEnv()
	defer unsetEnv()

	type args struct {
		ctx    context.Context
		issuer VoucherIssuer
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
			"disable voucher issuer for current profile works as expected.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
			},
			false,
			nil,
			noPre,
			func(w http.ResponseWriter, r *http.Request) {
				testMethod(t, r, "DELETE")
				w.WriteHeader(http.StatusNoContent)
			},
		},
		{
			"disable voucher issuer for current profile an error is returned from the server.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
			},
			true,
			fmt.Errorf("500 Internal Server Error: An internal server error occurred while processing your request"),
			noPre,
			errorHandler,
		},
		{
			"disable voucher issuer for current profile invalid url when building request.",
			args{
				context.Background(),
				PluxeeEcoVoucher,
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
				fmt.Sprintf("/v2/profiles/me/methods/voucher/issuers/%s",
					c.args.issuer,
				),
				c.handler,
			)

			res, err := tClient.Profiles.DisableVoucherIssuerForCurrent(c.args.ctx, c.args.issuer)
			if c.wantErr {
				assert.NotNil(t, err)
				assert.EqualError(t, err, c.err.Error())
			} else {
				assert.Nil(t, err)
				assert.IsType(t, &http.Response{}, res.Response)
				assert.Equal(t, http.StatusNoContent, res.StatusCode)
			}
		})
	}
}
