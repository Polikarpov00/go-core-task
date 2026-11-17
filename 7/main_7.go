package main

import (
	"fmt"
	"sync"
)

func main() {
	a := source(1, 2, 3)
	b := source(10, 20, 30)
	c := source(100, 200, 300)

	for v := range merge(a, b, c) {
		fmt.Println(v)
	}
}

func merge(chans ...<-chan int) <-chan int {
	out := make(chan int)

	var wg sync.WaitGroup
	for _, ch := range chans {
		wg.Add(1)
		go func(ch <-chan int) {
			defer wg.Done()
			for v := range ch {
				out <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func source(values ...int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for _, v := range values {
			ch <- v
		}
	}()
	return ch
}
