package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
		t, err := time.Parse("1/02/2006 15:04:05", date)
		if err != nil {
			panic(err)
		}
		return t
	}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	layout := "January 2, 2006 15:04:05"
	now := time.Now()
	date_passed , err := time.Parse(layout,date)
	if err != nil {
		panic(err)
	}
	return now.After(date_passed)
}
// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	layout :="Monday, January 2, 2006 15:04:05"
	date_passed, err := time.Parse(layout,date)
	if err != nil {
		panic(err)
	}

	if date_passed.Hour() >= 12 && date_passed.Hour() <18 {
		return true
	} else {
		return false
	}
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	layout := "1/2/2006 15:04:05"
	d,  err := time.Parse(layout, date)
	if err != nil {
		panic(err)
	}
	text := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", d.Weekday(), d.Month(), d.Day(), d.Year(), d.Hour(), d.Minute())
	return text 
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	layout := "2006-01-02 15:04:05 +0000"
	date := "2020-09-15 00:00:00 +0000"
	fecha, err := time.Parse(layout,date)
	res := time.Now().Year() - fecha.Year()
	if err != nil {
		panic(err)
	}
	return fecha.AddDate(res, 0, 0)
}
