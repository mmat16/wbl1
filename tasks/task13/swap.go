package main

import "fmt"

/*
BitwiseSwap использует битовый оператор XOR (исключающее или) для
последовательной перестановке значений переменных, без использования
вспомогательной переменной. XOR, как я описывал в одной из задач ранее,
возвращает битовую единицу в результате если только у одного из операндов на
данной позиции есть единица.
в случае перестановки получается если first = 10, а second = 5 - их битовое
представление выглядит следующим образом: first = 1010, second = 0101;
first = first ^ second (1010 ^ 0101 = 1111)
second = first ^ second (1111 ^ 0101 = 1010)
first = first ^ second (1111 ^ 1010 = 0101)
и в итоге мы получае что first = 0101 = 5, а second = 1010 = 10 - значения
поменялись местами, без использование дополнительной переменной!
*/
func BitwiseSwap(first, second *int) {
	*first = *first ^ *second
	*second = *first ^ *second
	*first = *first ^ *second
}

/*
SyntacticSwap использует синтаксическую "лазейку" Го и многих других языков,
позволяющую определять и переопределять несколько переменных одновременно
в одной строке, что позволяет менять местами значения в переменных без
аллокации дополнительной памяти
*/
func SyntacticSwap(first, second *int) {
	*first, *second = *second, *first
}

/*
SyntacticSwap использует арифметику для перестановки занчений. посмотрим на
происходящее поближе, для примера возьмём знакомые значение - first = 10,
second = 5;
на первом шаге прибавляем second к first и получаем first = 15
на втором шаге присваем в second разницу first и second, получаем second = 10
и на третьем шаге, отнимая от first новый second - получаем результат first = 5
и снова значение поменялись, не используя дополнительных переменных
*/
func ArithmeticSwap(first, second *int) {
	*first += *second
	*second = *first - *second
	*first -= *second
}

func main() {
	a, b := 10, 5

	fmt.Printf("\nbefore swap\na=%d, b=%d\n", a, b)
	BitwiseSwap(&a, &b)
	fmt.Printf("after swap\na=%d, b=%d\n", a, b)

	fmt.Printf("\nbefore swap\na=%d, b=%d\n", a, b)
	SyntacticSwap(&a, &b)
	fmt.Printf("after swap\na=%d, b=%d\n", a, b)

	fmt.Printf("\nbefore swap\na=%d, b=%d\n", a, b)
	ArithmeticSwap(&a, &b)
	fmt.Printf("after swap\na=%d, b=%d\n", a, b)
}
