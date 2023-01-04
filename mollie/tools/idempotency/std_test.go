package idempotency

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewStdGenerator(t *testing.T) {
	assert.Implements(t, new(KeyGenerator), NewStdGenerator())
}

func TestStandardGenerator(t *testing.T) {
	key := NewStdGenerator().Generate()

	assert.Len(t, key, 36)

	assert.NotEqual(t, key, NewStdGenerator().Generate())
}
