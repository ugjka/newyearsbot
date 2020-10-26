package nyb

import (
	"strings"
	"time"

	"github.com/hako/durafmt"
)

func zoneOffset(target time.Time, zone *time.Location) int {
	_, offset := time.Date(target.Year(), target.Month(), target.Day(),
		target.Hour(), target.Minute(), target.Second(),
		target.Nanosecond(), zone).Zone()
	return offset
}

var timeNow = func() time.Time {
	return time.Now()
}

func humanDur(d time.Duration) string {
	h := durafmt.Parse(d)
	arr := strings.Split(h.String(), " ")
	if len(arr) > 2 {
		return strings.Join(arr[:4], " ")
	}
	return strings.Join(arr[:2], " ")
}

func normalize(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)
	split := strings.Split(s, " ")
	s = ""
	for i, w := range split {
		if w == "" {
			continue
		}
		s += w
		if i != len(split)-1 {
			s += " "
		}
	}
	return s
}
