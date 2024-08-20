package nullable

import (
	"bytes"
	"database/sql"
	"encoding/json"
)

// Value is a type that can be used to represent a nullable value in a struct.
// It is similar to sql.Null[Type], but provides additional functionality.
type Value[T any] struct {
	sql.Null[T]      // The underlying value.
	Present     bool // Whether the value is present or not.
}

// ValueOf creates a new Value from the given value.
//
// v is the value to wrap.
// It sets the Present field to true.
func ValueOf[T any](v T) Value[T] {
	return Value[T]{sql.Null[T]{V: v, Valid: true}, true}
}

// MarshalJSON implements the json.Marshaler interface.
//
// It marshals the Value value into JSON.
// ([]byte, error)
func (v Value[T]) MarshalJSON() ([]byte, error) {
	if !v.Present {
		return nil, nil
	}
	if !v.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(v.V)
}

// UnmarshalJSON unmarshals the JSON data into the Value value.
//
// data is the JSON data to be unmarshaled.
// Returns an error if the unmarshaling process fails.
func (v *Value[T]) UnmarshalJSON(data []byte) error {
	v.Valid = false
	v.Present = false
	if len(data) == 0 {
		return nil
	}
	if bytes.Equal(data, []byte("null")) {
		v.Present = true
		return nil
	}
	if err := json.Unmarshal(data, &v.V); err != nil {
		return err
	}
	v.Valid = true
	v.Present = true
	return nil
}
