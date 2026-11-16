package main

import "testing"

// TestRandomGeneratorCount проверяет, что генератор выдаёт ровно n чисел.
// Ожидается: из канала прочитано 100 значений, после чего канал закрыт.
func TestRandomGeneratorCount(t *testing.T) {
	count := 0
	for range randomGenerator(100, 10) {
		count++
	}
	if count != 100 {
		t.Errorf("got %d values, want 100", count)
	}
}

// TestRandomGeneratorRange проверяет диапазон случайных чисел.
// Ожидается: все 1000 чисел в диапазоне от 0 до 9.
func TestRandomGeneratorRange(t *testing.T) {
	for v := range randomGenerator(1000, 10) {
		if v < 0 || v >= 10 {
			t.Fatalf("value %d out of [0, 10)", v)
		}
	}
}

// TestRandomGeneratorUnbuffered проверяет, что канал генератора небуферизированный.
// Ожидается: ёмкость канала равна 0.
func TestRandomGeneratorUnbuffered(t *testing.T) {
	ch := randomGenerator(1, 10)
	if cap(ch) != 0 {
		t.Errorf("cap = %d, want 0 (unbuffered)", cap(ch))
	}
	<-ch
}

// TestRandomGeneratorZero проверяет генератор при n = 0.
// Ожидается: канал сразу закрыт, значений нет.
func TestRandomGeneratorZero(t *testing.T) {
	_, ok := <-randomGenerator(0, 10)
	if ok {
		t.Error("channel should be closed when n = 0")
	}
}
