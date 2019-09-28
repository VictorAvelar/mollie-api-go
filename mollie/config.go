package mollie

// Config contains information that helps during the setup of a new Mollie client.
type Config struct {
	testing bool
	auth    string
}

//NewConfig build a Mollie configuration object,
// it takes t to indicate if our client is meant to create requests for testing
// and auth to indicate the authentication method we want to use.
func NewConfig(t bool, auth string) *Config {
	return &Config{
		testing: t,
		auth:    auth,
	}
}
