package meetingtime

import (
	"testing"
	"time"
)

func TestNextScheduleSlice(t *testing.T) {
	var tests = []struct {
		name         string
		schedules    ScheduleSlice
		inTime       time.Time
		expectedTime time.Time
		expectedErr  error
	}{
		{
			name: "1 day",
			schedules: ScheduleSlice{
				Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			},
			inTime:       time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "2 days, non meeting day",
			schedules: ScheduleSlice{
				Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 2},
			},
			inTime:       time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "1 day - before first meeting",
			schedules: ScheduleSlice{
				Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			},
			inTime:       time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "1 week",
			schedules: ScheduleSlice{
				Schedule{Type: Weekly, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
				NewMonthlySchedule(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 1),
				NewMonthlyScheduleByWeekday(time.Date(2015, time.November, 11, 0, 0, 0, 0, time.UTC)),
			},
			inTime:       time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "2nd Wednesday and 3rd Tuesday",
			schedules: ScheduleSlice{
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 14, 18, 0, 0, 0, time.UTC)),
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 20, 18, 0, 0, 0, time.UTC)),
			},
			inTime:       time.Date(2016, time.September, 10, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.September, 14, 18, 0, 0, 0, time.UTC),
		},
		{
			name: "2nd Wednesday and 3rd Tuesday, test 2",
			schedules: ScheduleSlice{
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 14, 18, 0, 0, 0, time.UTC)),
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 20, 18, 0, 0, 0, time.UTC)),
			},
			inTime:       time.Date(2016, time.October, 13, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.October, 18, 18, 0, 0, 0, time.UTC),
		},
		{
			name: "Yearly schedules",
			schedules: ScheduleSlice{
				Schedule{Type: Yearly, First: time.Date(2014, time.January, 2, 0, 0, 0, 0, time.UTC), Frequency: 1},
				Schedule{Type: Yearly, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 2},
			},
			inTime:       time.Date(2017, time.January, 3, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2018, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outTime, outErr := test.schedules.Next(test.inTime)
			if outErr != nil {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}

func TestPreviousScheduleSlice(t *testing.T) {
	var tests = []struct {
		name         string
		schedules    ScheduleSlice
		inTime       time.Time
		expectedTime time.Time
		expectedErr  error
	}{
		{
			name: "1 day",
			schedules: ScheduleSlice{
				Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			},
			inTime:       time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "2 days, non meeting day",
			schedules: ScheduleSlice{
				NewDailySchedule(time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 2),
				Schedule{Type: Daily, First: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), Frequency: 1},
			},
			inTime:       time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "No earlier meetings, some valid",
			schedules: ScheduleSlice{
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 14, 18, 0, 0, 0, time.UTC)),
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 20, 18, 0, 0, 0, time.UTC)),
			},
			inTime:       time.Date(2016, time.September, 19, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.September, 14, 18, 0, 0, 0, time.UTC),
		},
		{
			name: "No earlier meetings, no valid",
			schedules: ScheduleSlice{
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 14, 18, 0, 0, 0, time.UTC)),
				NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 20, 18, 0, 0, 0, time.UTC)),
			},
			inTime:      time.Date(2016, time.September, 14, 14, 0, 0, 0, time.UTC),
			expectedErr: ErrNoEarlierMeetings,
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			outTime, outErr := test.schedules.Previous(test.inTime)
			if outErr != test.expectedErr {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}
