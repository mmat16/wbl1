package main

import "fmt"

/*
DefineTypeExt использует конструкцию switch-type для определения типа переменной,
переданной как пустой интерфейс (сокращённо называемый any).
*/
func DefineTypeExt(item any) string {
	var msg string

	switch t := item.(type) {
	default:
		msg = fmt.Sprintf("value has type %T", t)
	}

	return msg
}

/*
DefineType проверяет входящую any переменную на соответсвие одному из типов и в
случае когда переменная не является ни одним из них - возврашает сообщение об
этом. Для определение типа так же используется конструкция switch-type
*/
func DefineType(item any) string {
	var msg string

	switch t := item.(type) {
	case int:
		msg = fmt.Sprintf("value has type %T", t)
	case string:
		msg = fmt.Sprintf("value has type %T", t)
	case bool:
		msg = fmt.Sprintf("value has type %T", t)
	case chan any:
		msg = fmt.Sprintf("value has type %T", t)
	default:
		msg = "unknown type"
	}

	return msg
}

func main() {
	var item map[int]bool
	fmt.Println(DefineType(item))
	fmt.Println(DefineTypeExt(item))
}
