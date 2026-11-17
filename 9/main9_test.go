package main

import (
	"slices"
	"testing"
)

func runPipeline(values []uint8) []float64 {
	in := make(chan uint8)
	out := make(chan float64)

	go func() {
		defer close(in)
		for _, v := range values {
			in <- v
		}
	}()
	go cubePipeline(in, out)

	var res []float64
	for v := range out {
		res = append(res, v)
	}
	return res
}

// TestCubePipeline проверяет конвейер: uint8 -> float64 -> куб.
// Ожидается: [1 2 3] -> [1 8 27], 0 -> 0, 255 -> 16581375 (без переполнения), пустой вход — пустой выход.
func TestCubePipeline(t *testing.T) {
	tests := []struct {
		name string
		in   []uint8
		want []float64
	}{
		{"basic", []uint8{1, 2, 3}, []float64{1, 8, 27}},
		{"zero", []uint8{0}, []float64{0}},
		{"max uint8", []uint8{255}, []float64{16581375}},
		{"empty", []uint8{}, nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runPipeline(tt.in); !slices.Equal(got, tt.want) {
				t.Errorf("pipeline(%v) = %v, want %v", tt.in, got, tt.want)
			}
		})
	}
}
