// Ṕackage isoweek calculates the starting date and time of ISO 8601 week.
//
// The Go standard library "time" package has ISOWeek() method for getting
// ISO week number of given time.Time, but there is no reverse functionality
// for getting the date from week number. This package implements that.
//
// Invalid week numbers are silently accepted. There is a separate Validate()
// function if validation is needed.
package isoweek

import "time"

// ISOWeekday returns the ISO weekday number of given day.
// (1 = Mon, 2 = Tue,.. 7 = Sun)
//
// This is different from Go's standard time.Weekday, which you should use
// normally. It is exposed because it may be useful for some calculations.
func ISOWeekday(year int, month time.Month, day int) (weekday int) {
	return DateToJulian(year, month, day)%7 + 1
}

// startOffset returns the offset (in days) from the start of a year to
// Monday of the given week. Offset may be negative.
func startOffset(y, week int) (offset int) {
	y = y - 1
	return week*7 - (y+y/4-y/100+y/400+3)%7 - 4
	// return week*7 - ISOWeekday(y, 1, 4) - 3
}

// StartTime returns the starting time (Monday 00:00) of the given ISO week.
func StartTime(wyear, week int, loc *time.Location) (start time.Time) {
	y, m, d := StartDate(wyear, week)
	return time.Date(y, m, d, 0, 0, 0, 0, loc)
}

// StartDate returns the starting date (Monday) of the given ISO week.
func StartDate(wyear, week int) (year int, month time.Month, day int) {
	return JulianToDate(
		DateToJulian(wyear, 1, 1) + startOffset(wyear, week))
}

// Validate checks if a week number is valid. Returns true if it is valid.
func Validate(wyear, week int) (ok bool) {
	if week < 1 || week > 53 {
		return false
	}
	wyear2, week2 := StartTime(wyear, week, time.UTC).ISOWeek()

	if wyear == wyear2 && week == week2 {
		return true
	}
	return false
}
