package common

import (
	"fmt"
)

type Flag int

func Interface() {
	var i interface{} = 3
	f, ok := i.(Flag)
	if !ok {
		fmt.Println("not a Flag")
		return
	}
	fmt.Println(f)
}
