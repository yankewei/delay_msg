# delay_msg
a delay message package build with go

想法来自于看的一篇微信公众号的文章[《如何快速实现“延时消息”》](https://mp.weixin.qq.com/s/9qt3JEvjv1wka57GBTng9g)。觉得还挺有趣的，就想实现一个看看。
### Usage
The first need Go installed (version 1.11+ is required), then you can use the below Go command to install:
```bash
go get -u github.com/yankewei/delay_msg
```

```go
package main

import (
	"fmt"
	"github.com/yankewei/delay_msg"
	"time"
)

type Hello struct {

}

func (h *Hello) Greet(name string) {
	fmt.Println("Welcome " + name)
}

func main() {
	r := delay_msg.GetDelayMessageApp()
	r.AddMessage(new(Hello), "Greet", 10 * time.Second, "yankewei")
	r.Run()
}
```
copy above code to a file, then run `go run example.go` and wait 10 seconds.
If correct, you can see a string "Welcome ***"。
