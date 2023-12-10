package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

/*
IntersectionON находит пересечение двух множеств, представленных двумя слайсами-
аргументами функции за время приблизительно равное O(n). Внутри функции
создаётся результирующий слайс для пересечения входных множеств и
вспомогательная мапа для хранения всех значений первого исходного слайса.
После того как все элементы первого слайса помещенны в мапу, в цикле через
второй слайс каждый его элемент проверяется на наличие в мапе и если он в ней
есть, то такой элемент добавляется в результирующий слайс.
Возвращается в качестве результата этот слайс, пропущенный через функцию,
убирающую дупликаты в слайсе.
*/
func IntersectionON[T constraints.Ordered](arr1 []T, arr2 []T) []T {
	isInFirstArr := make(map[T]bool, len(arr1))
	for _, num := range arr1 {
		isInFirstArr[num] = true
	}

	var intersection []T
	for _, num := range arr2 {
		if isInFirstArr[num] {
			intersection = append(intersection, num)
		}
	}
	return RemoveDups(intersection)
}

/*
IntersectionONSquare так же находит пересечение двух множеств, но за время,
приблизительно равное O(n^2) так как какждый элемент первого массива
сравнивается с каждым элементом второго массива. Каждый раз, когда элементы
равны - один из них добавляется в результирующий массив, из которого после всех
манипуляций удаляются дупликаты.
*/
func IntersectionONSquare[T constraints.Ordered](arr1 []T, arr2 []T) []T {
	var intersection []T
	for _, num := range arr1 {
		for _, num2 := range arr2 {
			if num == num2 {
				intersection = append(intersection, num)
			}
		}
	}
	return RemoveDups(intersection)
}

/*
RemoveDups удаляет повторяющиеся элементы в массиве. Для этого используется
мапа, в которой в цикле поочерёдно проверяется наличие текущего элемента. если
элемент ещё не присутсвует в ней - он добавляется в мапу и результирующий слайс,
который возвращается из функции по завершении функции.
*/
func RemoveDups[T constraints.Ordered](arr []T) []T {
	var res []T
	noDups := make(map[T]bool, len(arr))
	for _, num := range arr {
		if !noDups[num] {
			noDups[num] = true
			res = append(res, num)
		}
	}
	return res
}

func main() {
	first := []int{1, 2, 3, 2, 3}
	second := []int{5, 4, 1, 3, 1, 3, 1}
	ON := IntersectionON(first, second)
	ONSquare := IntersectionONSquare(first, second)
	fmt.Println(ON)
	fmt.Println(ONSquare)
	firstStr := []string{"h", "e", "l", "l", "o"}
	secondStr := []string{"t", "h", "e", "r", "e"}
	fmt.Println(IntersectionON(firstStr, secondStr))
	fmt.Println(IntersectionONSquare(firstStr, secondStr))
}
