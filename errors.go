package meetingtime

type errorStr string

func (e errorStr) Error() string { return string(e) }

// ErrNoEarlierMeetings indicates that Previous was called with a date before the first meeting of a Schedule
const ErrNoEarlierMeetings = errorStr("no meetings on or before this date")
