package main

import (
	"unicode"
)

// AreSymbolsUnique проверяет, что все символы в строке уникальные, с помощью
// map (ключ - руна, значение - true/false). В цикле перебираются все руны
// строки и записываются в map со значением true. После цикла проверяется что
// длина map равна длине строки (если есть повторяющиеся символы, то длина map
// будет меньше длины строки).
func AreSymbolsUnique(s string) bool {
	m := make(map[rune]bool)
	for _, r := range s {
		r = unicode.ToLower(r)
		m[r] = true
	}
	return len(m) == len(s)
}

func main() {
	println(AreSymbolsUnique("abc"))
	println(AreSymbolsUnique("aBc"))
	println(AreSymbolsUnique("aBca"))
	println(AreSymbolsUnique("Abca"))
}
