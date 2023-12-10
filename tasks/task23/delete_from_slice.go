package main

import "fmt"

/*
DeleteFromSliceOrdered удаляет элемент на указанной позиции в исходном слайсе и
возвращает результат этого удаления. функция сохраняет порядок элементов в
слайсе и является не очень эффективной (временная сложность O(n) в худшем
случае, то есть когда элемент удаляется не из конца массива), так как перемещает
все элементы после удалённого на одну позицию влево.
*/
func DeleteFromSliceOrdered[T any](slice []T, pos int) []T {
	return append(slice[:pos], slice[pos+1:]...)
}

/*
DeleteFromSliceUnordered следует использовать в случаях когда важна скорость
удаления элемента, но не важен порядок элементов, так как она просто
переставляет местами последний элемент с тем что следует удалить и возвращает
срез от оригинального слайса с первого по предпоследний элемент. временная
сложность данного подхода константна - O(1)

разницу производительности можно наглядно заметить на выполении бенчмарка:
❯ go test -bench="." -benchmem
goos: darwin
goarch: amd64
pkg: wb01/tasks/task23
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkDeleteFromSliceOrdered-8     	72529858	        17.43 ns/op	       0 B/op	       0 allocs/op
BenchmarkDeleteFromSliceUnordered-8   	175494182	         6.834 ns/op	       0 B/op	       0 allocs/op
PASS
ok  	wb01/tasks/task23	4.852s
*/
func DeleteFromSliceUnordered[T any](slice []T, pos int) []T {
	slice[pos] = slice[len(slice)-1]
	return slice[:len(slice)-1]
}

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(DeleteFromSliceOrdered(slice1, 5))
	fmt.Println(DeleteFromSliceUnordered(slice2, 5))
}
