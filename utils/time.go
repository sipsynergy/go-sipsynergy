package utils

import "time"

func UTCNow() time.Time {
	return UTCTime(time.Now())
}

func UTCTime(t time.Time) time.Time {
	location, _ := time.LoadLocation("UTC")
	return t.In(location)
}
