package main

import "fmt"

func main() {
	slice1 := []string{"apple", "banana", "cherry", "date", "43", "lead", "gno1"}
	slice2 := []string{"banana", "date", "fig"}

	fmt.Println(difference(slice1, slice2))
}

func difference(a, b []string) []string {
	exclude := make(map[string]struct{}, len(b))
	for _, s := range b {
		exclude[s] = struct{}{}
	}

	res := make([]string, 0, len(a))
	for _, s := range a {
		if _, found := exclude[s]; !found {
			res = append(res, s)
		}
	}
	return res
}
