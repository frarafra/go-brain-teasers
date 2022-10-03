package common

import (
	"fmt"

	"golang.org/x/text/unicode/norm"
)

func Norm() {
	city1, city2 := "Krakow", "Kraków"
	city1, city2 = norm.NFC.String(city1), norm.NFC.String(city2)
	fmt.Println(city1 == city2)
}
