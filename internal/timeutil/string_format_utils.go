package timeutil

import "time"

// FormatDateTimeForURL returns the date in the format YYYY-MM-DD_HH-MM-SS.
func FormatDateTimeForURL(date time.Time) string {
	return date.Format("2006-01-02_15-04-05")
}

// FormatDateTimeShort returns the date in the format YYYY-MM-DD HH:MM.
func FormatDateTimeShort(date time.Time) string {
	return date.Format("2006-01-02 15:04")
}

// FormatYYYYMMDD returns the date in the format YYYY-MM-DD.
func FormatYYYYMMDD(date time.Time) string {
	return date.Format("2006-01-02")
}

// FormatYYYYMMDD returns the date in the format YYYY/MM/DD.
func FormatYYYYMMDDSlashed(date time.Time) string {
	return date.Format("2006/01/02")
}

// FormatDDMMYYYYSlashed returns the date in the format MM/DD/YYYY
func FormatDDMMYYYYSlashed(date time.Time) string {
	return date.Format("02/01/2006")
}
