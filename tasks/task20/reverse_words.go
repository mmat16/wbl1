package main

import (
	"fmt"
	"strings"
)

/*
ReverseWords использует для записи результирующей строки strings.Builder как
более оптимальное решение для конкатенации строк, по сравнению с конкатенацией
"в лоб" при помощи арифметических операторов. так как строки в Го являются
read-only слайсами байт, каждое изменение строки приводит к перезаписи значений
в новую переменную, а следовательно и выделение новой памяти, что сказывается
на производительности программы. Builder решает эту проблему, предоставляя метод
Grow для предварительного выделения памяти, что в случае данного задания, когда
длина результирующей строки заранее известна, позволяет снизить количество
операций выделения памяти под результирующую строку до одной.
остальной алгоритм довольно прост - разбивка строки на слова по пробелам, с
использованием strings.Fields, затем итерация от конца до начала полученного
слайса строк и запись их в память Builder'a. по завершении цикла из функции
возвращается результат вызова метода builder.String, собирающий записанную в
билдер информацию в строковое представление.
*/
func ReverseWords(s string) string {
	words := strings.Fields(s)
	var builder strings.Builder

	builder.Grow(len(s))

	for i := len(words) - 1; i >= 0; i-- {
		builder.WriteString(words[i])
		builder.WriteRune(' ')
	}

	return builder.String()
}

/*
ReverseWordsConcat приведена для сравнения производительности со
strings.Builder. этот вариант функции использует самый очевидный способ
конкатенации строк, который однако является крайне неэффективным. ниже привожу
пример бенчмарка двух функций

❯ go test -bench="." -benchmem
goos: darwin
goarch: amd64
pkg: wb01/tasks/task20
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkReverseWords-8         	 1070458	      1108 ns/op	     496 B/op	       3 allocs/op
BenchmarkReverseWordsConcat-8   	  761736	      1492 ns/op	     928 B/op	      11 allocs/op
PASS
ok  	wb01/tasks/task20	3.338s

очевидно что первый вариант выигрывает и по скорости выполнения, и более ощутимо
обходит второй вариант по использованию памяти и количеству аллокаций.
*/
func ReverseWordsConcat(s string) string {
	words := strings.Fields(s)
	var res string

	for i := len(words) - 1; i >= 0; i-- {
		res += words[i] + " "
	}
	return res
}

func main() {
	s := "моя гордыня это дыня гор а не какой-нибудь равнинный помидор"
	fmt.Println(ReverseWords(s))
	fmt.Println(ReverseWordsConcat(s))
}
