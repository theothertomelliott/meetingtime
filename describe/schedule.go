package describe

import (
	"errors"
	"fmt"
	"time"

	"github.com/theothertomelliott/meetingtime"
)

// Schedule generates an English description of an instance of meetingtime.Schedule
func Schedule(schedule meetingtime.Schedule) (string, error) {
	switch schedule.Type {
	case meetingtime.Daily:
		return daily(schedule), nil
	case meetingtime.Weekly:
		return weekly(schedule), nil
	case meetingtime.Monthly:
		return monthly(schedule), nil
	case meetingtime.MonthlyByWeekday:
		return monthlyByWeekday(schedule), nil
	case meetingtime.Yearly:
		return yearly(schedule), nil
	}
	return "", errors.New("unknown schedule type")
}

func daily(schedule meetingtime.Schedule) string {
	if schedule.Frequency == 1 {
		return fmt.Sprintf("Every day starting %v", formatDate(schedule.First))
	}
	return fmt.Sprintf("Every %d days starting %v", schedule.Frequency, formatDate(schedule.First))
}

func weekly(schedule meetingtime.Schedule) string {
	if schedule.Frequency == 1 {
		return fmt.Sprintf("Every week starting %v", formatDate(schedule.First))
	}
	return fmt.Sprintf("Every %d weeks starting %v", schedule.Frequency, formatDate(schedule.First))
}

func monthly(schedule meetingtime.Schedule) string {
	if schedule.Frequency == 1 {
		return fmt.Sprintf("Every month starting %v", formatDate(schedule.First))
	}
	return fmt.Sprintf("Every %d months starting %v", schedule.Frequency, formatDate(schedule.First))
}

func monthlyByWeekday(schedule meetingtime.Schedule) string {
	weekday, n := meetingtime.GetWeekdayAndIndex(schedule.First)
	return fmt.Sprintf("Every %v%v %v, starting %v", n, ordSuffix(n), weekday.String(), formatDateNoDay(schedule.First))
}

func yearly(schedule meetingtime.Schedule) string {
	if schedule.Frequency == 1 {
		return fmt.Sprintf("Every year starting %v", formatDate(schedule.First))
	}
	return fmt.Sprintf("Every %d years starting %v", schedule.Frequency, formatDate(schedule.First))
}

func formatDate(d time.Time) string {
	return d.Format("Mon Jan 02 2006 at 3:04PM")
}

func formatDateNoDay(d time.Time) string {
	return d.Format("Jan 02 2006 at 3:04PM")
}

func ordSuffix(x int) string {
	switch x % 10 {
	case 1:
		if x%100 != 11 {
			return "st"
		}
	case 2:
		if x%100 != 12 {
			return "nd"
		}
	case 3:
		if x%100 != 13 {
			return "rd"
		}
	}
	return "th"
}
