package idempotency

import "github.com/google/uuid"

type stdGenerator struct{}

// Generate encapsulates the logic to return a string representation of
// a unique idempotency key.
//
// A string representation of a v4 uuid for this implementation.
func (ikg stdGenerator) Generate() string {
	return uuid.New().String()
}

// NewStdGenerator returns an standard and common way of generating
// idempotency unique keys.
func NewStdGenerator() KeyGenerator {
	return new(stdGenerator)
}
