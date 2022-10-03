package common

import (
	"fmt"
	"strconv"
)

func Parse() {
	i := 169
	s := fmt.Sprint(i)
	fmt.Println(s)
	s = strconv.Itoa(i)
	fmt.Println(s)
	s = string(i)
	fmt.Println(s)
}
