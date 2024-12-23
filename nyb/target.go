package nyb

import "time"

// Set the target year
var target = func() time.Time {
	tmp := now().UTC()
	if tmp.Month() == time.January && tmp.Day() < 2 {
		return time.Date(tmp.Year(), time.January, 1, 0, 0, 0, 0, time.UTC)
	}
	//Debug target
	//return time.Date(tmp.Year(), time.January, 20, 1, 0, 0, 0, time.UTC)
	return time.Date(tmp.Year()+1, time.January, 1, 0, 0, 0, 0, time.UTC)
}()

// Set now
var now = func() time.Time {
	return time.Now()
	// debug fake time
	//return time.Date(2024, 1, 1, 0, 0, 1, 100000011, time.Now().Location())
}
