package mollie

import (
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

	p, err := tClient.Profiles.Get(nil, id)
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

	p, err := tClient.Profiles.Get(nil, id)
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

	p, err := tClient.Profiles.Current(nil)
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

	pl, err := tClient.Profiles.List(nil, opt)
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

	p, err := tClient.Profiles.Create(nil, np)
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

	p, err := tClient.Profiles.Update(nil, id, np)
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

	err := tClient.Profiles.Delete(nil, id)
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

	res, err := tClient.Profiles.EnablePaymentMethod(nil, id, Bancontact)
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

	err := tClient.Profiles.DisablePaymentMethod(nil, id, Bancontact)
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

	res, err := tClient.Profiles.EnableGiftCardIssuer(nil, id, Festivalcadeau)

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

	err := tClient.Profiles.DisableGiftCardIssuer(nil, id, Festivalcadeau)

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

	res, err := tClient.Profiles.EnableGiftCardIssuerForCurrent(nil, Festivalcadeau)

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

	err := tClient.Profiles.DisableGiftCardIssuerForCurrent(nil, Festivalcadeau)

	if err != nil {
		t.Error(err)
	}
}

func TestProfilesService_HttpRequestErrors(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/v2/profiles/", errorHandler)

	p := Profile{}

	_, cerr := tClient.Profiles.Create(nil, &p)
	_, rerr := tClient.Profiles.List(nil, nil)
	_, uerr := tClient.Profiles.Update(nil, "1212", &p)
	derr := tClient.Profiles.Delete(nil, "1212")
	_, gerr := tClient.Profiles.Get(nil, "1212")
	_, ccerr := tClient.Profiles.Current(nil)
	dderr := tClient.Profiles.DisablePaymentMethod(nil, "1212", PayPal)
	_, eperr := tClient.Profiles.EnablePaymentMethod(nil, "1212", PayPal)
	dgcerr := tClient.Profiles.DisableGiftCardIssuer(nil, "1212", Festivalcadeau)
	_, egcerr := tClient.Profiles.EnableGiftCardIssuer(nil, "1212", Festivalcadeau)
	dgccerr := tClient.Profiles.DisableGiftCardIssuerForCurrent(nil, Festivalcadeau)
	_, egccerr := tClient.Profiles.EnableGiftCardIssuerForCurrent(nil, Festivalcadeau)

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

	_, cerr := tClient.Profiles.Create(nil, &p)
	_, rerr := tClient.Profiles.List(nil, nil)
	_, uerr := tClient.Profiles.Update(nil, "1212", &p)
	_, gerr := tClient.Profiles.Get(nil, "1212")
	_, ccerr := tClient.Profiles.Current(nil)
	_, eperr := tClient.Profiles.EnablePaymentMethod(nil, "1212", PayPal)
	_, egcerr := tClient.Profiles.EnableGiftCardIssuer(nil, "1212", Festivalcadeau)
	_, egccerr := tClient.Profiles.EnableGiftCardIssuerForCurrent(nil, Festivalcadeau)

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

	_, cerr := tClient.Profiles.Create(nil, &p)
	_, rerr := tClient.Profiles.List(nil, nil)
	_, uerr := tClient.Profiles.Update(nil, "1212", &p)
	_, gerr := tClient.Profiles.Get(nil, "1212")
	_, ccerr := tClient.Profiles.Current(nil)
	ddcerr := tClient.Profiles.Delete(nil, "1212")
	_, eperr := tClient.Profiles.EnablePaymentMethod(nil, "1212", PayPal)
	deperr := tClient.Profiles.DisablePaymentMethod(nil, "1212", PayPal)
	dgcerr := tClient.Profiles.DisableGiftCardIssuer(nil, "1212", Festivalcadeau)
	_, egcerr := tClient.Profiles.EnableGiftCardIssuer(nil, "1212", Festivalcadeau)

	tests := []error{cerr, rerr, uerr, gerr, ccerr, eperr, deperr, ddcerr, dgcerr, egcerr}

	for _, tt := range tests {
		if tt != errBadBaseURL {
			t.Error(tt)
		}
	}
}
