package delay_msg

import (
	"log"
	"reflect"
	"time"
)

/**
 * 消息环
 */
type Ring struct {
	Ring [3600]*MessageList
	P int
}

/**
 * 环上的每个元素都是一个切片
 */
type MessageList struct {
	List []*Message
}

/**
 * 环上每个切片的元素
 */
type Message struct {
	Object interface{}
	Method string
	Args   []interface{}
	Cycles int
}

/**
 * 项目应该暴露的几个Api
 * Run() 延迟消息的启动
 * AddMessage() 添加一条消息
 */

func (r *Ring) Run() {
	ticker := time.NewTicker(time.Second)
	for t := range ticker.C {
		go Process(r.Ring[r.P], t)
		r.P++
	}
}

// 初始化消息环，默认3600个长度，后期可能会自定义
func InitMessageRing() *Ring {
	r := new(Ring)
	for i := 0; i < 3600; i++ {
		var m []*Message
		r.Ring[i] = &MessageList{List: m}
	}
	r.P = 0
	return r
}

// 返回一个延迟消息
func GetDelayMessageApp() *Ring {
	return InitMessageRing()
}

/**
 * 解析时间，使消息可以放到正确的位置上
 */
func ParseTime(r *Ring, dalayTime time.Duration) (cycles, seconds int) {
	// 得到需要延迟多少秒执行
	seconds = r.P + int(dalayTime.Seconds())
	// 是否需要一小时后执行
	cycles = seconds / 3600
	seconds = seconds % 3600
	return
}

/**
 * object struct实例，可以通过 new Struct() 生成
 * method string struct的一个方法
 * delayTime time.Duration 要延迟多长时间执行
 * args 调用method所需要的参数
 */
func (r *Ring) AddMessage(object interface{}, method string, delayTime time.Duration, args ...interface{}) {
	cycles, seconds := ParseTime(r, delayTime)
	message := &Message{
		Object: object,
		Method: method,
		Args: args,
		Cycles: cycles,
	}
	messageList := r.Ring[r.P + seconds]
	if len(messageList.List) == 0 {
		messageList.List = append(messageList.List, message)
	}
}

func Process(messageList *MessageList, t time.Time) {
	if len(messageList.List) != 0 {
		for _, m := range messageList.List {
			if m.Cycles == 0 {
				go execMessage(m, t)
			} else {
				m.Cycles--
			}
		}
	}
}

func execMessage(m *Message, t time.Time) {
	inputs := make([]reflect.Value, len(m.Args))
	for i, v := range m.Args {
		inputs[i] = reflect.ValueOf(v)
	}
	reflect.ValueOf(m.Object).MethodByName(m.Method).Call(inputs)
	log.Println("execMessage: 时间:" + t.String())
}
