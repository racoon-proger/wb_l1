package main

import (
	"fmt"
	"sync"
	"time"
)

// реализация остановки горутины с использованием канала для передачи
// сигнала об остановке
func worker(wg *sync.WaitGroup, done chan bool) {
	defer wg.Done()
	for {
		select {
		default:
			// Выполнение работы в горутине
			fmt.Println("выполняется")
			time.Sleep(1 * time.Second)
		case <-done:
			// Получен сигнал об остановке, горутина завершается
			fmt.Println("завершена")
			return
		}
	}
}

func main() {
	// создаем канал
	done := make(chan bool)
	defer close(done)

	wg := sync.WaitGroup{}
	wg.Add(1)
	go worker(&wg, done)

	go func() {
		// после опредленного времени,
		// решаем завершить горутину
		time.Sleep(5 * time.Second)
		done <- true
	}()

	wg.Wait()
}
