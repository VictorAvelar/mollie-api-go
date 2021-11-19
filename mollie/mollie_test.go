package mollie

import (
	"context"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"strings"
	"testing"
	"time"
)

func TestNewClient(t *testing.T) {
	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	conf := NewConfig(true, APITokenEnv)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewClient(tt.client, conf)
			if err != nil {
				t.Errorf("not nil error received: %v", err)
			}
		})
	}
}

func TestNewClientWithEnvVars(t *testing.T) {
	setEnv()
	defer unsetEnv()

	var c = http.DefaultClient
	{
		c.Timeout = 25 * time.Second
	}

	tests := []struct {
		name   string
		client *http.Client
	}{
		{
			"nil returns a valid client",
			nil,
		},
		{
			"a passed client is decorated",
			c,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewClient(tt.client, tConf)
			if err != nil {
				t.Errorf("not nil error received: %v", err)
			}

			if got.authentication == "" {
				t.Errorf("got empty api key %v, value %s expected", got.authentication, "token_X12b31ggg23")
			}
		})
	}
}

func TestClient_NewAPIRequest(t *testing.T) {
	setup()
	defer teardown()
	b := []string{"hello", "bye"}
	inURL, outURL := "test", tServer.URL+"/test"
	inBody, outBody := b, `["hello","bye"]`+"\n"
	_ = tClient.WithAuthenticationValue("test_token")
	req, _ := tClient.NewAPIRequest(context.TODO(), "GET", inURL, inBody)

	testHeader(t, req, "Accept", RequestContentType)
	testHeader(t, req, AuthHeader, "Bearer test_token")
	// test that relative URL was expanded
	if got, want := req.URL.String(), outURL; got != want {
		t.Errorf("NewRequest(%q) URL is %v, want %v", inURL, got, want)
	}

	// test that body was JSON encoded
	body, _ := ioutil.ReadAll(req.Body)
	if got, want := string(body), outBody; got != want {
		t.Errorf("NewRequest(%q) Body is %v, want %v", inBody, got, want)
	}
}

func TestClient_IsAccessToken(t *testing.T) {
	cases := []struct {
		name  string
		value string
		want  bool
	}{
		{
			"Test dummytext",
			"dummytext",
			false,
		},
		{
			"Test pattern is matched",
			"access_testingtokenhere",
			true,
		},
		{
			"Test patter is not matched if not in the right position",
			"testing_tokent_access_here",
			false,
		},
		{
			"Test other tokens are not matched",
			"test_yourtokenhere",
			false,
		},
		{
			"Test empty is not a matched",
			"",
			false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			client := &Client{
				authentication: c.value,
			}

			got := client.HasAccessToken()

			if got != c.want {
				t.Errorf("Mismatching token check, want %v, got %v", c.want, got)
			}
		})
	}
}

func TestClient_NewAPIRequest_ErrTrailingSlash(t *testing.T) {
	uri, _ := url.Parse("http://localhost")
	tClient = &Client{
		BaseURL: uri,
	}
	_, err := tClient.NewAPIRequest(context.TODO(), "GET", "test", nil)

	if err == nil {
		t.Errorf("expected error %v not occurred, got %v", errBadBaseURL, err)
	}
}

func TestClient_NewAPIRequest_HTTPReqNativeError(t *testing.T) {
	setup()
	defer teardown()
	_, err := tClient.NewAPIRequest(context.TODO(), "\\\\\\", "test", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "invalid method") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_OrgTokenOverApiKey(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("org_token")
	req, _ := tClient.NewAPIRequest(context.TODO(), "GET", "test", nil)

	testHeader(t, req, AuthHeader, "Bearer org_token")
}

func TestClient_WithAuthenticationValue_Error(t *testing.T) {
	setup()
	defer teardown()
	err := tClient.WithAuthenticationValue("")

	if err == nil {
		t.Errorf("unexpected error, want %v and got %v", errEmptyAuthKey, err)
	}
}

func TestClient_NewAPIRequest_ErrorBodySerialization(t *testing.T) {
	setup()
	defer teardown()
	b := make(chan int)
	_, err := tClient.NewAPIRequest(context.TODO(), "GET", "test", b)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "unsupported type") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_NativeURLParseError(t *testing.T) {
	setup()
	defer teardown()
	_, err := tClient.NewAPIRequest(context.TODO(), "GET", ":", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "parse") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_Do(t *testing.T) {
	setup()
	defer teardown()

	tMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, AuthHeader, "Bearer test_token")
		w.WriteHeader(http.StatusOK)
	})
	_ = tClient.WithAuthenticationValue("test_token")
	req, _ := tClient.NewAPIRequest(context.TODO(), "GET", "test", nil)
	res, err := tClient.Do(req)

	if err != nil {
		t.Errorf("unexpected error received: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("request failed: %+v", res)
	}
}

