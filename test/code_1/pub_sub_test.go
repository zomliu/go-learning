package code_1

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPubSub(t *testing.T) {
	pub := NewPublisher()

	sub1 := pub.Subscribe()
	sub2 := pub.Subscribe()

	go func() {
		for msg := range sub1 {
			fmt.Printf("sub1: %s\n", msg)
		}
	}()

	go func() {
		for msg := range sub2 {
			fmt.Printf("sub2: %s\n", msg)
		}
	}()

	for idx := range 5 {
		pub.Publish(fmt.Sprintf("msg %d", idx + 1))
	}

	timer := time.After(time.Second * 3)
	tm := <- timer

	fmt.Printf("time: %s \n", tm)

	pub.Unsubscribe(sub1)
	pub.Unsubscribe(sub2)

	fmt.Println("done")
}

type Publisher struct {
	subscribers map[chan string]struct{}
	L           sync.Mutex
}

func NewPublisher() *Publisher {
	return &Publisher{
		subscribers: make(map[chan string]struct{}),
	}
}

func (p *Publisher) Subscribe() chan string {
	p.L.Lock()
	defer p.L.Unlock()
	ch := make(chan string, 1)
	p.subscribers[ch] = struct{}{}
	return ch
}

func (p *Publisher) Unsubscribe(ch chan string) {
	p.L.Lock()
	defer p.L.Unlock()
	delete(p.subscribers, ch)
	close(ch)
}

func (p *Publisher) Publish(msg string) {
	p.L.Lock()
	defer p.L.Unlock()
	for ch := range p.subscribers {
		ch <- msg
	}
}
