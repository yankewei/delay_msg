package delay_msg

import (
	"testing"
	"time"
)

type UniTest struct {

}

func TestParseTime(t *testing.T) {
	r := InitMessageRing()
	duration, _ := time.ParseDuration("10m")
	cycles, seconds := ParseTime(r, duration)
	if cycles != 0 {
		t.Errorf("want cycles is 0, got %d", cycles)
	}
	if seconds != 600 {
		t.Errorf("want seconds is 600s, go %ds", seconds)
	}
}

func TestRing_AddMessage(t *testing.T) {
	ring := GetDelayMessageApp()
	object := new(UniTest)
	duration, _ := time.ParseDuration("10m")
	ring.AddMessage(object, "returnHello", duration)
	messageList := ring.Ring[600]
	message := messageList.List[0]
	if message.Method != "returnHello" {
		t.Errorf("want method is %s, got is %s", "returnHello", message.Method)
	}
}


func (u *UniTest)returnHello() string {
	return "Hello World"
}