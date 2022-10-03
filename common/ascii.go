package common

import (
	"fmt"
	"unicode/utf8"
)

func Ascii() {
	city := "Kraków"
	fmt.Println(len(city))
	fmt.Println(utf8.RuneCountInString(city))
}
