# nullable

A library for nullable types in Go.

The library provides two types:

- `Time`: a nullable `time.Time` type
- `Value[T]`: a nullable type of type `T`

The `Time` type can be used to represent a nullable `time.Time` value in a struct.
The `Value[T]` type can be used to represent a nullable value of type `T` in a struct.

The `Time` and `Value[T]` types implement the `json.Marshaler` and `json.Unmarshaler` interfaces,
so they can be used with the `encoding/json` package.

The `Time` and `Value[T]` types also implement the `sql.Scanner` and `driver.Valuer` interfaces,
so they can be used with the `database/sql` package.

The `Time` type has two fields: `NullTime` and `Present`.
The `NullTime` field is a `sql.NullTime` type, and the `Present` field is a `bool` type.
The `Present` field is used to indicate whether the `NullTime` field is valid or not.

The `Value[T]` type has two fields: `Null` and `Present`.
The `Null` field is a `sql.Null[Type]` type, and the `Present` field is a `bool` type.
The `Present` field is used to indicate whether the `Null` field is valid or not.

The `Time` and `Value[T]` types have several methods:

- `MarshalJSON() ([]byte, error)`: marshals the `Time` or `Value[T]` value into JSON
- `UnmarshalJSON(data []byte) error`: unmarshals the JSON data into the `Time` or `Value[T]` value
- `Scan(value interface{}) error`: scans the value into the `Time` or `Value[T]` value
- `Value() (driver.Value, error)`: returns the value of the `Time` or `Value[T]` value

The `Time` and `Value[T]` types also have several functions:

- `TimeOf(t time.Time) Time`: creates a new `Time` value from a `time.Time` value
- `ValueOf[T](v T) Value[T]`: creates a new `Value[T]` value from a value of type `T`

The `Time` and `Value[T]` types can be used with the `database/sql` package to scan nullable values from a database,
and to marshal nullable values into JSON.

The `Time` and `Value[T]` types can be used with the `encoding/json` package to marshal nullable values into JSON,
and to unmarshal nullable values from JSON.
