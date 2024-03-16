package mollie

import (
	"encoding/json"
)

// ContextValues is a map of TransactionType to ContextValue.
type ContextValues map[TransactionType]ContextValue

// UnmarshalJSON implements the json.Unmarshaler interface on ContextValues.
//
// See: https://github.com/VictorAvelar/mollie-api-go/issues/251
func (cv *ContextValues) UnmarshalJSON(data []byte) error {
	var d map[TransactionType]ContextValue

	if err := json.Unmarshal(data, &d); err != nil {
		if _, ok := err.(*json.UnmarshalTypeError); ok {
			*cv = make(ContextValues)

			return nil
		}

		return err
	}

	*cv = ContextValues(d)

	return nil
}
