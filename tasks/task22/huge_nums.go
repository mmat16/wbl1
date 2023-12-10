package main

import (
	"flag"
	"fmt"
	"log"
	"math/big"
)

// переменные для хранения значений аргументов командной строки
var (
	sum  bool
	sub  bool
	mul  bool
	div  bool
	num1 string
	num2 string
)

/*
Sum складывает два числа типа *big.Int (специальный тип для целых чисел
произвольного длины) и возвращает результат вычислений.
*/
func Sum(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Add(a, b)
	return res
}

/*
Substract вычитает b из a - чисел типа *big.Int и возвращает результат
вычислений
*/
func Substract(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Sub(a, b)
	return res
}

/*
Multiply возвращает произведение двух чисел типа *big.Int
*/
func Multiply(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Mul(a, b)
	return res
}

/*
Divide делит число a типа *big.Int на число b типа *big.Int и возвращает
результат вычислений
*/
func Divide(a, b *big.Int) *big.Int {
	res := new(big.Int)
	res.Div(a, b)
	return res
}

/*
proccess создаёт два числа типа *big.Int, выставляет в них значения, прочитанные
из аргументов командной строки и проверяет таким образом валидность прочитанных
значений - в случае если присвоить данные не удалось (данные оказались не
валидны), программа аварийно завершается так как произвести вычисления не
возможно. далее выполняется операция, выбранная с помощью оператора switch-case
определённая так же с помощью аргумента командной строки и результат вычислений
выводится в stdout с сообщением о типе выполненной операции.
*/
func proccess() {
	a, b := new(big.Int), new(big.Int)

	_, ok := a.SetString(num1, 10)
	if !ok {
		log.Fatal("invalid value for num1: ", num1)
	}

	_, ok = b.SetString(num2, 10)
	if !ok {
		log.Fatal("invalid value for num2: ", num2)
	}

	switch {
	case sum:
		fmt.Println("Sum:", Sum(a, b).String())
	case sub:
		fmt.Println("Difference:", Substract(a, b).String())
	case mul:
		fmt.Println("Product:", Multiply(a, b).String())
	case div:
		fmt.Println("Quotient:", Divide(a, b).String())
	}
}

/*
функция инициализации, выполняется перед функцией main. назначает переменные
для чтения аргументов командной строки и парсит их с помощью стандартной
библиотеки flag
*/
func init() {
	flag.BoolVar(&sum, "sum", false, "proccess addition of two numbers")
	flag.BoolVar(&sub, "sub", false, "proccess substraction of two numbers")
	flag.BoolVar(&mul, "mul", false, "proccess multiplication of two numbers")
	flag.BoolVar(&div, "div", false, "procces division of two numbers")
	flag.StringVar(&num1, "first", "", "value of the first to number to operate")
	flag.StringVar(&num2, "second", "", "value of the second number to operate")
	flag.Parse()
}

/*
main функция вызывает функцию proccess, выполняющую операцию над числами.
*/
func main() {
	proccess()
}
