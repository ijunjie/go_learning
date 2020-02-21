package ch5

import "testing"

const (
	SecondsPerMinute = 60
	SecondsPerHour = SecondsPerMinute * 60
	SecondsPerDay = SecondsPerHour * 24
)

func resolveTime(seconds int) (day int, hour int, minute int) {
	day = seconds / SecondsPerDay
	hour = seconds / SecondsPerHour
	minute = seconds / SecondsPerMinute
	return
}

func TestSeconds(t *testing.T) {
	t.Log(resolveTime(1000))
	_, hour, minute := resolveTime(18000)
	t.Log(hour, minute)

	day, _, _ := resolveTime(90000)
	t.Log(day)

}