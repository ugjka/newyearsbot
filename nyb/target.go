package nyb

import "time"

//Set target year
var target = func() time.Time {
	tmp := time.Now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	return time.Date(tmp.Year(), time.February, 15, 0, 0, 0, 0, time.UTC)
	//return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

var timeNow = func() time.Time {
	return time.Now()
}
