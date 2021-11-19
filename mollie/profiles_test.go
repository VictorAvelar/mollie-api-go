package mollie

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
)

func TestProfilesService_Get(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetProfileResponse))
	})

	p, err := tClient.Profiles.Get(context.TODO(), id)
	if err != nil {
		t.Error(err)
	}

	if p.ID != id {
		t.Errorf("unexpected response, want: %v got: %v", p.ID, id)
	}
}

func ExampleProfilesService_Get() {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/"+id, func(w http.ResponseWriter, r *http.Request) {
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetProfileResponse))
	})

	p, err := tClient.Profiles.Get(context.TODO(), id)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(p.ID)
	//Output: pfl_v9hTwCvYqw
}

func TestProfilesService_Current(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/me", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetProfileResponse))
	})

	p, err := tClient.Profiles.Current(context.TODO())
	if err != nil {
		t.Error(err)
	}

	if p.ID != id {
		t.Errorf("unexpected response, want: %v got: %v", p.ID, id)
	}
}

func TestProfilesService_List(t *testing.T) {
	setup()
	defer teardown()

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "GET")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetProfilesListResponse))
	})

	opt := &ProfileListOptions{
		Limit: 10,
	}

	pl, err := tClient.Profiles.List(context.TODO(), opt)
	if err != nil {
		t.Error(err)
	}

	if len(pl.Embedded.Profiles) < 1 {
		t.Errorf("returned empty list")
	}
}

func TestProfilesService_Create(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusCreated)
		_, _ = w.Write([]byte(testdata.CreateProfileResponse))
	})
	np := &Profile{
		ID:           "pfl_v9hTwCvYqw",
		CategoryCode: GeneralMerchandise,
		Email:        "info@mywebsite.com",
		Mode:         LiveMode,
		Name:         "My website name",
	}

	p, err := tClient.Profiles.Create(context.TODO(), np)
	if err != nil {
		t.Error(err)
	}

	if p.ID != np.ID {
		t.Errorf("unexpected response, want: %v got: %v", p.ID, id)
	}
}

func TestProfilesService_Update(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "PATCH")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.GetProfileResponse))
	})

	np := &Profile{
		ID:           "pfl_v9hTwCvYqw",
		CategoryCode: GeneralMerchandise,
		Email:        "info@mywebsite.com",
		Mode:         LiveMode,
		Name:         "My website name",
		Phone:        "",
	}

	p, err := tClient.Profiles.Update(context.TODO(), id, np)
	if err != nil {
		t.Error(err)
	}

	if p.ID != id {
		t.Errorf("unexpected response, want: %v got: %v", p.ID, id)
	}
}

func TestProfilesService_Delete(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/"+id, func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "DELETE")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	err := tClient.Profiles.Delete(context.TODO(), id)
	if err != nil {
		t.Error(err)
	}
}

func TestProfilesService_EnablePaymentMethod(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/"+id+"/methods/bancontact", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "POST")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(testdata.EnablePaymentMethodResponse))
	})

	res, err := tClient.Profiles.EnablePaymentMethod(context.TODO(), id, Bancontact)
	if err != nil {
		t.Error(err)
	}

	if res.ID != string(Bancontact) {
		t.Errorf("unexpected response, want: %v got: %v", Bancontact, res.ID)
	}
}

func TestProfilesService_DisablePaymentMethod(t *testing.T) {
	setup()
	defer teardown()
	id := "pfl_v9hTwCvYqw"

	_ = tClient.WithAuthenticationValue("test_token")

	tMux.HandleFunc("/v2/profiles/"+id+"/methods/bancontact", func(w http.ResponseWriter, r *http.Request) {
		testHeader(t, r, AuthHeader, "Bearer test_token")
		testMethod(t, r, "DELETE")
		if _, ok := r.Header[AuthHeader]; !ok {
			w.WriteHeader(http.StatusUnauthorized)
		}

		w.WriteHeader(http.StatusNoContent)
	})

	err := tClient.Profiles.DisablePaymentMethod(context.TODO(), id, Bancontact)
	if err != nil {
		t.Error(err)
	}
}

func TestProfilesService_EnableGiftCardIssuer(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		teardown()
		unsetEnv()
	}()

	id := "pfl_v9hTwCvYqw"

	tMux.HandleFunc(
		fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", id, Festivalcadeau),
		func(rw http.ResponseWriter, r *http.Request) {
			testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
			testMethod(t, r, "POST")
			if _, ok := r.Header[AuthHeader]; !ok {
				rw.WriteHeader(http.StatusUnauthorized)
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(testdata.EnableGiftCardIssuerResponse))
		},
	)

	res, err := tClient.Profiles.EnableGiftCardIssuer(context.TODO(), id, Festivalcadeau)

	if err != nil {
		t.Error(err)
	}

	if res.ID != Festivalcadeau {
		t.Errorf("unexpected id received: want %v, got %v", Festivalcadeau, res.ID)
	}
}

func TestProfilesService_DisableGiftCardIssuer(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		teardown()
		unsetEnv()
	}()

	id := "pfl_v9hTwCvYqw"

	tMux.HandleFunc(
		fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", id, Festivalcadeau),
		func(rw http.ResponseWriter, r *http.Request) {
			testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
			testMethod(t, r, "DELETE")
			if _, ok := r.Header[AuthHeader]; !ok {
				rw.WriteHeader(http.StatusUnauthorized)
			}

			rw.WriteHeader(http.StatusNoContent)
		},
	)

	err := tClient.Profiles.DisableGiftCardIssuer(context.TODO(), id, Festivalcadeau)

	if err != nil {
		t.Error(err)
	}
}

