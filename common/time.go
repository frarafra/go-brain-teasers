package common

import (
	"fmt"
	"time"
)

func Time() {
	timeout := 3000
	fmt.Printf("before ")
	time.Sleep(time.Duration(timeout) * time.Millisecond)
	fmt.Println("after")
}
