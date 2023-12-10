package main

import (
	"fmt"
	"sync"
)

/*
Counter является структурой-счётчиком со "встроенным" sync.RWMutex, который
позволяет безопасно изменять поле count структуры из разных горутин
*/
type Counter struct {
	sync.RWMutex
	count int
}

/*
NewCounter возвращает указатель на новый экземпляр типа Counter
*/
func NewCounter() *Counter {
	return &Counter{}
}

/*
Increment "блокирует", либо же встаёт в очередь на блокировку общей памяти
горутин для безопасного увеличения значения поля count
*/
func (c *Counter) Increment() {
	c.Lock()
	c.count++
	c.Unlock()
}

/*
Decrement "блокирует", либо же встаёт в очередь на блокировку общей памяти
горутин для безопасного уменьшения значения поля count
*/
func (c *Counter) Decrement() {
	c.Lock()
	c.count--
	c.Unlock()
}

/*
Count печатает в stdout текущее значение поля count, предварительно блокируя на
чтение память в которой находится данная переменная.
*/
func (c *Counter) Count() {
	c.RLock()
	fmt.Println("current Counter state =", c.count)
	c.RUnlock()
}

/*
Reset блокирует память поля count и сбрасывает его на стандартное значение,
равное нулю
*/
func (c *Counter) Reset() {
	c.Lock()
	c.count = 0
	c.Unlock()
}

/*
SetVal блокирует память поля count и присваивает ему переданное значение, в
обход операций инкремента, или декремента
*/
func (c *Counter) SetVal(val int) {
	c.Lock()
	c.count = val
	c.Unlock()
}

func main() {
	counter := NewCounter()
	var wg sync.WaitGroup

	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func(c *Counter) {
			c.Increment()
			wg.Done()
		}(counter)
	}
	wg.Wait()

	counter.Count()
}
