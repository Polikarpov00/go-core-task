package main

import (
	"sync/atomic"
	"testing"
	"time"
)

// TestWaitWithoutAdd проверяет Wait без вызова Add.
// Ожидается: Wait сразу возвращается и не блокирует.
func TestWaitWithoutAdd(t *testing.T) {
	wg := NewWaitGroup()
	wg.Wait()
}

// TestWaitGroupWaitsForAll проверяет, что Wait дожидается всех горутин.
// Ожидается: после Wait все 100 горутин отработали, счётчик равен 100.
func TestWaitGroupWaitsForAll(t *testing.T) {
	wg := NewWaitGroup()
	var counter atomic.Int64

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			time.Sleep(time.Millisecond)
			counter.Add(1)
		}()
	}

	wg.Wait()
	if counter.Load() != 100 {
		t.Errorf("counter = %d, want 100", counter.Load())
	}
}

// TestWaitGroupReuse проверяет повторное использование WaitGroup.
// Ожидается: три цикла Add/Done/Wait подряд проходят без зависаний.
func TestWaitGroupReuse(t *testing.T) {
	wg := NewWaitGroup()
	for i := 0; i < 3; i++ {
		wg.Add(2)
		go wg.Done()
		go wg.Done()
		wg.Wait()
	}
}

// TestNegativeCounterPanics проверяет вызов Done без Add.
// Ожидается: паника из-за отрицательного счётчика.
func TestNegativeCounterPanics(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Error("expected panic on negative counter")
		}
	}()
	NewWaitGroup().Done()
}
