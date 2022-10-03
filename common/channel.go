package common

import (
	"context"
	"fmt"
)

func fibs(ctx context.Context, n int) chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		a, b := 1, 1
		for i := 0; i < n; i++ {
			select {
			case ch <- a:
				a, b = b, a+b
			case <-ctx.Done():
				fmt.Println("cancelled")
			}
		}
	}()

	return ch
}

func Channel() {
	ctx, cancel := context.WithCancel(context.Background())
	for i := range fibs(ctx, 5) {
		fmt.Printf("%d ", i)
	}
	fmt.Println()
	ch := fibs(ctx, 5)
	for i := 0; i < 3; i++ {
		val := <-ch
		fmt.Printf("%d ", val)
	}
	fmt.Println()
	cancel()
}
