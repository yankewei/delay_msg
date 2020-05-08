package delay_msg

import (
	"fmt"
	"github.com/yankewei/delay_msg/mine"
	"reflect"
)

func main() {
	object := new(mine.MyStruct)
	m := AddItemToMessage(object, "Add", 1, 2)
	inputs := make([]reflect.Value, 2)
	inputs[0] = reflect.ValueOf(1)
	inputs[1] = reflect.ValueOf(2)
	ret := reflect.ValueOf(m.Object).MethodByName(m.Method).Call(inputs)
	fmt.Println(ret)
}
