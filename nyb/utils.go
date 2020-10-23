package nyb

import (
	"strings"
	"time"

	"github.com/hako/durafmt"
)

func roundDuration(dur *durafmt.Durafmt) string {
	arr := strings.Split(dur.String(), " ")
	if len(arr) > 2 {
		return strings.Join(arr[:4], " ")
	}
	return strings.Join(arr[:2], " ")
}

func getOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}

var timeNow = func() time.Time {
	return time.Now()
}
