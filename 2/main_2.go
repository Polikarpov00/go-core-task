package main

import (
	"errors"
	"fmt"
	"math/rand/v2"
)

var ErrIndexOutOfRange = errors.New("index out of range")

func main() {
	originalSlice := randomSlice(10, 100)
	fmt.Println("originalSlice:", originalSlice)

	fmt.Println("sliceExample:", sliceExample(originalSlice))
	fmt.Println("addElements:", addElements(originalSlice, 777))

	copied := copySlice(originalSlice)
	originalSlice[0] = -1
	fmt.Println("после изменения оригинала: ", originalSlice)
	fmt.Println("копия не изменилась: ", copied)

	removed, err := removeElement(originalSlice, 3)
	if err != nil {
		fmt.Println("removeElement:", err)
		return
	}
	fmt.Println("removeElement:", removed)
	fmt.Println("originalSlice:", originalSlice)
}

func randomSlice(n, max int) []int {
	s := make([]int, n)
	for i := range s {
		s[i] = rand.IntN(max)
	}
	return s
}

func sliceExample(s []int) []int {
	res := make([]int, 0, len(s))
	for _, v := range s {
		if v%2 == 0 {
			res = append(res, v)
		}
	}
	return res
}

func addElements(s []int, v int) []int {
	res := make([]int, len(s), len(s)+1)
	copy(res, s)
	return append(res, v)
}

func copySlice(s []int) []int {
	res := make([]int, len(s))
	copy(res, s)
	return res
}

func removeElement(s []int, idx int) ([]int, error) {
	if idx < 0 || idx >= len(s) {
		return nil, ErrIndexOutOfRange
	}
	res := make([]int, 0, len(s)-1)
	res = append(res, s[:idx]...)
	res = append(res, s[idx+1:]...)
	return res, nil
}
