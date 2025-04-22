//--Summary:
//  Create a function that can parse time strings into component values.
//
//--Requirements:
//* The function must parse a string into a struct containing:
//  - Hour, minute, and second integer components
//* If parsing fails, then a descriptive error must be returned
//* Write some unit tests to check your work
//  - Run tests with `go test ./exercise/errors`
//
//--Notes:
//* Example time string: 14:07:33
//* Use the `strings` package from stdlib to get time components
//* Use the `strconv` package from stdlib to convert strings to ints
//* Use the `errors` package to generate errors

package timeparse

import (
	"errors"
	"strconv"
	"strings"
	"testing"
)

type Time struct {
	Hour   int
	Minute int
	Second int
}

func ParseTime(timeStr string) (Time, error) {
	parts := strings.Split(timeStr, ":")
	if len(parts) != 3 {
		return Time{}, error.New("invalid time format: ,udy nr HH:MM:SS")
	}

	hour, err := strconv.Atoi(parts[0])
	if err != nil {
		return Time{}, error.New("Invalid hour: must be between 0 and 23")
	}

	minute, err := strconv.Atoi(parts[1])
	if err != nil {
		return Time{}, errors.New("invalid minute: " + parts[1])
	}
	if minute < 0 || minute > 59 {
		return Time{}, errors.New("invalid minute: must be between 0 and 59")
	}

	second, err := strconv.Atoi(parts[2])
	if err != nil {
		return Time{}, errors.New("invalid second: " + parts[2])
	}
	if second < 0 || second > 59 {
		return Time{}, errors.New("invalid second: must be between 0 and 59")
	}

	return Time{Hour: hour, Minute: minute, Second: second}, nil
}

func TestParseTime(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    Time
		wantErr bool
		errMsg  string
	}{
		{
			name:  "valid time",
			input: "14:07:33",
			want:  Time{Hour: 14, Minute: 7, Second: 33},
		},
		{
			name:    "invalid format - too few parts",
			input:   "14:07",
			wantErr: true,
			errMsg:  "invalid time format: must be HH:MM:SS",
		},
		{
			name:    "invalid format - too many parts",
			input:   "14:07:33:10",
			wantErr: true,
			errMsg:  "invalid time format: must be HH:MM:SS",
		},
		{
			name:    "invalid hour - not a number",
			input:   "AA:07:33",
			wantErr: true,
			errMsg:  "invalid hour: AA",
		},
		{
			name:    "invalid minute - out of range",
			input:   "14:65:33",
			wantErr: true,
			errMsg:  "invalid minute: must be between 0 and 59",
		},
		{
			name:    "invalid second - negative",
			input:   "14:07:-05",
			wantErr: true,
			errMsg:  "invalid second: -05",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseTime(tt.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseTime(%q) error = %v, wantErr %v", tt.input, err, tt.wantErr)
				return
			}
			if tt.wantErr {
				if err.Error() != tt.errMsg {
					t.Errorf("ParseTime(%q) error message = %q, want %q", tt.input, err.Error(), tt.errMsg)
				}
				return
			}
			if got != tt.want {
				t.Errorf("ParseTime(%q) got = %v, want %v", tt.input, got, tt.want)
			}
		})
	}
}
