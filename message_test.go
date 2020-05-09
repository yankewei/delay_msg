package delay_msg

import (
	"testing"
	"time"
)

func TestParseTime(t *testing.T) {
	duration, _ := time.ParseDuration("10m")
	cycles, seconds := ParseTime(duration)
	if cycles != 0 {
		t.Errorf("want cycles is 0, got %d", cycles)
	}
	if seconds != 600 {
		t.Errorf("want seconds is 600s, go %ds", seconds)
	}
}
