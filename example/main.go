package main

import (
	"github.com/yankewei/delay_msg"
	"github.com/yankewei/delay_msg/example/Mine"
	"time"
)

func main() {
	r := delay_msg.GetDelayMessageApp()
	m := new(Mine.Mine)
	r.AddMessage(m, "Hello", 10 * time.Second)
	r.Run()
}