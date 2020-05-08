package delay_msg

import (
	"github.com/yankewei/delay_msg/mine"
	"reflect"
	"testing"
)

func TestMessage(t *testing.T) {
	object := new(mine.MyStruct)
	m := AddItemToMessage(object, "Add", 1, 2)
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(1)
	inputs[1] = reflect.ValueOf(2)
	ret := reflect.ValueOf(m.Object).MethodByName(m.Method).Call(inputs)
	got := ret[0].Int()
	want := 3
	if got != int64(want) {
		t.Errorf("want is %d, go is %d", want, got)
	}
}