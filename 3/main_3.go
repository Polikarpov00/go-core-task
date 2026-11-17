package main

import "fmt"

func main() {
	m := NewStringIntMap()
	m.Add("one", 1)
	m.Add("two", 2)
	m.Add("three", 3)

	fmt.Println("Exists(\"two\"):", m.Exists("two"))

	v, ok := m.Get("three")
	fmt.Println("Get(\"three\"):", v, ok)

	m.Remove("two")
	fmt.Println("после Remove(\"two\"), Exists:", m.Exists("two"))

	c := m.Copy()
	c["new"] = 100
	fmt.Println("копия:", c)
	fmt.Println("оригинал не изменился, Exists(\"new\"):", m.Exists("new"))
}

type StringIntMap struct {
	data map[string]int
}

func NewStringIntMap() *StringIntMap {
	return &StringIntMap{data: make(map[string]int)}
}

func (m *StringIntMap) Add(key string, value int) {
	m.data[key] = value
}

func (m *StringIntMap) Remove(key string) {
	delete(m.data, key)
}

func (m *StringIntMap) Copy() map[string]int {
	res := make(map[string]int, len(m.data))
	for k, v := range m.data {
		res[k] = v
	}
	return res
}

func (m *StringIntMap) Exists(key string) bool {
	_, ok := m.data[key]
	return ok
}

func (m *StringIntMap) Get(key string) (int, bool) {
	v, ok := m.data[key]
	return v, ok
}
