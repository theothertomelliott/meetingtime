# meetingtime

[![Build Status](https://travis-ci.org/theothertomelliott/meetingtime.svg?branch=master)](https://travis-ci.org/theothertomelliott/meetingtime)
[![GoDoc](https://godoc.org/github.com/theothertomelliott/meetingtime?status.svg)](https://godoc.org/github.com/theothertomelliott/meetingtime)
[![Go Report Card](https://goreportcard.com/badge/github.com/theothertomelliott/meetingtime)](https://goreportcard.com/report/github.com/theothertomelliott/meetingtime)
[![Coverage Status](https://coveralls.io/repos/github/theothertomelliott/meetingtime/badge.svg?branch=master)](https://coveralls.io/github/theothertomelliott/meetingtime?branch=master)

Package `meetingtime` provides tools for calculating dates and times for regularly occurring meetings.

# Basic Usage

Start by defining a Schedule:

    // Create a Schedule for a meeting that occurs every other month on the 10th at 6pm
    schedule := meetingtime.NewMonthlySchedule(time.Date(2016, time.January, 10, 18, 0, 0, 0, time.UTC), 2)

Then query the Schedule to obtain the date of the next meeting as a `time.Time` value:

    // Get the next meeting after the current time
    nextMeeting, err := schedule.Next(time.Now())

# Schedule Types

`meetingtime` provides a variety of schedule types:

* Daily
* Weekly
* Monthly
* Monthly by Weekday
* Yearly

*Daily*, *Weekly*, *Monthly* and *Yearly* schedules accept a frequency value, to allow for schedules such as "every other Monday". The *Monthly by Weekday* type permits schedules like "the second Tuesday of each month", this type does not take a frequency, and any frequency value will be ignored.

# Complex schedules

More complicated schedules can be represented by combinations of Schedule values using the ScheduleSlice type.

    // Create a ScheduleSlice for a meeting on the 1st and 3rd Monday of each month at 7pm
    schedule := ScheduleSlice{
        NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 5, 19, 0, 0, 0, time.UTC)),
        NewMonthlyScheduleByWeekday(time.Date(2016, time.September, 19, 19, 0, 0, 0, time.UTC)),
    }

    // Get the first meeting in October
    firstInOct, err := schedule.Next(time.Date(2016, time.October, 1, 0, 0, 0, 0, time.UTC))

# Describing a Schedule

The `describe` package provides a function (`describe`.`Schedule`) for creating English descriptions for `meetingtime`.`Schedule` values.

See the [describe godoc](https://godoc.org/github.com/theothertomelliott/meetingtime/describe) for more details.

`describe` does not currently have i18n support.

# Having Trouble?

If you're having trouble with `meetingtime`, please raise a [GitHub issue](https://github.com/theothertomelliott/meetingtime/issues) and we'll do what we can to help, or make fixes as needed.

# Contributing

Contributions are always more than welcome, feel free to pick up an issue and/or send a PR. Alternatively, just raise an issue with your idea for improvements.

# License

`meetingtime` is provided under the [MIT License](https://github.com/theothertomelliott/meetingtime/blob/master/LICENSE).
