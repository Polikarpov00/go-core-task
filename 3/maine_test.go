package main

import (
	"maps"
	"testing"
)

// TestAddAndGet проверяет добавление и получение значения по ключу.
// Ожидается: Get возвращает добавленное значение и true; повторный Add перезаписывает значение.
func TestAddAndGet(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)

	v, ok := m.Get("a")
	if !ok || v != 1 {
		t.Errorf("Get(a) = %d, %v; want 1, true", v, ok)
	}

	m.Add("a", 2)
	if v, _ := m.Get("a"); v != 2 {
		t.Errorf("after overwrite Get(a) = %d, want 2", v)
	}
}

// TestGetMissing проверяет получение несуществующего ключа.
// Ожидается: 0 и false.
func TestGetMissing(t *testing.T) {
	m := NewStringIntMap()
	v, ok := m.Get("missing")
	if ok || v != 0 {
		t.Errorf("Get(missing) = %d, %v; want 0, false", v, ok)
	}
}

// TestRemove проверяет удаление ключа.
// Ожидается: после Remove ключа нет; удаление несуществующего ключа не вызывает панику.
func TestRemove(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Remove("a")
	if m.Exists("a") {
		t.Error("key a still exists after Remove")
	}

	m.Remove("missing")
}

// TestExists проверяет наличие ключа в карте.
// Ожидается: true для существующего ключа (даже со значением 0), false для отсутствующего.
func TestExists(t *testing.T) {
	m := NewStringIntMap()
	m.Add("zero", 0)

	if !m.Exists("zero") {
		t.Error("Exists(zero) = false, want true")
	}
	if m.Exists("missing") {
		t.Error("Exists(missing) = true, want false")
	}
}

// TestCopy проверяет копирование карты.
// Ожидается: копия содержит те же элементы, изменения копии не влияют на оригинал.
func TestCopy(t *testing.T) {
	m := NewStringIntMap()
	m.Add("a", 1)
	m.Add("b", 2)

	c := m.Copy()
	want := map[string]int{"a": 1, "b": 2}
	if !maps.Equal(c, want) {
		t.Fatalf("Copy() = %v, want %v", c, want)
	}

	c["a"] = 100
	c["new"] = 3
	if v, _ := m.Get("a"); v != 1 {
		t.Errorf("original changed after modifying copy: a = %d", v)
	}
	if m.Exists("new") {
		t.Error("key added to copy appeared in original")
	}
}
