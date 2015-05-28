package timefns

import (
	"time"
)

func Now() {
	now := time.Now()
	println(now)
}

func Hello(arg interface{}) {
	println("hello", arg)
}
