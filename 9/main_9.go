package main

import "fmt"

func main() {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		for i := uint8(1); i <= 10; i++ {
			in <- i
		}
	}()

	go cubePipeline(in, out)

	for v := range out {
		fmt.Println(v)
	}
}

func cubePipeline(in <-chan uint8, out chan<- float64) {
	defer close(out)
	for v := range in {
		f := float64(v)
		out <- f * f * f
	}
}
