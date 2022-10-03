package common

import (
	"fmt"
	"sync"
)

func Mutex() {
	var count int
	var wg sync.WaitGroup
	var m sync.Mutex

	for i := 0; i < 1_000_000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			count++
		}()

	}
	wg.Wait()
	fmt.Println(count)
}