func TestClient_DoErrInvalidJSON(t *testing.T) {
	setup()
	defer teardown()
	tMux.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		testMethod(t, r, "GET")
		testHeader(t, r, AuthHeader, "Bearer test_token")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("{"))
	})
	_ = tClient.WithAuthenticationValue("test_token")
	req, _ := tClient.NewAPIRequest(context.TODO(), "GET", "test", nil)
	req.URL = nil
	_, err := tClient.Do(req)

	if err == nil {
		t.Error(err)
	}
	if !strings.Contains(err.Error(), "nil Request.URL") {
		t.Errorf("unexpected response, got %v", err)
	}
}

func TestClient_DoErr(t *testing.T) {
	setup()
	defer teardown()
	req, _ := tClient.NewAPIRequest(context.TODO(), "GET", "test", nil)
	req.URL = nil
	_, err := tClient.Do(req)

	if err == nil {
		t.Error(err)
	}

	if !strings.Contains(err.Error(), "nil Request.URL") {
		t.Errorf("unexpected response, got %v", err)
	}
}

func TestCheckResponse(t *testing.T) {
	res1 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(strings.NewReader("not found ok")),
	}

	res3 := &http.Response{
		StatusCode: http.StatusNotFound,
		Status:     http.StatusText(http.StatusNotFound),
		Body:       ioutil.NopCloser(strings.NewReader("")),
	}

	res2 := &http.Response{
		StatusCode: http.StatusOK,
		Status:     http.StatusText(http.StatusOK),
		Body:       ioutil.NopCloser(strings.NewReader("success ok")),
	}

	tests := []struct {
		name string
		code string
		arg  *http.Response
	}{
		{
			"successful response",
			"",
			res2,
		},
		{
			"not found response",
			"Not Found",
			res1,
		},
		{
			"success with empty body",
			"",
			res3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := CheckResponse(tt.arg); err != nil {
				if !strings.Contains(err.Error(), tt.code) {
					t.Error(err)
				}
			}
		})
	}
}

// <----- Testing helpers ----->

var (
	tMux    *http.ServeMux
	tServer *httptest.Server
	tClient *Client
	tConf   *Config
)

// the parameter indicates if you want to prepare your tests against the US sandbox
// just to be used when doing integration testing.
func setup() {
	tMux = http.NewServeMux()
	tServer = httptest.NewServer(tMux)
	tConf = NewConfig(true, APITokenEnv)
	tClient, _ = NewClient(nil, tConf)
	u, _ := url.Parse(tServer.URL + "/")
	tClient.BaseURL = u
}

func teardown() {
	tServer.Close()
}

func setEnv() {
	_ = os.Setenv(APITokenEnv, "token_X12b31ggg23")
	_ = os.Setenv(OrgTokenEnv, "ey1923n23123n1k3b123jv12g312h31v32g13")
}

func unsetEnv() {
	_ = os.Unsetenv(APITokenEnv)
	_ = os.Unsetenv(OrgTokenEnv)
}

func testMethod(t *testing.T, r *http.Request, want string) {
	if got := r.Method; got != want {
		t.Errorf("Request method: %v, want %v", got, want)
	}
}

func testHeader(t *testing.T, r *http.Request, header string, want string) {
	if got := r.Header.Get(header); got != want {
		t.Errorf("Header.Get(%q) returned %q, want %q", header, got, want)
	}
}

// <----- .Testing helpers ----->
