package main

import "fmt"

/*
ReverseStringSymbols создаёт руновое представление входной строки (руны -
целочисленний тип int32 в Го, используемый для кодировки Unicode символов) и
затем цикле с двух сторон переставляет элементы местами в этом слайсе рун.
в итоге мы получаем все символы переставленными задом наперёд и возвращаем
строковое представление полученного слайса рун
*/
func ReverseStringSymbols(s string) string {
	runes := []rune(s)

	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func main() {
	s := "hello world"
	fmt.Println(ReverseStringSymbols(s))
	s = "леопард ядра поел"
	fmt.Println(ReverseStringSymbols(s))
}
