package main

import (
	"slices"
	"testing"
)

// TestDifference проверяет поиск элементов, которые есть в первом слайсе, но нет во втором.
// Ожидается: для примера из задания — [apple cherry 43 lead gno1]; для пустых и nil-слайсов — пустой слайс.
func TestDifference(t *testing.T) {
	tests := []struct {
		name string
		a, b []string
		want []string
	}{
		{
			"example",
			[]string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"},
			[]string{"banana", "date", "fig"},
			[]string{"apple", "cherry", "43", "lead", "gno1"},
		},
		{"empty second", []string{"a", "b"}, []string{}, []string{"a", "b"}},
		{"empty first", []string{}, []string{"a"}, []string{}},
		{"all excluded", []string{"a", "b"}, []string{"b", "a"}, []string{}},
		{"nil inputs", nil, nil, []string{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := difference(tt.a, tt.b); !slices.Equal(got, tt.want) {
				t.Errorf("difference(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
