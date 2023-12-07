package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

/*
stopWithTimeout останавливает горутину по истечении таймаута контекста (или
другого фактора, обуславливающего его длительность жизни).
в цикле используется оператор select для проверки состояния контекста - если
время его жизни истекло, либо же на стороне вызова была вызвана функция cancel,
сгенерированная вместе с контекстом - выполнение функции stopWithTimeout
прерывается. Иначе выполняется ветка default, в которой происходит вывод о
работе функции.

Так же в целом возможно использование контекста с отменой context.WithCancel,
контекста с дедлайном context.WithDeadline и др виды контекстов.
*/
func stopWithTimeout(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			fmt.Println("stopWithTimeout stopped")
			return
		default:
			fmt.Println("stopWithTimeout running")
			time.Sleep(time.Second)
		}
	}
}

/*
stopThruChannel использует для остановки канал, принимаюший пустой интерфейс. В
цикле с помощью оператора select проверяется что в канал done отправлено
сообщение и в таком случае завершает свою работу. Иначе срабатывает ветка
default и в канал сообщений отправляется сообщение о работе функции.
*/
func stopThruChannel(msg chan string, done <-chan struct{}) {
	for {
		select {
		case <-done:
			fmt.Println("stopThruChannel stopped")
			return
		default:
			msg <- "stopThruChannel running"
			time.Sleep(time.Second)
		}
	}
}

/*
stopThruClosingChannel использует конструкцию for range для канала, которая при
очередном считывании значений из канала проверяет что канал не был закрыт.
Если канал закрыт, цикл завершается, в stdout пишется сообщение об этом и
функция завершается
*/
func stopThruClosingChannel(msg <-chan string, wg *sync.WaitGroup) {
	defer wg.Done()
	for message := range msg {
		fmt.Println(message)
	}
	fmt.Println("stopThruClosingChannel stopped")
}

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg.Add(1)
	go stopWithTimeout(ctx, &wg)

	wg.Wait()

	msg := make(chan string)
	done := make(chan struct{})

	go stopThruChannel(msg, done)

	for i := 0; i < 5; i++ {
		fmt.Println(<-msg)
	}
	done <- struct{}{}

	wg.Add(1)
	go stopThruClosingChannel(msg, &wg)

	for i := 0; i < 5; i++ {
		msg <- "stopThruClosingChannel running"
		time.Sleep(time.Second)
	}
	close(msg)
	wg.Wait()
}
