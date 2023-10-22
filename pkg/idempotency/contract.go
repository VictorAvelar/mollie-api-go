package idempotency

// KeyGenerator describes the service in charge
// of generating a unique idempotency key to be passed in
// POST requests to ensure operation uniqueness.
//
// See: https://docs.mollie.com/overview/api-idempotency
type KeyGenerator interface {
	// Generate encapsulates the logic to return a string representation of
	// a unique idempotency key.
	Generate() string
}
