package nullable

import (
	"database/sql"
	"testing"
	"time"
)

const MOCK_TIME_STRING = "2022-01-01T12:00:00+07:00"

func TestTimeOf(t *testing.T) {
	testTime, _ := time.Parse(time.RFC3339, MOCK_TIME_STRING)
	tests := []struct {
		name string
		time time.Time
		want Time
	}{
		{
			name: "valid time",
			time: testTime,
			want: Time{
				NullTime: sql.NullTime{
					Time:  testTime,
					Valid: true,
				},
				Present: true,
			},
		},
		{
			name: "zero time",
			time: time.Time{},
			want: Time{
				NullTime: sql.NullTime{
					Time:  time.Time{},
					Valid: true,
				},
				Present: true,
			},
		},
		{
			name: "specific date and time",
			time: time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
			want: Time{
				NullTime: sql.NullTime{
					Time:  time.Date(2022, 1, 1, 12, 0, 0, 0, time.UTC),
					Valid: true,
				},
				Present: true,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := TimeOf(tt.time)
			if got != tt.want {
				t.Errorf("TimeOf() = %+v, want %+v", got, tt.want)
			}
		})
	}
}

func TestTime_MarshalJSON(t *testing.T) {
	testTime, _ := time.Parse(time.RFC3339, MOCK_TIME_STRING)
	tests := []struct {
		name      string
		timeValue Time
		want      []byte
		wantErr   bool
	}{
		{
			name:      "Time value is not present",
			timeValue: Time{Present: false},
			want:      []byte{},
			wantErr:   false,
		},
		{
			name:      "Time value is present but not valid",
			timeValue: Time{sql.NullTime{Valid: false}, true},
			want:      []byte("null"),
			wantErr:   false,
		},
		{
			name:      "Time value is present and valid",
			timeValue: Time{sql.NullTime{Time: testTime, Valid: true}, true},
			want:      []byte(`"` + testTime.Format(time.RFC3339Nano) + `"`),
			wantErr:   false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.timeValue.MarshalJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("Time.MarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if string(got) != string(tt.want) {
				t.Errorf("Time.MarshalJSON() = %s, want %s", got, tt.want)
			}
		})
	}
}

func TestTime_UnmarshalJSON(t *testing.T) {
	tests := []struct {
		name    string
		data    []byte
		wantErr bool
	}{
		{"empty data", []byte{}, false},
		{"null data", []byte("null"), false},
		{"valid data", []byte(`"2022-01-01T12:00:00Z"`), false},
		{"invalid data", []byte(`"invalid"`), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := &Time{}
			if err := v.UnmarshalJSON(tt.data); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalJSON() error = %v, wantErr %v", err, tt.wantErr)
			}
			if !tt.wantErr {
				if tt.name == "empty data" {
					if v.Present || v.Valid {
						t.Errorf("UnmarshalJSON() = %+v, want %+v", v, &Time{})
					}
				} else if tt.name == "null data" {
					if !v.Present || v.Valid {
						t.Errorf("UnmarshalJSON() = %+v, want %+v", v, &Time{Present: true})
					}
				} else if tt.name == "valid data" {
					wantTime, _ := time.Parse(time.RFC3339, "2022-01-01T12:00:00Z")
					if v.Time != wantTime || !v.Present || !v.Valid {
						t.Errorf("UnmarshalJSON() = %+v, want %+v", v, &Time{sql.NullTime{Time: wantTime, Valid: true}, true})
					}
				}
			}
		})
	}
}
