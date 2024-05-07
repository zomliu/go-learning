package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

func main() {
	// generate a random int between 0 and 100
	randInt, err := rand.Int(rand.Reader, big.NewInt(10))
	if err != nil {
		panic(err)
	}
	switch randInt.Int64() {
	case 1, 2, 3, 4, 5:
		fmt.Println("less than 5")
	case 6, 7, 8, 9, 10:
		fmt.Println("greater than 5")
	default:
		fmt.Println("default")
	}
}