func TestProfilesService_EnableGiftCardIssuerForCurrent(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		teardown()
		unsetEnv()
	}()

	id := "me"

	tMux.HandleFunc(
		fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", id, Festivalcadeau),
		func(rw http.ResponseWriter, r *http.Request) {
			testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
			testMethod(t, r, "POST")
			if _, ok := r.Header[AuthHeader]; !ok {
				rw.WriteHeader(http.StatusUnauthorized)
			}

			rw.WriteHeader(http.StatusOK)
			_, _ = rw.Write([]byte(testdata.EnableGiftCardIssuerResponse))
		},
	)

	res, err := tClient.Profiles.EnableGiftCardIssuerForCurrent(context.TODO(), Festivalcadeau)

	if err != nil {
		t.Error(err)
	}

	if res.ID != Festivalcadeau {
		t.Errorf("unexpected id received: want %v, got %v", Festivalcadeau, res.ID)
	}
}

func TestProfilesService_DisableGiftCardIssuerForCurrent(t *testing.T) {
	setEnv()
	setup()
	defer func() {
		teardown()
		unsetEnv()
	}()

	id := "me"

	tMux.HandleFunc(
		fmt.Sprintf("/v2/profiles/%s/methods/giftcard/issuers/%s", id, Festivalcadeau),
		func(rw http.ResponseWriter, r *http.Request) {
			testHeader(t, r, AuthHeader, "Bearer token_X12b31ggg23")
			testMethod(t, r, "DELETE")
			if _, ok := r.Header[AuthHeader]; !ok {
				rw.WriteHeader(http.StatusUnauthorized)
			}

			rw.WriteHeader(http.StatusNoContent)
		},
	)

	err := tClient.Profiles.DisableGiftCardIssuerForCurrent(context.TODO(), Festivalcadeau)

	if err != nil {
		t.Error(err)
	}
}

func TestProfilesService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/profiles/", errorHandler)

	p := Profile{}

	_, cerr := tClient.Profiles.Create(context.TODO(), &p)
	_, rerr := tClient.Profiles.List(context.TODO(), nil)
	_, uerr := tClient.Profiles.Update(context.TODO(), "1212", &p)
	derr := tClient.Profiles.Delete(context.TODO(), "1212")
	_, gerr := tClient.Profiles.Get(context.TODO(), "1212")
	_, ccerr := tClient.Profiles.Current(context.TODO())
	dderr := tClient.Profiles.DisablePaymentMethod(context.TODO(), "1212", PayPal)
	_, eperr := tClient.Profiles.EnablePaymentMethod(context.TODO(), "1212", PayPal)
	dgcerr := tClient.Profiles.DisableGiftCardIssuer(context.TODO(), "1212", Festivalcadeau)
	_, egcerr := tClient.Profiles.EnableGiftCardIssuer(context.TODO(), "1212", Festivalcadeau)
	dgccerr := tClient.Profiles.DisableGiftCardIssuerForCurrent(context.TODO(), Festivalcadeau)
	_, egccerr := tClient.Profiles.EnableGiftCardIssuerForCurrent(context.TODO(), Festivalcadeau)

	tests := []error{cerr, rerr, uerr, derr, gerr, ccerr, eperr, dderr, dgcerr, egcerr, dgccerr, egccerr}

	for _, tt := range tests {
		if tt == nil {
			t.Fail()
		}
	}
}

func TestProfilesService_EncodingResponseErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/profiles/", encodingHandler)

	p := Profile{}

	_, cerr := tClient.Profiles.Create(context.TODO(), &p)
	_, rerr := tClient.Profiles.List(context.TODO(), nil)
	_, uerr := tClient.Profiles.Update(context.TODO(), "1212", &p)
	_, gerr := tClient.Profiles.Get(context.TODO(), "1212")
	_, ccerr := tClient.Profiles.Current(context.TODO())
	_, eperr := tClient.Profiles.EnablePaymentMethod(context.TODO(), "1212", PayPal)
	_, egcerr := tClient.Profiles.EnableGiftCardIssuer(context.TODO(), "1212", Festivalcadeau)
	_, egccerr := tClient.Profiles.EnableGiftCardIssuerForCurrent(context.TODO(), Festivalcadeau)

	tests := []error{cerr, rerr, uerr, gerr, ccerr, eperr, egcerr, egccerr}

	for _, tt := range tests {
		if tt == nil {
			t.Error(tt)
		} else if !strings.Contains(tt.Error(), "invalid character") {
			t.Errorf("unexpected error %v", tt)
		}
	}
}

func TestProfilesService_NewAPIRequestBaseError(t *testing.T) {
	setup()
	defer teardown()
	u, _ := url.Parse(tServer.URL)
	tClient.BaseURL = u
	tMux.HandleFunc("/v2/profiles/", errorHandler)

	p := Profile{}

	_, cerr := tClient.Profiles.Create(context.TODO(), &p)
	_, rerr := tClient.Profiles.List(context.TODO(), nil)
	_, uerr := tClient.Profiles.Update(context.TODO(), "1212", &p)
	_, gerr := tClient.Profiles.Get(context.TODO(), "1212")
	_, ccerr := tClient.Profiles.Current(context.TODO())
	ddcerr := tClient.Profiles.Delete(context.TODO(), "1212")
	_, eperr := tClient.Profiles.EnablePaymentMethod(context.TODO(), "1212", PayPal)
	deperr := tClient.Profiles.DisablePaymentMethod(context.TODO(), "1212", PayPal)
	dgcerr := tClient.Profiles.DisableGiftCardIssuer(context.TODO(), "1212", Festivalcadeau)
	_, egcerr := tClient.Profiles.EnableGiftCardIssuer(context.TODO(), "1212", Festivalcadeau)

	tests := []error{cerr, rerr, uerr, gerr, ccerr, eperr, deperr, ddcerr, dgcerr, egcerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}
