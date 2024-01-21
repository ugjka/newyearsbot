package nyb

import (
	"strings"
	"time"

	"github.com/hako/durafmt"
)

func zoneOffset(target time.Time, zone *time.Location) time.Duration {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return time.Second * time.Duration(offset)
}

func humanDur(d time.Duration) string {
	hdur := durafmt.Parse(d)
	hdur = hdur.LimitToUnit("weeks")
	hdur = hdur.LimitFirstN(2)
	return hdur.String()
}

func normalize(s string) string {
	s = strings.ToLower(s)
	s = strings.Join(strings.Fields(s), " ")
	return s
}
