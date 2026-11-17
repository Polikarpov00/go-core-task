package main
import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"reflect"
	"strconv"
	"strings"
)
const salt = "go-2024"

func main() {
	var numDecimal int = 42
	var numOctal int = 052
	var numHexadecimal int = 0x2A
	var pi float64 = 3.14
	var name string = "Golang"
	var isActive bool = true
	var complexNum complex64 = 1 + 2i
	
	values := []any{numDecimal, numOctal, numHexadecimal, pi, name, isActive, complexNum}
	
	fmt.Println("Типы значений:")
	for _, v := range values {
		fmt.Printf("Значение: %v; Тип: %s\n", v, getType(v))
	}
	
	s := concat(values...)
	fmt.Println("Строка:", s)
	
	r := toRunes(s)
	fmt.Println("Руны:", r)
	salted := insertSalt(r, salt)
	fmt.Println("Соль:", string(salted))
	fmt.Println("SHA256:", hashRunes(salted))
}


func getType(v any) string {
	return reflect.TypeOf(v).String()
}

func toString(v any) string {
	switch x := v.(type) {
	case int:
		return strconv.Itoa(x)
	case float64:
		return strconv.FormatFloat(x, 'f', -1, 64)
	case string:
		return x
	case bool:
		return strconv.FormatBool(x)
	case complex64:
		return strconv.FormatComplex(complex128(x), 'f', -1, 64)
	default:
		return fmt.Sprint(x)
	}
}

func concat(values ...any) string {
	var sb strings.Builder
	for _, v := range values {
		sb.WriteString(toString(v))
	}
	return sb.String()
}

func toRunes(s string) []rune {
	return []rune(s)
}

func insertSalt(r []rune, salt string) []rune {
	mid := len(r) / 2
	saltRunes := []rune(salt)
	res := make([]rune, 0, len(r)+len(saltRunes))
	res = append(res, r[:mid]...)
	res = append(res, saltRunes...)
	res = append(res, r[mid:]...)
	return res
}

func hashRunes(r []rune) string {
	sum := sha256.Sum256([]byte(string(r)))
	return hex.EncodeToString(sum[:])
}