package common

import (
	"fmt"
)

func Append() {
	a := []int{1, 2, 3}
	b := append(a[:1], 10)
	c := append(b[:1], 10, 11, 12)
	fmt.Printf("a=%v, b=%v, c=%v\n", a, b, c)
}
