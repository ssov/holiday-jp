package holiday

import "time"

var jst = time.FixedZone("JST", 9*60*60)

func IsHoliday(t time.Time) bool {
	timeJST := t.In(jst)
	normalizedTime := time.Date(
		timeJST.Year(),
		timeJST.Month(),
		timeJST.Day(),
		0, 0, 0, 0, jst)
	if _, ok := holidays[normalizedTime]; !ok {
		return false
	}
	return true
}
