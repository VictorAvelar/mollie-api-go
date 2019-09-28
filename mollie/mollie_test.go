package mollie

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"reflect"
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
	req, _ := tClient.NewAPIRequest("GET", inURL, inBody)

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

func TestClient_NewAPIRequest_ErrTrailingSlash(t *testing.T) {
	uri, _ := url.Parse("http://localhost")
	tClient = &Client{
		BaseURL: uri,
	}
	_, err := tClient.NewAPIRequest("GET", "test", nil)

	if !reflect.DeepEqual(err, errBadBaseURL) {
		t.Errorf("expected error %v not ocurred, got %v", errBadBaseURL, err)
	}
}

func TestClient_NewAPIRequest_HTTPReqNativeError(t *testing.T) {
	setup()
	defer teardown()
	_, err := tClient.NewAPIRequest("\\\\\\", "test", nil)

	if err == nil {
		t.Fatal("nil error produced")
	}

	if !strings.Contains(err.Error(), "invalid method") {
		t.Errorf("unexpected err received %v", err)
	}
}

func TestClient_NewAPIRequest_NoAuthKey(t *testing.T) {
	setup()
	defer teardown()
	_, err := tClient.NewAPIRequest("GET", "test", nil)

	if err != errEmptyAPIKey {
		t.Errorf("unexpected error %v", err)
	}
}

func TestClient_NewAPIRequest_OrgTokenOverApiKey(t *testing.T) {
	setup()
	defer teardown()
	_ = tClient.WithAuthenticationValue("org_token")
	req, _ := tClient.NewAPIRequest("GET", "test", nil)

	testHeader(t, req, AuthHeader, "Bearer org_token")
}

func TestClient_NewAPIRequest_ErrorBodySerialization(t *testing.T) {
	setup()
	defer teardown()
	b := make(chan int)
	_, err := tClient.NewAPIRequest("GET", "test", b)

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
	_, err := tClient.NewAPIRequest("GET", ":", nil)

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
	req, _ := tClient.NewAPIRequest("GET", "test", nil)
	res, err := tClient.Do(req)

	if err != nil {
		t.Errorf("unexpected error received: %v", err)
	}

	if res.StatusCode != http.StatusOK {
		t.Errorf("request failed: %+v", res)
	}
}

// ----- mollie examples -----
func ExampleNewClient() {
	c, err := NewClient(nil, &Config{
		testing: false,
		auth:    APITokenEnv,
	})
	log.Printf("%+v", c)
	fmt.Println(err == nil)
	// Output: true
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

func testURLParseError(t *testing.T, err error) {
	if err == nil {
		t.Errorf("Expected error to be returned")
	}
	if err, ok := err.(*url.Error); !ok || err.Op != "parse" {
		t.Errorf("Expected URL parse error, got %+v", err)
	}
}

func testBody(t *testing.T, r *http.Request, want string) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		t.Errorf("Error reading request body: %v", err)
	}
	if got := string(b); got != want {
		t.Errorf("request Body is %s, want %s", got, want)
	}
}

func testJSONMarshal(t *testing.T, v interface{}, want string) {
	j, err := json.Marshal(v)
	if err != nil {
		t.Errorf("Unable to marshal JSON for %v", v)
	}

	w := new(bytes.Buffer)
	err = json.Compact(w, []byte(want))
	if err != nil {
		t.Errorf("String is not valid json: %s", want)
	}

	if w.String() != string(j) {
		t.Errorf("json.Marshal(%q) returned %s, want %s", v, j, w)
	}

	// now go the other direction and make sure things unmarshal as expected
	u := reflect.ValueOf(v).Interface()
	if err := json.Unmarshal([]byte(want), u); err != nil {
		t.Errorf("Unable to unmarshal JSON for %v: %v", want, err)
	}

	if !reflect.DeepEqual(v, u) {
		t.Errorf("json.Unmarshal(%q) returned %s, want %s", want, u, v)
	}
}

// <----- .Testing helpers ----->
