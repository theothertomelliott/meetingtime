package meetingtime

import (
	"errors"
	"time"
)

// Schedule defines a regular schedule for a meeting
type Schedule struct {
	Daily         int
	Monthly       int
	MonthDay      int
	Weekday       *time.Weekday
	WeekdayNumber *int
	Yearly        int
}

// NewDailySchedule creates a schedule recurring every n days
func NewDailySchedule(n int) (Schedule, error) {
	return Schedule{Daily: n}, nil
}

func (s Schedule) IsDaily() bool { return s.Daily != 0 }

// NewWeeklySchedule creates a schedule recurring on the specified day of the week
func NewWeeklySchedule(weekday time.Weekday) (Schedule, error) {
	return Schedule{Weekday: &weekday}, nil
}

func (s Schedule) IsWeekly() bool { return s.Weekday != nil && s.WeekdayNumber == nil }

// NewMonthlySchedule creates a schedule recurring on the specified day in the month, every n months.
func NewMonthlySchedule(dayOfMonth int, n int) (Schedule, error) {
	return Schedule{MonthDay: dayOfMonth, Monthly: n}, nil
}

func (s Schedule) IsMonthly() bool { return s.Monthly != 0 }

// NewMonthlyScheduleByWeekday creates a schedule recurring on the nth weekday (eg: 2nd Tuesday)
func NewMonthlyScheduleByWeekday(weekday time.Weekday, n int) (Schedule, error) {
	return Schedule{Weekday: &weekday, WeekdayNumber: &n}, nil
}

func (s Schedule) IsMonthlyByWeekday() bool { return s.Weekday != nil && s.WeekdayNumber != nil }

// NewYearlySchedule creates a schedule recurring every n years
func NewYearlySchedule(n int) (Schedule, error) {
	return Schedule{Yearly: n}, nil
}

func (s Schedule) IsYearly() bool { return s.Yearly != 0 }

/*
Next returns the time of the next meeting after the given time.

For daily and yearly schedules, assumes that the given time is the date of the current meeting.
For monthly schedules, the closest valid date after the provided one will be returned.
*/
func (s Schedule) Next(t time.Time) (time.Time, error) {
	if s.isMultipleTypes() {
		return time.Time{}, errors.New("schedules of multiple types not supported")
	}
	if s.IsDaily() {
		return t.AddDate(0, 0, s.Daily), nil
	}
	if s.IsYearly() {
		return t.AddDate(s.Yearly, 0, 0), nil
	}
	return time.Time{}, errors.New("not implemented")
}

/*
Previous returns the time of the meeting before the given time.

For daily and yearly schedules, assumes that the given time is the date of the current meeting.
For monthly schedules, the closest valid date before the provided one will be returned.
*/
func (s Schedule) Previous(t time.Time) (time.Time, error) {
	if s.isMultipleTypes() {
		return time.Time{}, errors.New("schedules of multiple types not supported")
	}
	if s.IsDaily() {
		return t.AddDate(0, 0, -s.Daily), nil
	}
	if s.IsYearly() {
		return t.AddDate(-s.Yearly, 0, 0), nil
	}
	return time.Time{}, errors.New("not implemented")
}

func (s Schedule) isMultipleTypes() bool {
	typeCount := 0
	if s.IsDaily() {
		typeCount++
	}
	if s.IsWeekly() {
		typeCount++
	}
	if s.IsMonthly() {
		typeCount++
	}
	if s.IsMonthly() {
		typeCount++
	}
	if s.IsYearly() {
		typeCount++
	}
	return typeCount > 1
}
