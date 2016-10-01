package meetingtime

import (
	"errors"
	"time"
)

// ScheduleSlice allows Schedule instances to be grouped to create more complex schedules.
type ScheduleSlice []Schedule

/*
Next returns the earliest next meeting from all Schedules in the slice.
*/
func (schedules ScheduleSlice) Next(t time.Time) (time.Time, error) {
	if len(schedules) == 0 {
		return time.Time{}, errors.New("no schedules")
	}
	var next *time.Time
	for _, s := range schedules {
		sn, err := s.Next(t)
		if err != nil {
			return time.Time{}, err
		}
		if next == nil || sn.Before(*next) {
			next = &sn
		}
	}
	return *next, nil
}

/*
Previous returns the latest previous from all Schedules in the slice.
*/
func (schedules ScheduleSlice) Previous(t time.Time) (time.Time, error) {
	var next *time.Time
	for _, s := range schedules {
		sn, err := s.Previous(t)
		if err != nil {
			if err == ErrNoEarlierMeetings {
				continue
			}
			return time.Time{}, err
		}
		if next == nil || sn.After(*next) {
			next = &sn
		}
	}
	if next == nil {
		return time.Time{}, ErrNoEarlierMeetings
	}
	return *next, nil
}
