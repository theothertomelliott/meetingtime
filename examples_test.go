package meetingtime

import (
	"fmt"
	"time"
)

func ExampleSchedule() {
	// Create a Schedule for a meeting that occurs every other month on the 10th at 6pm
	schedule := NewMonthlySchedule(time.Date(2016, time.January, 10, 18, 0, 0, 0, time.UTC), 2)

	// Get the second meeting by asking for the next meeting from 1 minute after the first
	secondMeeting, err := schedule.Next(time.Date(2016, time.January, 10, 18, 1, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(secondMeeting.Format(time.UnixDate))

	// Get the first meeting by asking for the meeting previous to the second
	first, err := schedule.Previous(secondMeeting)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(first.Format(time.UnixDate))

	// Create a Schedule for a daily meeting at 2.30pm
	schedule = NewDailySchedule(time.Date(2016, time.February, 1, 14, 30, 0, 0, time.UTC), 1)

	// An ErrNoEarlierMeetings error is returned if you try to get the meeting before the first
	_, err = schedule.Previous(time.Date(2016, time.February, 1, 14, 30, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
	}

	// An ErrNoEarlierMeetings error is also returned for calling previous on any earlier date
	_, err = schedule.Previous(time.Date(2016, time.January, 31, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
	}

	// Output: Thu Mar 10 18:00:00 UTC 2016
	// Sun Jan 10 18:00:00 UTC 2016
	// no meetings on or before this date
	// no meetings on or before this date
}

func ExampleScheduleSlice() {
	// Create a ScheduleSlice for a meeting on the 1st and 3rd Monday of each month at 7pm
	schedule := ScheduleSlice{
		NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 5, 19, 0, 0, 0, time.UTC)),
		NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 19, 19, 0, 0, 0, time.UTC)),
	}

	// Get the first meeting in October
	firstInOct, err := schedule.Next(time.Date(2016, time.October, 1, 0, 0, 0, 0, time.UTC))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(firstInOct.Format(time.UnixDate))

	// Get the next meeting after this
	secondInOct, err := schedule.Next(firstInOct)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(secondInOct.Format(time.UnixDate))
	// Output: Mon Oct  3 19:00:00 UTC 2016
	// Mon Oct 17 19:00:00 UTC 2016
}
