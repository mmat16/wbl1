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

func main() {
	var wg sync.WaitGroup

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	wg.Add(1)
	go stopWithTimeout(ctx, &wg)

	wg.Wait()
}
