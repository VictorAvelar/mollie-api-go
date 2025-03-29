package idempotency

// NOpIdempotencyGenerator is a dummy implementation of the
// IdempotencyKeyGenerator interface.
//
// Good for testing or when a predictable result is required.
type nOpIdempotencyGenerator struct {
	expected string
}

// TestKeyExpected is the default value for the NOpGenerator.
const (
	TestKeyExpected = "test_ikg_key"
)

// Generate encapsulates the logic to return a string representation of
// a unique idempotency key.
func (nopIKG nOpIdempotencyGenerator) Generate() string {
	return nopIKG.expected
}

// NewNopGenerator returns a dummy implementation of the
// IdempotencyKeyGenerator interface.
//
// Good for testing or when a predictable result is required.
//
// If exp is an empty string, then TestKeyExpected is used as
// default value for the NOpGenerator.
func NewNopGenerator(exp string) KeyGenerator {
	if exp == "" {
		exp = TestKeyExpected
	}

	return nOpIdempotencyGenerator{
		expected: exp,
	}
}
