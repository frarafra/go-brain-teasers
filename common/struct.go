package common

import (
	"fmt"
	"time"
)

// Log is a log message
type Log struct {
	Message string
	Time    time.Time
}

func Struct() {
	ts := time.Date(2009, 11, 10, 0, 0, 0, 0, time.UTC)
	log := Log{"Hello", ts}
	fmt.Printf("%v\n", log)
}
