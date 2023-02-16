package mollie

// Config contains information that helps during the setup of a new Mollie client.
type Config struct {
	testing        bool
	auth           string
	reqIdempotency bool
}

// ToggleTesting enables/disables the test-mode in the current Config.
func (c *Config) ToggleTesting() bool {
	c.testing = !c.testing

	return c.testing
}

// ToggleIdempotency enables/disables the request idempotency feature
// in the current Config.
func (c *Config) ToggleIdempotency() bool {
	c.reqIdempotency = !c.reqIdempotency

	return c.reqIdempotency
}

// SwitchAuthStrategy changes the environment variable used to fetch the
// auth tokens.
//
// Known values are: [MOLLIE_API_TOKEN,MOLLIE_ORG_TOKEN], if you use a custom
// environment variable pass it as argument.
func (c *Config) SwitchAuthStrategy(auth string) string {
	c.auth = auth

	return c.auth
}

/* Configuration init helpers.  */

// NewConfig builds a Mollie configuration object,
// it takes t to indicate if our client is meant to create requests for testing,
// and auth to indicate the authentication method we want to use.
func NewConfig(t bool, auth string) *Config {
	return createConfig(t, false, auth)
}

// NewAPITestingConfig builds a configuration object with the following settings:
// tests mode: enabled
// api token source: MOLLIE_API_TOKEN
//
// it receives `reqIdem (boolean)` to enable the request idempotency feature.
func NewAPITestingConfig(reqIdem bool) *Config {
	return createConfig(true, reqIdem, APITokenEnv)
}

// NewAPIConfig builds a configuration object with the following settings:
// tests mode: disabled
// api token source: MOLLIE_API_TOKEN
//
// it receives `reqIdem (boolean)` to enable the request idempotency feature.
func NewAPIConfig(reqIdem bool) *Config {
	return createConfig(false, reqIdem, APITokenEnv)
}

// NewOrgTestingConfig builds a configuration object with the following settings:
// tests mode: enabled
// api token source: MOLLIE_ORG_TOKEN
//
// it receives `reqIdem (boolean)` to enable the request idempotency feature.
func NewOrgTestingConfig(reqIdem bool) *Config {
	return createConfig(true, reqIdem, OrgTokenEnv)
}

// NewOrgConfig builds a configuration object with the following settings:
// tests mode: disabled
// Org token source: MOLLIE_ORG_TOKEN
//
// it receives `reqIdem (boolean)` to enable the request idempotency feature.
func NewOrgConfig(reqIdem bool) *Config {
	return createConfig(false, reqIdem, OrgTokenEnv)
}

func createConfig(test, reqIdem bool, auth string) *Config {
	return &Config{
		testing:        test,
		auth:           auth,
		reqIdempotency: reqIdem,
	}
}
