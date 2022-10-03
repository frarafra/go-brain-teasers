package common

import (
	"fmt"
)

func Rune() {
	msg := "π = 3.14159265358..."
	fmt.Printf("%T ", msg[0])
	for _, c := range msg {
		fmt.Printf("%T\n", c)
		break
	}
}
