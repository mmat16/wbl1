package main

import "fmt"

/*
MakeSet создаёт множество с уникальными элементами взодного слайса за линейное
время при помощи мапы, ключами которой являются элементы входного массива. так
как ключами любой мапы могут быть только уникальные значения - это гарантирует
уникальность элементов в результирующем слайсе, представляющем множество.
*/
func MakeSet(objects []string) []string {
	set := make(map[string]bool)
	for _, obj := range objects {
		set[obj] = true
	}
	var res []string
	for k := range set {
		res = append(res, k)
	}
	return res
}

func main() {
	objects := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(MakeSet(objects))
}
