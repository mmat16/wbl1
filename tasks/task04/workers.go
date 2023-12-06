package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"sync"
	"time"
)

/*
константы для хранения сообщений о пользовании флагом и ошибке пользователю.
*/
const (
	infoMsg  = "number of workers reading from main channel\nshould be equal or greater than 1"
	panicMsg = "Number of workers reading should be equal or greater than 1"
)

/*
переменная для хранения количетсва воркеров, получающих сообщения из основного
канала.
*/
var numWorkers uint

/*
функция инициализации (выполняется перед функцией main) устанавливает либо
стандартное количество воркеров, либо количество, указанное пользователем. а
так же проверяет валидность значения - воркеров не может быть менее 1.
*/
func init() {
	flag.UintVar(&numWorkers, "w", 4, infoMsg)
	flag.Parse()
	if numWorkers < 1 {
		panic(panicMsg)
	}
}

/*
main функция создаёт канал для сигналов операционной системы и связывает signal
interrupt (ctrl+c) с этим каналом. Затем создаётся канал для передачи сообщений
из "основной" горутины воркерам с буфером, соответствующим количеству воркеров.
затем создаётся вейтгруппа и запускается горутина для чтения из канала quit.
затем создаются и запускаются воркеры. и в цикле используется оператор select
чтобы обрабатывать закрытие канала передачи сообщений и ожидания завершения
работы воркеров перед завершением программы, либо же для передачи нового
сообщения воркерам.
я выбрал данный способ завершения работы, потому что он гарантирует то что все
воркеры успеют отправить сообщение в stdout до того как "основная" горутина
завершится и ни одно сообщение не будет потеряно.
*/
func main() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)

	messages := make(chan int, numWorkers)
	var wg sync.WaitGroup

	go WaitForInterrupt(quit, messages)

	Workers(numWorkers, messages, &wg)

	for {
		select {
		case <-messages:
			fmt.Println("waiting for workers to be done")
			wg.Wait()
			fmt.Println("quiting")
			return
		default:
			messages <- rand.Int()
			fmt.Println("sent new message")
			time.Sleep(time.Millisecond * 100)
		}
	}
}

/*
WaitForInterrupt ждёт пока в канал quit придёт signal interrupt (ctrl+c), после
чего закрывает канал messages, использующийся для передачи сообщений воркерам
*/
func WaitForInterrupt(quit <-chan os.Signal, messages chan int) {
	<-quit
	fmt.Println("received signal interrupt. closing messages channel")
	close(messages)
}

/*
Workers создаёт необходимое количество воркеров и инкрементирует счётчик
WaitGroup на это число.
*/
func Workers(numWorkers uint, ch <-chan int, wg *sync.WaitGroup) {
	wg.Add(int(numWorkers))
	for i := 0; i < int(numWorkers); i++ {
		go Work(ch, i+1, wg)
	}
}

/*
Work в цикле читает значения из канала, проверяя не закрыт ли он. Если канал не
закрыт и сообщение прочитать удалось - оно печатается в stdou. Если прочитать
очередное значение не удалось, цикл прерывается и, перед выходом из функции,
срабатывает отложенная функция wg.Done, декрементирующая счётчик sync.WaitGroup
*/
func Work(ch <-chan int, number int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		num, ok := <-ch
		if !ok {
			fmt.Printf("Worker %d: done\n", number)
			break
		}
		fmt.Printf("Worker %d says: %d\n", number, num)
	}
}
