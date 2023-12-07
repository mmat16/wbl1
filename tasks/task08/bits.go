package main

import "fmt"

/*
SetBit выставляет бит на заданной позиции в 1, если он равен 0, в противном же
случае оставляет его как есть. это достигается битовым сдвигом единицы в
специальном числе-маске на указанную позицию и затем наложения этой маски на
число которое следует изменить с помощью битового оператора "или". принцип его
работы аналогичен логическому оператору "или" - логическому сложению. в данном
случае в результирующем числе на заданной позиции бит выставляется в единицу в
том случае, если хотя бы в одном из исходных чисел данный бит выставлен в
единицу, а это гарантируется созданным числом-маской.
*/
func SetBit(num int64, bitPosition int) int64 {
	return num | (1 << bitPosition)
}

/*
ResetBit выставляет бит на заданной позиции в 0. в нём по аналогии с предыдущей
функцией создаётся битовая маска с битом, выставленным в единицу на указанной
позиции, а затем она накладывается на оригинальное число и в результате
получается число с битом выставленным в ноль на указанной позиции. достигается
это с помощью сочетания битовых операторов & - "и" (единиуа в бите на указанной
позиции результирующего числа возможна только если оба бита на указанной позиции
оригинальных чисел выставлены в единцу) и ^ - "исключающее или" (единица в бите
на указанной позиции результирующего числа возможно только если ровно в одном
бите оригинальных чисел выставлена единица) - такое сочетание называется битовым
сбросом и может называться AND NOT - бит в результирующем числе на указанной
позиции будет в любом случае равен нулю, так как при помощи данного оператора,
если в правом операнде указанный бит установлен в единицу, то вне зависимости от
того какой бит на указанной позиции находится в левом операнде - в результирующем
числе этот бит будет равен нулю и обратно - если в правом операнде бит выставлен
в ноль, то бит в правом операнде останется неизменным
*/
func ResetBit(num int64, bitPosition int) int64 {
	return num &^ (1 << bitPosition)
}

/*
IsSetBit проверяет что бит на указанной позиции в числе выставлен в единицу при
помощи итового оператора & - "и", который "возвращает" единицу только в том
случае когда бит оригинального числа и маски оба имеют единицу в указанном бите.
возрващается результат сравнения маски с оригинальным числом с нулём - то есть
если бит выставлен в оригинальном числе в единицу, то результатом битовой
операции будет единица, которая не равна нулю и результатом работы функции будет
true.
*/
func IsSetBit(num int64, bitPosition int) bool {
	return (num & (1 << bitPosition)) != 0
}

func main() {
	var num int64

	num = SetBit(num, 1)

	fmt.Println("Bit number 1 is set =", IsSetBit(num, 1))

	num = ResetBit(num, 1)

	fmt.Println("Bit number 1 is set =", IsSetBit(num, 1))
}
