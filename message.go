package delay_msg

type Message struct {
	Object interface{}
	Method string
	Args interface{}
}

var m Message

func AddItemToMessage (object interface{}, method string, args ...interface{}) Message {
	m.Object = object
	m.Method = method
	m.Args = args
	return m
}