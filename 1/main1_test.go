package main

import (
	"crypto/sha256"
	"encoding/hex"
	"slices"
	"testing"
)

// TestGetType проверяет, что getType правильно определяет тип переменной.
// Ожидается: int, float64, string, bool, complex64; для 052 и 0x2A тоже int.
func TestGetType(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{"int", 42, "int"},
		{"octal int", 052, "int"},
		{"hex int", 0x2A, "int"},
		{"float64", 3.14, "float64"},
		{"string", "Golang", "string"},
		{"bool", true, "bool"},
		{"complex64", complex64(1 + 2i), "complex64"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getType(tt.in); got != tt.want {
				t.Errorf("getType(%v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// TestToString проверяет перевод значений разных типов в строку.
// Ожидается: 42 -> "42", 3.14 -> "3.14", true -> "true", 1+2i -> "(1+2i)".
func TestToString(t *testing.T) {
	tests := []struct {
		name string
		in   any
		want string
	}{
		{"int", 42, "42"},
		{"octal int", 052, "42"},
		{"float64", 3.14, "3.14"},
		{"string", "Golang", "Golang"},
		{"bool", true, "true"},
		{"complex64", complex64(1 + 2i), "(1+2i)"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toString(tt.in); got != tt.want {
				t.Errorf("toString(%v) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// TestConcat проверяет склейку всех переменных в одну строку.
// Ожидается: "4242423.14Golangtrue(1+2i)".
func TestConcat(t *testing.T) {
	got := concat(42, 052, 0x2A, 3.14, "Golang", true, complex64(1+2i))
	want := "4242423.14Golangtrue(1+2i)"
	if got != want {
		t.Errorf("concat() = %q, want %q", got, want)
	}
}

// TestToRunes проверяет перевод строки в срез рун.
// Ожидается: "привет" даёт 6 рун (а не 12 байт).
func TestToRunes(t *testing.T) {
	got := toRunes("привет")
	if len(got) != 6 {
		t.Errorf("len = %d, want 6", len(got))
	}
}

// TestInsertSalt проверяет, что соль вставляется в середину среза рун.
// Ожидается: "abcd" -> "abgo-2024cd", пустая строка -> "go-2024", кириллица не ломается.
func TestInsertSalt(t *testing.T) {
	tests := []struct {
		name string
		in   string
		want string
	}{
		{"even", "abcd", "abgo-2024cd"},
		{"odd", "abc", "ago-2024bc"},
		{"empty", "", "go-2024"},
		{"cyrillic", "мама", "маgo-2024ма"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := string(insertSalt([]rune(tt.in), salt)); got != tt.want {
				t.Errorf("insertSalt(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

// TestInsertSaltDoesNotModifyInput проверяет, что insertSalt не меняет исходный срез.
// Ожидается: после вызова исходный срез остаётся "abcd".
func TestInsertSaltDoesNotModifyInput(t *testing.T) {
	in := []rune("abcd")
	orig := slices.Clone(in)
	insertSalt(in, salt)
	if !slices.Equal(in, orig) {
		t.Errorf("input modified: %q, want %q", string(in), string(orig))
	}
}

// TestHashRunes проверяет SHA256-хэш строки с солью.
// Ожидается: хэш совпадает с sha256("abgo-2024cd") и имеет длину 64 символа.
func TestHashRunes(t *testing.T) {
	sum := sha256.Sum256([]byte("abgo-2024cd"))
	want := hex.EncodeToString(sum[:])

	got := hashRunes(insertSalt([]rune("abcd"), salt))
	if got != want {
		t.Errorf("hashRunes() = %s, want %s", got, want)
	}
	if len(got) != 64 {
		t.Errorf("hash length = %d, want 64", len(got))
	}
}
