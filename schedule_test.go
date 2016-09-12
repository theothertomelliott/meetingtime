package meetingtime

import (
	"testing"
	"time"
)

func TestNextSchedule(t *testing.T) {
	var tests = []struct {
		name         string
		schedule     Schedule
		inTime       time.Time
		expectedTime time.Time
		expectedErr  error
	}{
		{
			name:         "1 day",
			schedule:     Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			inTime:       time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2 days, non meeting day",
			schedule:     Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 2},
			inTime:       time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "1 day - before first meeting",
			schedule:     Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			inTime:       time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "1 week",
			schedule:     Schedule{Type: Weekly, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			inTime:       time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "1 month",
			schedule:     NewMonthlySchedule(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 1),
			inTime:       time.Date(2016, time.February, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.March, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2nd Wednesday",
			schedule:     NewMonthlyScheduleByWeekday(time.Date(2015, time.November, 11, 0, 0, 0, 0, time.UTC)),
			inTime:       time.Date(2016, time.September, 10, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.September, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2nd Wednesday - 5 minutes prior",
			schedule:     NewMonthlyScheduleByWeekday(time.Date(2015, time.November, 11, 18, 30, 0, 0, time.UTC)),
			inTime:       time.Date(2016, time.September, 14, 18, 25, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.September, 14, 18, 30, 0, 0, time.UTC),
		},
		{
			name:         "2nd Wednesday - 5 minutes after",
			schedule:     NewMonthlyScheduleByWeekday(time.Date(2015, time.November, 11, 18, 30, 0, 0, time.UTC)),
			inTime:       time.Date(2016, time.September, 14, 18, 35, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.October, 12, 18, 30, 0, 0, time.UTC),
		},
		{
			name:         "1 year",
			schedule:     Schedule{Type: Yearly, First: time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			inTime:       time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2 years, non meeting day",
			schedule:     Schedule{Type: Yearly, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 2},
			inTime:       time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outTime, outErr := test.schedule.Next(test.inTime)
			if outErr != nil {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}

func TestPreviousSchedule(t *testing.T) {
	var tests = []struct {
		name         string
		schedule     Schedule
		inTime       time.Time
		expectedTime time.Time
		expectedErr  error
	}{
		{
			name:         "1 day",
			schedule:     Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			inTime:       time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2 days, non meeting day",
			schedule:     NewDailySchedule(time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 2),
			inTime:       time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "1 week",
			schedule:     NewWeeklySchedule(time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 1),
			inTime:       time.Date(2016, time.January, 16, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 9, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "1 month",
			schedule:     NewMonthlySchedule(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 1),
			inTime:       time.Date(2016, time.March, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.February, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2nd Wednesday",
			schedule:     NewMonthlyScheduleByWeekday(time.Date(2015, time.November, 11, 0, 0, 0, 0, time.UTC)),
			inTime:       time.Date(2016, time.September, 20, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.September, 14, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "1 year",
			schedule:     NewYearlySchedule(time.Date(2014, time.January, 1, 0, 0, 0, 0, time.UTC), 1),
			inTime:       time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "2 years, non meeting day",
			schedule:     NewYearlySchedule(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 2),
			inTime:       time.Date(2018, time.January, 20, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:        "No earlier meeting",
			schedule:    Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			inTime:      time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedErr: ErrNoEarlierMeetings,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outTime, outErr := test.schedule.Previous(test.inTime)
			if outErr != test.expectedErr {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}

func TestGetWeekdayAndN(t *testing.T) {
	var tests = []struct {
		name            string
		time            time.Time
		expectedWeekday time.Weekday
		expectedN       int
	}{
		{
			name:            "First Monday",
			time:            time.Date(2016, time.September, 5, 0, 0, 0, 0, time.UTC),
			expectedWeekday: time.Monday,
			expectedN:       1,
		},
		{
			name:            "Third Monday",
			time:            time.Date(2016, time.September, 19, 0, 0, 0, 0, time.UTC),
			expectedWeekday: time.Monday,
			expectedN:       3,
		},
		{
			name:            "Second Wednesday",
			time:            time.Date(2015, time.November, 11, 0, 0, 0, 0, time.UTC),
			expectedWeekday: time.Wednesday,
			expectedN:       2,
		},
		{
			name:            "5th Sunday",
			time:            time.Date(2015, time.November, 29, 0, 0, 0, 0, time.UTC),
			expectedWeekday: time.Sunday,
			expectedN:       5,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			weekday, n := getWeekdayAndIndex(test.time)
			if weekday != test.expectedWeekday {
				t.Errorf("Weekday: expected %v, got %v", test.expectedWeekday, weekday)
			}
			if n != test.expectedN {
				t.Errorf("n: expected %v, got %v", test.expectedN, n)
			}
		})
	}
}
