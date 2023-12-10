package sum_slice_squares

import "sync"

/*
Данный вариант использует канал для передачи квадратов отдельных чисел слайса
из дополнительных горутин в "основную".
Выполнение всех дополнительных горутин гарантируется тем что операции чтения
происходят в цикле, колоичество итераций которого соответсвует количеству
итераций первого цикла в котором и запускаются дополнительные горутины.
Во втором цикле прочитанные из канала значения суммируются в результирующую
переменную.
*/
func SumSliceSquares(nums []int) int {
	results := make(chan int, len(nums))
	for _, num := range nums {
		go func(num int, results chan<- int) {
			results <- num * num
		}(num, results)
	}

	var res int

	for range nums {
		res += <-results
	}
	return res
}

/*
Этот вариант использует мьютекс и вейтгруппу для синхронизации горутин и
избежания состояния гонки - результаты промежуточных вычислений квадратов
элементов массива записываются сразу в результирующую переменную, предварительно
залоченную мьютексом.
*/
func SumSliceSquaresMu(nums []int) int {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var res int

	wg.Add(len(nums))
	for _, num := range nums {
		go func(num int) {
			mu.Lock()
			res += num * num
			mu.Unlock()
			wg.Done()
		}(num)
	}
	wg.Wait()

	return res
}

/*
и снова - как и в предыдущем задании, использование пакета sync обходит по
скорости использование каналов, но использование каналов в свою очеред экономит
память.

❯ go test -bench="." -benchmem -run="kw"
goos: darwin
goarch: amd64
pkg: wb01/tasks/task03
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkSumSliceSquares-8     	  131148	      9293 ns/op	     816 B/op	      19 allocs/op
BenchmarkSumSliceSquaresMu-8   	  198459	      6074 ns/op	    1040 B/op	      39 allocs/op
PASS
ok  	wb01/tasks/task03	5.338s
*/
