package main

import "time"

// Sleep останавливает выполнение программы на заданное количество секунд при
// помощи time.After - канала, который закрывается через указанное количество
// секунд. Таким образом выполняемая горутина блокируется до тех пор, пока канал,
// не будет закрыт
func Sleep(seconds int) {
	<-time.After(time.Second * time.Duration(seconds))
}

func main() {
	Sleep(5)
}
