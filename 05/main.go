package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// создаем контекст, который отменится через определенное время
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	// создаем канал
	newChan := make(chan int)
	defer close(newChan)

	go worker(newChan)
	num := 0
	for {
		select {
		// проверяем, завершен ли контекст
		case <-ctx.Done():
			return
		default:
			// записываем данные в канал
			newChan <- num
			num++
			time.Sleep(time.Millisecond * 300)
		}
	}
}

func worker(newChan chan int) {
	// читаем данные из канала
	for data := range newChan {
		fmt.Println(data)
	}
}
