package meetingtime

import (
	"testing"
	"time"
)

func TestNextDailySchedule(t *testing.T) {
	var tests = []struct {
		name         string
		daily        int
		inTime       time.Time
		expectedTime time.Time
	}{
		{
			name:         "1 day",
			daily:        1,
			inTime:       time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "365 days",
			daily:        365,
			inTime:       time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := NewDailySchedule(test.daily)
			if err != nil {
				t.Fatal(err)
			}
			outTime, outErr := s.Next(test.inTime)
			if outErr != nil {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}

func TestNextYearlySchedule(t *testing.T) {
	var tests = []struct {
		name         string
		daily        int
		inTime       time.Time
		expectedTime time.Time
	}{
		{
			name:         "1 year",
			daily:        1,
			inTime:       time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "5 years",
			daily:        5,
			inTime:       time.Date(2015, time.March, 20, 0, 0, 0, 0, time.UTC),
			expectedTime: time.Date(2020, time.March, 20, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := NewYearlySchedule(test.daily)
			if err != nil {
				t.Fatal(err)
			}
			outTime, outErr := s.Next(test.inTime)
			if outErr != nil {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}

func TestPreviousDailySchedule(t *testing.T) {
	var tests = []struct {
		name         string
		daily        int
		inTime       time.Time
		expectedTime time.Time
	}{
		{
			name:         "1 day",
			daily:        1,
			expectedTime: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			inTime:       time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "365 days",
			daily:        365,
			expectedTime: time.Date(2015, time.January, 1, 0, 0, 0, 0, time.UTC),
			inTime:       time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := NewDailySchedule(test.daily)
			if err != nil {
				t.Fatal(err)
			}
			outTime, outErr := s.Previous(test.inTime)
			if outErr != nil {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}

func TestPreviousYearlySchedule(t *testing.T) {
	var tests = []struct {
		name         string
		daily        int
		inTime       time.Time
		expectedTime time.Time
	}{
		{
			name:         "1 year",
			daily:        1,
			expectedTime: time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC),
			inTime:       time.Date(2017, time.January, 1, 0, 0, 0, 0, time.UTC),
		},
		{
			name:         "5 years",
			daily:        5,
			expectedTime: time.Date(2015, time.March, 20, 0, 0, 0, 0, time.UTC),
			inTime:       time.Date(2020, time.March, 20, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			s, err := NewYearlySchedule(test.daily)
			if err != nil {
				t.Fatal(err)
			}
			outTime, outErr := s.Previous(test.inTime)
			if outErr != nil {
				t.Errorf("error: %v", outErr.Error())
			} else if test.expectedTime != outTime {
				t.Errorf("times: expected '%v' got '%v'", test.expectedTime, outTime)
			}
		})
	}
}
