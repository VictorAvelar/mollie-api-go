package idempotency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNopGenerator(t *testing.T) {
	assert.Implements(t, new(KeyGenerator), NewNopGenerator("dummy"))
}

func TestGenerateDefault(t *testing.T) {
	g := NewNopGenerator("")

	assert.Equal(t, TestKeyExpected, g.Generate())
}

func TestGenerateNonDefault(t *testing.T) {
	g := NewNopGenerator("testing")

	assert.NotEqual(t, TestKeyExpected, g.Generate())
}
