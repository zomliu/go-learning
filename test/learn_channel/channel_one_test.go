package learnchannel

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

// implement a concurrency counter by using atomic package
func TestCounterOne(t *testing.T) {
	counter := atomic.Int32{}

	for range 15 {
		go func() {
			counter.Add(1)
		}()
	}

	timer := time.After(time.Second * 2)
	<-timer
	fmt.Println(counter.Load())
	fmt.Println("done")
}

// implement a concurrency counter by using sync.Mutex
func TestCounterTwo(t *testing.T) {
	var counter int32
	lock := sync.Mutex{}

	for range 15 {
		go func() {
			lock.Lock()
			defer lock.Unlock()
			counter++
		}()
	}
	timer := time.After(time.Second * 2)
	<-timer
	fmt.Println("result: " + fmt.Sprint(counter))
	fmt.Println("done")
}

// reading data from a channel, return immediately if timeout
func TestReadChannelUnderTimeout(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second * 10)
	defer cancel()
	ch := make(chan bool)

	go func() {
		
		timer := time.After(time.Second * 2)

		// 无条件的 for 循环一定要有一个 return
		for {
			select {
			case _, ok := <-ch:  // better read the ready signal in case the channel is closed
				if !ok {
					fmt.Println("channel closed")
				}
				fmt.Println("receive data")
			case <-timer:
				fmt.Println("timeout")
				return
			case <-ctx.Done():
				fmt.Println("consumer context done")
				return
			}
		}
	}()

	go func() {
		for i := range 4 {
			time.Sleep(time.Second * time.Duration(i))
			select {
			case ch <- true:
				fmt.Println("send data")
			case <-ctx.Done():
				fmt.Println("producer context done")
				return
			}
		}
		// only writer can close a channel, otherwise program will panic if writer try to write data into a closed channel
		close(ch)
	}()

	<- ctx.Done()
	fmt.Println("all done")
}
