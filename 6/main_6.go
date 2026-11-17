package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	for v := range randomGenerator(10, 100) {
		fmt.Println(v)
	}
}

func randomGenerator(n, max int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := 0; i < n; i++ {
			ch <- rand.IntN(max)
		}
	}()
	return ch
}
