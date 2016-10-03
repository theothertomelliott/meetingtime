package describe

import (
	"errors"
	"testing"
	"time"

	"github.com/theothertomelliott/meetingtime"
)

func TestScheduleDescription(t *testing.T) {
	var tests = []struct {
		name        string
		schedule    meetingtime.Schedule
		expectedOut string
		expectedErr error
	}{
		{
			name:        "Daily",
			schedule:    meetingtime.NewDailySchedule(time.Date(2016, time.January, 1, 0, 0, 0, 0, time.UTC), 1),
			expectedOut: "Every day starting Fri Jan 01 2016 at 12:00AM",
		},
		{
			name:        "Every 5 days",
			schedule:    meetingtime.NewDailySchedule(time.Date(2016, time.January, 2, 0, 0, 0, 0, time.UTC), 5),
			expectedOut: "Every 5 days starting Sat Jan 02 2016 at 12:00AM",
		},
		{
			name:        "Weekly",
			schedule:    meetingtime.NewWeeklySchedule(time.Date(2016, time.January, 3, 0, 0, 0, 0, time.UTC), 1),
			expectedOut: "Every week starting Sun Jan 03 2016 at 12:00AM",
		},
		{
			name:        "Every 4 weeks",
			schedule:    meetingtime.NewWeeklySchedule(time.Date(2016, time.January, 4, 0, 0, 0, 0, time.UTC), 4),
			expectedOut: "Every 4 weeks starting Mon Jan 04 2016 at 12:00AM",
		},
		{
			name:        "Monthly",
			schedule:    meetingtime.NewMonthlySchedule(time.Date(2016, time.January, 5, 0, 0, 0, 0, time.UTC), 1),
			expectedOut: "Every month starting Tue Jan 05 2016 at 12:00AM",
		},
		{
			name:        "Every 6 months",
			schedule:    meetingtime.NewMonthlySchedule(time.Date(2016, time.January, 6, 0, 0, 0, 0, time.UTC), 6),
			expectedOut: "Every 6 months starting Wed Jan 06 2016 at 12:00AM",
		},
		{
			name:        "Yearly",
			schedule:    meetingtime.NewYearlySchedule(time.Date(2016, time.January, 7, 0, 0, 0, 0, time.UTC), 1),
			expectedOut: "Every year starting Thu Jan 07 2016 at 12:00AM",
		},
		{
			name:        "Every 6 years",
			schedule:    meetingtime.NewYearlySchedule(time.Date(2016, time.January, 8, 0, 0, 0, 0, time.UTC), 2),
			expectedOut: "Every 2 years starting Fri Jan 08 2016 at 12:00AM",
		},
		{
			name:        "Every 2nd Wednesday",
			schedule:    meetingtime.NewMonthlyScheduleByWeekday(time.Date(2016, time.October, 12, 0, 0, 0, 0, time.UTC)),
			expectedOut: "Every 2nd Wednesday, starting Oct 12 2016 at 12:00AM",
		},
		{
			name:        "Every 1st Friday",
			schedule:    meetingtime.NewMonthlyScheduleByWeekday(time.Date(2016, time.October, 7, 0, 0, 0, 0, time.UTC)),
			expectedOut: "Every 1st Friday, starting Oct 07 2016 at 12:00AM",
		},
		{
			name:        "Every 3rd Monday",
			schedule:    meetingtime.NewMonthlyScheduleByWeekday(time.Date(2016, time.October, 17, 0, 0, 0, 0, time.UTC)),
			expectedOut: "Every 3rd Monday, starting Oct 17 2016 at 12:00AM",
		},
		{
			name:        "Invalid type",
			schedule:    meetingtime.Schedule{Type: 100},
			expectedErr: errors.New("unknown schedule type"),
		},
	}
	for _, test := range tests {
		test := test
		t.Run(test.name, func(t *testing.T) {
			out, err := Schedule(test.schedule)
			if test.expectedErr != nil {
				if err == nil || test.expectedErr.Error() != err.Error() {
					t.Errorf("Error not as expected. Expected '%v', got '%v'", test.expectedErr, err)
				}
			}
			if out != test.expectedOut {
				t.Errorf("Description expected '%v', got '%v'", test.expectedOut, out)
			}
		})
	}
}
