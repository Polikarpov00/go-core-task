package main

import (
	"fmt"
	"time"
)

func main() {
	wg := NewWaitGroup()

	for i := range 5 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Duration(i*100) * time.Millisecond)
			fmt.Println("горутина", i, "завершилась")
		}()
	}

	wg.Wait()
	fmt.Println("все горутины завершились")
}

type WaitGroup struct {
	sem   chan struct{}
	count int
	done  chan struct{}
}

func NewWaitGroup() *WaitGroup {
	done := make(chan struct{})
	close(done)
	return &WaitGroup{
		sem:  make(chan struct{}, 1),
		done: done,
	}
}

func (wg *WaitGroup) lock()   { wg.sem <- struct{}{} }
func (wg *WaitGroup) unlock() { <-wg.sem }

func (wg *WaitGroup) Add(n int) {
	wg.lock()
	defer wg.unlock()

	next := wg.count + n
	if next < 0 {
		panic("negative WaitGroup counter")
	}
	if wg.count == 0 && next > 0 {
		wg.done = make(chan struct{})
	}
	if wg.count > 0 && next == 0 {
		close(wg.done)
	}
	wg.count = next
}

func (wg *WaitGroup) Done() {
	wg.Add(-1)
}

func (wg *WaitGroup) Wait() {
	wg.lock()
	done := wg.done
	wg.unlock()
	<-done
}
