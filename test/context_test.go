package test

import (
	"context"
	"fmt"
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
	context1()
}

func context1() {
	ctx, _ := context.WithTimeout(context.TODO(), 5*time.Second)
	for i := 0; i < 2; i++ {
		go func() {
			fmt.Println("go routine start...")
			select {
			case <-ctx.Done():
				fmt.Println("Done")
			default:
				fmt.Println("nothing happen")
			}
		}()
	}

	time.Sleep(6 * time.Second)
}
