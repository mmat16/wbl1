package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"time"
)

// переменная для хранения времени в секундах, в течение которого будет
// происходить передача сообщений в канал
var duration int

/*
функция инициализации устанавливает либо стандартное время передачи сообщений
(3 секунды), либо назначает пользовательское время работы переданное через
аргументы командной строки и валидирует полученое значение - время работы должно
быть больше или равно одной секунде
*/
func init() {
	flag.IntVar(&duration, "t", 3, "program uptime in seconds")
	flag.Parse()
	if duration < 1 {
		panic("uptime should be equal or greater than one second")
	}
}

/*
main функция создаёт контекст с таймаутом, равным пользовательскому или
стандартному времени работы, затем создаёт канал для передачи сообщений и в
отдельной горутине начинает отправлять сообщения в него. основная горутина
читает значения из канала
*/
func main() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(duration)*time.Second)
	defer cancel()

	messages := make(chan int)

	go Send(ctx, messages)

	Receive(messages)
}

/*
Send принимает контекст и канал на запись для передачи сообщений и в цикле с
помощью оператора select проверяет - если контекст "истёк" то сообщает об этом
в stdout, закрывает канал и завершается, иначе отправляет в канал очередное
значение и "засыпает" на одну секунду
*/
func Send(ctx context.Context, messages chan<- int) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("sender saying: \"time is up!\"")
			close(messages)
			return
		default:
			messages <- rand.Int()
			time.Sleep(time.Second)
		}
	}
}

/*
Receive принимает канал на чтение и в цикле проверяет чтение из него - если
канал не закрыт, то выводит полученое сообщение в stdout, иначе сообщает туда же
о том что канал закрыт пишущей стороной и завершается.
*/
func Receive(messages <-chan int) {
	for {
		msg, ok := <-messages
		if !ok {
			fmt.Println("receiver complaining: \"channel is closed!\"")
			return
		}
		fmt.Println("received message:", msg)
	}
}

/*
важно закрывать канал именно со стороны "писателя", так как попытка записи в
закрытый канал приводит к панике. в нашем случае канал закрывается "безопасно" -
единственной функцией, пишущей в этот канал, которая завершается сразу после его
закрытия.
*/
