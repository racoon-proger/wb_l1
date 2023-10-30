package main

import (
	"context"
	"fmt"
	"time"
)

func worker(ctx context.Context) {
	for {
		select {
		default:
			// Выполнение работы в горутине
			fmt.Println("выполняется")
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			// Получен сигнал об остановке, горутина завершается
			fmt.Println("завершена")
			return
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	go worker(ctx)

	<-ctx.Done()
	time.Sleep(time.Second)
}
