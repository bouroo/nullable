package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
	"time"
)

// Time is a type that can be used to represent a nullable time.Time value in a struct.
// It is similar to sql.NullTime, but provides additional functionality.
type Time struct {
	sql.NullTime      // The underlying time.Time value.
	Present      bool // Whether the value is present or not.
}

// TimeOf creates a new Time value from a time.Time value.
//
// t is the time.Time value to wrap.
// Returns a new Time value with the given time.Time value and Present set to true.
func TimeOf(t time.Time) Time {
	return Time{sql.NullTime{Time: t, Valid: true}, true}
}

// MarshalJSON implements the json.Marshaler interface.
//
// It marshals the Time value into JSON.
// ([]byte, error)
func (v Time) MarshalJSON() ([]byte, error) {
	if !v.Present {
		return nil, nil
	}
	if !v.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(v.Time)
}

// UnmarshalJSON unmarshals the JSON data into the Time value.
//
// data is the JSON data to be unmarshaled.
// Returns an error if the unmarshaling process fails.
func (v *Time) UnmarshalJSON(data []byte) error {
	v.Valid = false
	v.Present = false
	if len(data) == 0 {
		return nil
	} else if bytes.Equal(data, []byte("null")) {
		v.Present = true
		return nil
	}
	if err := json.Unmarshal(data, &v.Time); err != nil {
		return err
	}
	v.Valid = true
	v.Present = true
	return nil
}
