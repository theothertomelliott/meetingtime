package meetingtime

import (
	"errors"
	"time"
)

// Schedule defines a regular schedule for a meeting
type Schedule struct {
	Type      ScheduleType // Type of recurrence
	First     time.Time    // Time and date of first meeting
	Frequency uint         // How frequent the meeting is. For a daily meeting, 2 would mean every other day.
}

// ScheduleType specifies the way in which this schedule recurs
type ScheduleType uint8

const (
	// Daily specifies a meeting that recurs daily.
	Daily ScheduleType = iota
	// Weekly specifies a meeting that recurs weekly.
	Weekly
	// Monthly specifies a meeting that recurs monthly.
	Monthly
	// MonthlyByWeekday specifies a meeting that recurs on the nth weekday of the month (2nd Wednesday, for example), based on the first meeting date.
	MonthlyByWeekday
	// Yearly specifes a meeting that recurs yearly.
	Yearly
)

// NewDailySchedule creates a schedule recurring every n days
func NewDailySchedule(first time.Time, n uint) Schedule {
	return Schedule{Type: Daily, First: first, Frequency: n}
}

// NewWeeklySchedule creates a schedule recurring on the same day every n weeks
func NewWeeklySchedule(first time.Time, n uint) Schedule {
	return Schedule{Type: Weekly, First: first, Frequency: n}
}

// NewMonthlySchedule creates a schedule recurring on the specified day in the month, every n months.
func NewMonthlySchedule(first time.Time, n uint) Schedule {
	return Schedule{Type: Monthly, First: first, Frequency: n}
}

// NewMonthlyScheduleByWeekday creates a schedule recurring every month on the same day of the week as the first meeting (for example, the 2nd Wednesday).
func NewMonthlyScheduleByWeekday(first time.Time) Schedule {
	return Schedule{Type: MonthlyByWeekday, First: first, Frequency: 1}
}

// NewYearlySchedule creates a schedule recurring every n years
func NewYearlySchedule(first time.Time, n uint) Schedule {
	return Schedule{Type: Yearly, First: first, Frequency: n}
}

/*
Next returns the time of the next meeting after the given time.

For daily and yearly schedules, assumes that the given time is the date of the current meeting.
For monthly schedules, the closest valid date after the provided one will be returned.
*/
func (s Schedule) Next(t time.Time) (time.Time, error) {
	var err error
	c := s.First
	for c.Before(t) || c.Equal(t) {
		c, err = s.increment(c)
		if err != nil {
			return time.Time{}, err
		}
	}
	return c, nil
}

/*
Previous returns the time of the meeting before the given time.

For daily and yearly schedules, assumes that the given time is the date of the current meeting.
For monthly schedules, the closest valid date before the provided one will be returned.
*/
func (s Schedule) Previous(t time.Time) (time.Time, error) {
	var err error
	c := s.First
	prev := s.First
	for c.Before(t) {
		prev = c
		c, err = s.increment(c)
		if err != nil {
			return time.Time{}, err
		}
	}
	return prev, nil
}

func (s *Schedule) increment(t time.Time) (time.Time, error) {
	if s.Type == Daily {
		return t.AddDate(0, 0, int(s.Frequency)), nil
	}
	if s.Type == Weekly {
		return t.AddDate(0, 0, 7*int(s.Frequency)), nil
	}
	if s.Type == Monthly {
		return t.AddDate(0, int(s.Frequency), 0), nil
	}
	if s.Type == MonthlyByWeekday {
		// Identify the weekday and index
		weekday, n := getWeekdayAndIndex(s.First)
		c := t.AddDate(0, 0, 1)
		w, cn := getWeekdayAndIndex(c)
		for w != weekday || n != cn {
			c = c.AddDate(0, 0, 1)
			w, cn = getWeekdayAndIndex(c)
		}
		return c, nil
	}
	if s.Type == Yearly {
		return t.AddDate(int(s.Frequency), 0, 0), nil
	}
	return time.Time{}, errors.New("not implemented")
}

func getWeekdayAndIndex(t time.Time) (weekday time.Weekday, n int) {
	// Identify the weekday and index
	weekday = t.Weekday()
	n = 0
	d := t.AddDate(0, 0, -t.Day())
	for d.Before(t) {
		d = d.AddDate(0, 0, 1)
		if d.Weekday() == weekday {
			n++
		}
	}
	return weekday, n
}
