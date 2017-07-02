package timeformat

import (
	"strings"
	"time"
)

var (
	lm = []string{
		"YYYY", "2006",
		"YY", "06",
		"MMMM", "January",
		"MMM", "Jan",
		"MM", "01",
		"M", "1",
		"DD", "02",
		"D", "2",
		"hh", "03",
		"HH", "15",
		"h", "3",
		"wwww", "Monday",
		"www", "Mon",
		"mm", "04",
		"m", "4",
		"ss", "05",
		"s", "5",
		"f", "0",
		"F", "9",
		"a", "pm",
		"A", "PM",
		"z", "MST",
		"-Z:Z:Z", "-07:00:00",
		"Z:Z:Z", "Z07:00:00",
		"-Z:Z", "-07:00",
		"Z:Z", "Z07:00",
		"-ZZZ", "-070000",
		"ZZZ", "Z070000",
		"-ZZ", "-0700",
		"ZZ", "Z0700",
		"-Z", "-07",
		"Z", "Z07",
	}
	replacer *strings.Replacer
)

func init() {
	replacer = strings.NewReplacer(lm...)
}

// Time is wrapper around the default time.Time object which can be formatted
// using the custom layout
type Time struct {
	time.Time
}

// T returns a wrapper around the time.Time object which can be formatted
// using the custom layout
func T(t time.Time) Time {
	return Time{t}
}

// Format returns a textual representation of the time value formatted
// according to the custom layout
func (t Time) Format(s string) string {
	return t.Time.Format(replacer.Replace(s))
}

// Layout returns the parsed layout which can be passed to the Format
// method of time.Time
func Layout(s string) string {
	return replacer.Replace(s)
}
