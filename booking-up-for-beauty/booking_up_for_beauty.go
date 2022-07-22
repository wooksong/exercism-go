package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	const layout string = "1/2/2006 15:04:05"
	t, err := time.Parse(layout, date)
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}
	return t
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	const layout string = "January 2, 2006 15:04:05"

	tDate, err := time.Parse(layout, date)
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	tNow := time.Now()

	return !tDate.After(tNow)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	const layout string = "Monday, January 2, 2006 15:04:05"
	tDate, err := time.Parse(layout, date)
	if err != nil { // Always check errors even if they should not happen.
		panic(err)
	}

	appointmentHour := tDate.Hour()
	return (appointmentHour >= 12) && (appointmentHour < 18)
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	tDate := Schedule(date)

	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.",
		tDate.Weekday().String(),
		tDate.Month().String(),
		tDate.Day(),
		tDate.Year(),
		tDate.Hour(),
		tDate.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), 9, 15, 0, 0, 0, 0, time.UTC)
}
