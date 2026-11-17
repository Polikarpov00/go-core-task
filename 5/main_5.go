package main

import "fmt"

func main() {
	a := []int{65, 3, 58, 678, 64}
	b := []int{64, 2, 3, 43}

	fmt.Println(intersection(a, b))
}

func intersection(a, b []int) (bool, []int) {
	inA := make(map[int]struct{}, len(a))
	for _, v := range a {
		inA[v] = struct{}{}
	}

	res := []int{}
	added := make(map[int]struct{})
	for _, v := range b {
		if _, ok := inA[v]; !ok {
			continue
		}
		if _, dup := added[v]; dup {
			continue
		}
		added[v] = struct{}{}
		res = append(res, v)
	}
	return len(res) > 0, res
}
