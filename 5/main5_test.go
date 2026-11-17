package main

import (
	"slices"
	"testing"
)

// TestIntersection проверяет поиск общих значений двух слайсов.
// Ожидается: для примера из задания — true и [64 3]; без пересечений — false и пустой (не nil) слайс;
// дубликаты в результат попадают один раз.
func TestIntersection(t *testing.T) {
	tests := []struct {
		name   string
		a, b   []int
		wantOk bool
		want   []int
	}{
		{"example", []int{65, 3, 58, 678, 64}, []int{64, 2, 3, 43}, true, []int{64, 3}},
		{"no intersection", []int{1, 2}, []int{3, 4}, false, []int{}},
		{"duplicates", []int{1, 1, 2}, []int{1, 1, 1}, true, []int{1}},
		{"empty", []int{}, []int{1}, false, []int{}},
		{"nil", nil, nil, false, []int{}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ok, got := intersection(tt.a, tt.b)
			if ok != tt.wantOk {
				t.Errorf("ok = %v, want %v", ok, tt.wantOk)
			}
			if got == nil {
				t.Error("result is nil, want empty slice")
			}
			if !slices.Equal(got, tt.want) {
				t.Errorf("intersection(%v, %v) = %v, want %v", tt.a, tt.b, got, tt.want)
			}
		})
	}
}
