package common

import (
	"fmt"
)

func FileExists(path string) (bool, error) {
	var err error
	return false, err
}

func Error() {
	if _, err := FileExists("/no/such/file"); err != nil {
		fmt.Printf("error: %s\n", err)
	} else {
		fmt.Println("OK")
	}
}
