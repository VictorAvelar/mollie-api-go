package mollie

import (
	"encoding/json"
	"testing"

	"github.com/VictorAvelar/mollie-api-go/v3/testdata"
	"github.com/stretchr/testify/assert"
)

func TestBaseError(t *testing.T) {
	cases := []struct {
		name      string
		body      []byte
		wantField bool
	}{
		{
			"standard base error",
			[]byte(testdata.UnauthorizedErrorResponse),
			false,
		},
		{
			"error with field included",
			[]byte(testdata.UnprocessableEntityErrorResponse),
			true,
		},
	}

	for _, c := range cases {
		g := &BaseError{}
		t.Run(c.name, func(t *testing.T) {
			err := json.Unmarshal(c.body, g)
			assert.Nil(t, err)
			if c.wantField {
				assert.Contains(t, g.Error(), "field")
			} else {
				assert.NotContains(t, g.Error(), "field")
			}
		})
	}
}
