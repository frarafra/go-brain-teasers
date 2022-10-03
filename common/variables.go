package common

import (
	"fmt"
)

func nextID(id *int) int {
	*id++
	return *id
}

func Variables() {
	var id int
	fmt.Println(id)
	nextID(&id)
	fmt.Println(id)
}
