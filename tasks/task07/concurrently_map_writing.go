package main

import (
	"fmt"
	"sync"
)

/*
WriteWithMutex использует мьютекс для конкуретной записи данных в map. в самом
начале фунцкии вызывается метод Lock, который блокирует общий для горутин
участок памяти на запись и чтение, после чего записывает в мапу значение по
переданному в функцию ключу. после этого в отложенных функциях мьютекс
разблокирует память и декрементирует счётчик вейтгруппы давая понять, что работа
функции WriteWithMutex завершена.
*/
func WriteWithMutex(mu *sync.Mutex, wg *sync.WaitGroup, storage map[int]int, key int) {
	mu.Lock()
	defer mu.Unlock()
	defer wg.Done()
	storage[key]++
}

/*
WriteWithSyncMap использует sync.Map для конкуретной записи. sync.Map является
структурой, содержащей в себе обычную map и методы для обеспечения
безопасного доступа к общей памяти мапы из разных горутин.
sync.Map была реализована в стандартной библиотеке го для решения нескольких
проблем, одна из них - линейное увеличение времени чтения из мапы с обычным
sync.Mutex, либо даже sync.RWMutex при наличии большого количества горутин и
задействованных ядер системы, работающих с памятью этой мапы, так как большое
количество времени тратится на выстраивание очередей не только на запись, или
чтение "полезной" информации - доступа к общей памяти, но и на переключения
статуса мьютекса. Использование sync.Map оправдано при условии что память будет
читаться большим количеством горутин на системе с большим количеством ядер
процессора - на системе до 4 ядер процессора sync.Map проигрывает по скорости
обычной мапе с мьютексом, но гарантирует константное время доступа к памяти при
любом количестве горутин и ядер процессора, в то время как у обычной мапы с
мьютексом эта скорость начинает расти линейно на системах с количеством ядер
более 4.
*/
func WriteWithSyncMap(storage *sync.Map, wg *sync.WaitGroup, key int) {
	defer wg.Done()
	storage.Store(key, 1)
}

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	storage := make(map[int]int, 5)

	fmt.Println("using sync.Mutex")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go WriteWithMutex(&mu, &wg, storage, i+1)
	}
	wg.Wait()

	for k, v := range storage {
		fmt.Printf("key %d, val %d\n", k, v)
	}

	m := &sync.Map{}

	fmt.Println("using sync.Map")
	wg.Add(5)
	for i := 0; i < 5; i++ {
		go WriteWithSyncMap(m, &wg, i+1)
	}
	wg.Wait()

	m.Range(func(key, value any) bool {
		fmt.Printf("key %d, val %d\n", key.(int), value.(int))
		return true
	})
}

/*
и информация подтверждается. на моей системе как раз 4 ядра и вариант функции с
sync.Map в данном случае проигрывает по скорости обычной мапе с мьютексом -
разница почти в три раза, что по объёму используемой памяти, что и по времени
выполнении операции.

❯ go test -bench="." -benchmem
goos: darwin
goarch: amd64
pkg: wb01/tasks/task07
cpu: Intel(R) Core(TM) i5-8257U CPU @ 1.40GHz
BenchmarkWriteWithMutex-8     	 1500722	       698.3 ns/op	     134 B/op	       1 allocs/op
BenchmarkWriteWithSyncMap-8   	 1000000	      1695 ns/op	     321 B/op	       5 allocs/op
PASS
ok  	wb01/tasks/task07	4.694s
*/
