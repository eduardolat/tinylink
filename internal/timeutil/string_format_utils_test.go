package timeutil

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFormatDateTimeForURL(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Known date 1",
			date:     time.Date(2022, 1, 1, 13, 45, 30, 0, time.UTC),
			expected: "2022-01-01_13-45-30",
		},
		{
			name:     "Known date 2",
			date:     time.Date(1999, 12, 31, 23, 59, 0, 0, time.UTC),
			expected: "1999-12-31_23-59-00",
		},
		{
			name:     "Single digit month and day",
			date:     time.Date(2000, 2, 3, 4, 5, 0, 0, time.UTC),
			expected: "2000-02-03_04-05-00",
		},
		{
			name:     "Zero value for time",
			date:     time.Time{},
			expected: "0001-01-01_00-00-00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, FormatDateTimeForURL(tt.date))
		})
	}
}

func TestFormatDateTimeShort(t *testing.T) {
	tests := []struct {
		name     string
		date     time.Time
		expected string
	}{
		{
			name:     "Known date 1",
			date:     time.Date(2022, 1, 1, 13, 45, 0, 0, time.UTC),
			expected: "2022-01-01 13:45",
		},
		{
			name:     "Known date 2",
			date:     time.Date(1999, 12, 31, 23, 59, 0, 0, time.UTC),
			expected: "1999-12-31 23:59",
		},
		{
			name:     "Single digit month and day",
			date:     time.Date(2000, 2, 3, 4, 5, 0, 0, time.UTC),
			expected: "2000-02-03 04:05",
		},
		{
			name:     "Zero value for time",
			date:     time.Time{},
			expected: "0001-01-01 00:00",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.expected, FormatDateTimeShort(tt.date))
		})
	}
}

func TestFormatYYYYMMDD(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want string
	}{
		{
			name: "Test date 1",
			date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			want: "2022-01-01",
		},
		{
			name: "Test date 2",
			date: time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
			want: "2000-12-31",
		},
		{
			name: "Test date 3",
			date: time.Date(1999, 6, 15, 0, 0, 0, 0, time.UTC),
			want: "1999-06-15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYYYYMMDD(tt.date); got != tt.want {
				t.Errorf("FormatYYYYMMDD() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatYYYYMMDDSlashed(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want string
	}{
		{
			name: "Test date 1",
			date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			want: "2022/01/01",
		},
		{
			name: "Test date 2",
			date: time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
			want: "2000/12/31",
		},
		{
			name: "Test date 3",
			date: time.Date(1999, 6, 15, 0, 0, 0, 0, time.UTC),
			want: "1999/06/15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatYYYYMMDDSlashed(tt.date); got != tt.want {
				t.Errorf("FormatYYYYMMDDSlashed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatDDMMYYYYSlashed(t *testing.T) {
	tests := []struct {
		name string
		date time.Time
		want string
	}{
		{
			name: "Test date 1",
			date: time.Date(2022, 1, 1, 0, 0, 0, 0, time.UTC),
			want: "01/01/2022",
		},
		{
			name: "Test date 2",
			date: time.Date(2000, 12, 31, 0, 0, 0, 0, time.UTC),
			want: "31/12/2000",
		},
		{
			name: "Test date 3",
			date: time.Date(1999, 6, 15, 0, 0, 0, 0, time.UTC),
			want: "15/06/1999",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatDDMMYYYYSlashed(tt.date); got != tt.want {
				t.Errorf("FormatDDMMYYYYSlashed() = %v, want %v", got, tt.want)
			}
		})
	}
}
