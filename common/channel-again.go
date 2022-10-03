package common

import (
	"fmt"
)

func ChannelAgain() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	a, ok := <-ch
	fmt.Println(a, ok)
	b, ok := <-ch
	fmt.Println(b, ok)
	close(ch)
	c, ok := <-ch
	fmt.Println(c, ok)
}
