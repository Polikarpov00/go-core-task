package main

import (
	"slices"
	"testing"
)

func collect(ch <-chan int) []int {
	var res []int
	for v := range ch {
		res = append(res, v)
	}
	return res
}

// TestMerge проверяет слияние нескольких каналов в один.
// Ожидается: на выходе все значения из всех каналов — [1 2 3 4 5 6] (после сортировки).
func TestMerge(t *testing.T) {
	got := collect(merge(source(1, 2, 3), source(4, 5), source(6)))
	slices.Sort(got)

	want := []int{1, 2, 3, 4, 5, 6}
	if !slices.Equal(got, want) {
		t.Errorf("merge = %v, want %v", got, want)
	}
}

// TestMergeNoChannels проверяет merge без входных каналов.
// Ожидается: выходной канал сразу закрыт, значений нет.
func TestMergeNoChannels(t *testing.T) {
	if got := collect(merge()); len(got) != 0 {
		t.Errorf("merge() = %v, want empty", got)
	}
}

// TestMergeEmptyChannels проверяет merge с пустыми каналами.
// Ожидается: выходной канал закрыт, значений нет.
func TestMergeEmptyChannels(t *testing.T) {
	if got := collect(merge(source(), source())); len(got) != 0 {
		t.Errorf("merge = %v, want empty", got)
	}
}

// TestMergeSingleChannel проверяет merge с одним каналом.
// Ожидается: значения [1 2 3] в исходном порядке.
func TestMergeSingleChannel(t *testing.T) {
	got := collect(merge(source(1, 2, 3)))
	if !slices.Equal(got, []int{1, 2, 3}) {
		t.Errorf("merge = %v, want [1 2 3]", got)
	}
}
