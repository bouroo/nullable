package nullable

import (
	"bytes"
	"database/sql"
	"reflect"
	"testing"
)

func TestValueOf(t *testing.T) {
	tests := []struct {
		name string
		v    int
		want Value[int]
	}{
		{"valid value", 1, Value[int]{sql.Null[int]{V: 1, Valid: true}, true}},
		{"zero value", 0, Value[int]{sql.Null[int]{V: 0, Valid: true}, true}},
		{"non-nil pointer value", 1, Value[int]{sql.Null[int]{V: 1, Valid: true}, true}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ValueOf(tt.v)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ValueOf() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValue_MarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		value   Value[int]
		want    []byte
		wantErr bool
	}{
		{
			name: "not present",
			value: Value[int]{
				Null:    sql.Null[int]{Valid: false},
				Present: false,
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "present but not valid",
			value: Value[int]{
				Null:    sql.Null[int]{Valid: false},
				Present: true,
			},
			want:    []byte("null"),
			wantErr: false,
		},
		{
			name: "present and valid",
			value: Value[int]{
				Null:    sql.Null[int]{V: 42, Valid: true},
				Present: true,
			},
			want:    []byte("42"),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.value.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Value.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !bytes.Equal(got, tt.want) {
				t.Errorf("Value.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestValue_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"empty data", []byte{}, false},
		{"null data", []byte("null"), false},
		{"valid data", []byte(`"hello"`), false},
		{"invalid data", []byte(`invalid`), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Value[string]{}
			if err := v.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.name == "empty data" {
					if v.Present || v.Valid {
						t.Errorf("UnmarshalJSON() = %+v, want %+v", v, &Value[string]{})
					}
				} else if tt.name == "null data" {
					if !v.Present || v.Valid {
						t.Errorf("UnmarshalJSON() = %+v, want %+v", v, &Value[string]{Present: true})
					}
				} else if tt.name == "valid data" {
					wantValue := "hello"
					if v.V != wantValue || !v.Present || !v.Valid {
						t.Errorf("UnmarshalJSON() = %+v, want %+v", v, &Value[string]{sql.Null[string]{V: wantValue, Valid: true}, true})
					}
				}
			}
		})
	}
}
