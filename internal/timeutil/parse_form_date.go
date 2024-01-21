package timeutil

import "time"

// convertDateStringToDate takes a date string in the format
// YYYY-MM-DD and converts it to time.Time
func ParseFormDate(dateString string) (time.Time, error) {
	layout := "2006-01-02"
	date, err := time.Parse(layout, dateString)
	if err != nil {
		return time.Time{}, err
	}
	return date, nil
}
