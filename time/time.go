package time

import "time"

const SQLDateFormat = "2006-01-02 15:04:05"

func UTCNow() time.Time {
	return UTCTime(time.Now())
}

func UTCTime(t time.Time) time.Time {
	location, _ := time.LoadLocation("UTC")
	return t.In(location)
}
