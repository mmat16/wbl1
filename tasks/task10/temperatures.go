package main

import "fmt"

/*
SplitByGroups принимает слайс чисел формата float64 и возвращает мапу с
целочисленным ключом и слайсом чисел float64 как значение ключа. в цикле
перебираются все значения входящего слайса. на каждой итерации текущее значение
конвертируется в целочисленное и округляется до десятка "вниз", далее полученное
число становится ключом мапы и в массив по нему дополняется оригинальный элемент
слайса.
*/
func SplitByGroups(nums []float64) map[int][]float64 {
	res := make(map[int][]float64)
	for _, num := range nums {
		raw := int(num) / 10 * 10
		res[raw] = append(res[raw], num)
	}
	return res
}

func main() {
	res := SplitByGroups([]float64{1, 3.5, 21.1, -11, 13, 12.1, 30.7, -20.1, -45.14})
	fmt.Println(res)
}
