package main

import (
	"errors"
	"slices"
	"testing"
)

// TestRandomSlice проверяет генерацию случайного слайса.
// Ожидается: длина 10, все числа в диапазоне от 0 до 99.
func TestRandomSlice(t *testing.T) {
	s := randomSlice(10, 100)
	if len(s) != 10 {
		t.Fatalf("len = %d, want 10", len(s))
	}
	for _, v := range s {
		if v < 0 || v >= 100 {
			t.Errorf("value %d out of [0, 100)", v)
		}
	}
}

// TestSliceExample проверяет отбор чётных чисел из слайса.
// Ожидается: только чётные числа (включая 0 и отрицательные), для пустого — пустой слайс.
func TestSliceExample(t *testing.T) {
	tests := []struct {
		name string
		in   []int
		want []int
	}{
		{"mixed", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 4, 6, 8, 10}},
		{"no evens", []int{1, 3, 5}, []int{}},
		{"negative", []int{-4, -3, 0}, []int{-4, 0}},
		{"empty", []int{}, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sliceExample(tt.in); !slices.Equal(got, tt.want) {
				t.Errorf("sliceExample(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}

// TestAddElements проверяет добавление числа в конец слайса.
// Ожидается: [1 2 3] + 4 -> [1 2 3 4], исходный слайс при этом не меняется.
func TestAddElements(t *testing.T) {
	in := make([]int, 3, 10)
	copy(in, []int{1, 2, 3})

	got := addElements(in, 4)
	if !slices.Equal(got, []int{1, 2, 3, 4}) {
		t.Errorf("addElements = %v, want [1 2 3 4]", got)
	}

	got[0] = 100
	if in[0] != 1 {
		t.Errorf("addElements result shares memory with input")
	}
}

// TestCopySlice проверяет копирование слайса.
// Ожидается: копия равна оригиналу, изменение оригинала не влияет на копию.
func TestCopySlice(t *testing.T) {
	in := []int{1, 2, 3}
	got := copySlice(in)
	if !slices.Equal(got, in) {
		t.Fatalf("copySlice = %v, want %v", got, in)
	}

	in[0] = 100
	if got[0] != 1 {
		t.Errorf("copy changed after modifying original: %v", got)
	}
}

// TestRemoveElement проверяет удаление элемента по индексу.
// Ожидается: элемент удалён, исходный слайс не изменён; при неверном индексе — ошибка ErrIndexOutOfRange.
func TestRemoveElement(t *testing.T) {
	tests := []struct {
		name    string
		in      []int
		idx     int
		want    []int
		wantErr error
	}{
		{"middle", []int{1, 2, 3, 4}, 1, []int{1, 3, 4}, nil},
		{"first", []int{1, 2, 3}, 0, []int{2, 3}, nil},
		{"last", []int{1, 2, 3}, 2, []int{1, 2}, nil},
		{"single", []int{1}, 0, []int{}, nil},
		{"negative idx", []int{1, 2}, -1, nil, ErrIndexOutOfRange},
		{"idx too big", []int{1, 2}, 2, nil, ErrIndexOutOfRange},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			orig := slices.Clone(tt.in)
			got, err := removeElement(tt.in, tt.idx)
			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("err = %v, want %v", err, tt.wantErr)
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("removeElement(%v, %d) = %v, want %v", tt.in, tt.idx, got, tt.want)
			}
			if !slices.Equal(tt.in, orig) {
				t.Errorf("input modified: %v, want %v", tt.in, orig)
			}
		})
	}
}
