package delay_msg

import (
	"time"
)

type Message struct {
	Object interface{}
	Method string
	Args   interface{}
	Cycles int
}

var m Message

/**
 * object struct实例，可以通过 new Struct() 生成
 * method string struct的一个方法
 * delayTime time.Duration 要延迟多长时间执行
 * args 调用method所需要的参数
 */
func AddItemToMessage(object interface{}, method string, delayTime time.Duration, args ...interface{}) Message {
	m.Object = object
	m.Method = method
	m.Args = args
	m.Cycles, _ = ParseTime(delayTime)
	return m
}

func ParseTime(dalayTime time.Duration) (cycles, seconds int) {
	// 得到需要延迟多少秒执行
	seconds = int(dalayTime.Seconds())
	// 是否需要一小时后执行
	cycles = seconds / 3600
	seconds = seconds % 3600
	return
}
