package code_1

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

// context cancel 作用
func TestContext(t *testing.T) {
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx1 context.Context) {

		go func(ctx2 context.Context) {

			go func(ctx3 context.Context) {

				select {
				case <-ctx3.Done():
					fmt.Println("loop-3 done")
					return
				}

			}(ctx2)

			select {
			case <-ctx2.Done():
				fmt.Println("loop-2 done")
				return
			}

		}(ctx1)

		select {
		case <-ctx1.Done():
			fmt.Println("loop-1 done")
		}
	}(ctx)

	time.Sleep(2 * time.Second)
	cancel()
	fmt.Println("cancel done")
	time.Sleep(time.Second)
}

func TestContex(t *testing.T) {
	wg := sync.WaitGroup{}

	for i:=range 10 {
		wg.Add(1)
		go func(idx int) {
			name := idx
			defer wg.Done()
			var s []int
			for range 100 {
				s = append(s, idx)
				idx = idx + 10
			}
			fmt.Printf("Groutine%d, %v \n\n", name, s)
		}(i)

	}

	wg.Wait()
	fmt.Println("done")
}