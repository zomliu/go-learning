package test

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

type Animal interface {
	GetName()
	GetAge()
}

type Person struct{}

func (p Person) GetName() {

}

func (p *Person) GetAge() {

}

func TestInterface(t *testing.T) {
	//var a Animal
	a := Person{}

	a.GetAge()
	a.GetName()
}

func TestContext(t *testing.T) {
	ctx := context.TODO()

	ctx1 := context.WithValue(ctx, "ctx1", "ctx1-value")
	ctx2 := context.WithValue(ctx1, "ctx2", "ctx2-value2")

	fmt.Println(ctx)
	fmt.Println(ctx1)
	fmt.Println(ctx2)

	fmt.Println(ctx2.Value("ctx1"))

	//context1()
}

func context1() {
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	for i := range 2 {
		go func(idx int) {
			fmt.Printf("go %d routine start... \n", idx)

			select {
			case <-ctx.Done():
				fmt.Println("Done")
			default:
				fmt.Println("nothing happen")
			}

		}(i)
	}
	// concelFunc()

	time.Sleep(6 * time.Second)
}

func TestAtomic(t *testing.T) {
	a := atomic.Int32{}
	a.Store(1)
	aa := a.Add(1)
	t.Logf("aa=%v", aa)
	t.Logf("a=%v", a.Load())
}

func TestCloseFunc(t *testing.T) {
	fc1 := counter(1)
	for range 3 {
		t.Log(fc1())
	}
	print("-------------\n")
	fc2 := fibonacci()
	for range 10 {
		t.Log(fc2())
	}
	print("-------------\n")
	next, reset := loop([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	for range 5 {
		t.Log(next())
	}
	reset()
	print("------ Reset -------\n")
	for range 4 {
		t.Log(next())
	}
}

// Number generator
func counter(start int) func() int {
	i := start
	return func() int {
		i++
		return i - 1
	}
}

func fibonacci() func() int {
	a, b := 0, 1
	return func() int {
		ret := a
		a, b = b, a+b
		return ret
	}
}

func loop(arr []int) (next func() int, reset func()) {
	i := 0
	return func() int {
			if i < len(arr) {
				i++
				return arr[i-1]
			}
			return -1
		}, func() {
			i = 0
		}
}

func TestTypeAssert(t *testing.T) {
	var i any = 1
	switch v := i.(type) {
	case int:
		t.Logf("int=%v", v)
	case string:
		t.Logf("string=%v", v)
	default:
		t.Logf("unknown")
	}

	wg := sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		print("")
	}()
	wg.Wait()

	ctx, cancel := context.WithCancel(context.TODO())
	defer cancel()
	ctx.Done()
}

func TestContextCancel(t *testing.T) {
	ctx, cancel := context.WithCancel(context.TODO())

	go func() {
		timer := time.After(3 * time.Second)
		<-timer
		cancel()
		print("goroutine timeout \n")

	}()
	<-ctx.Done()
	print("done")

	ch := make(chan int, 1)
	go func() {
		for {
			time.Sleep(1 * time.Second)
			_, ok := <-ch
			if !ok {
				print("channel closed")
				return
			}
			print("channel not closed \n")
		}
	}()
	go func() {
		for range 3 {
			ch <- 1
		}
	}()
	time.Sleep(5 * time.Second)
	close(ch)
}
