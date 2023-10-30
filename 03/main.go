package main

import (
	"fmt"
	"sync"
)

var (
	mu  = &sync.Mutex{}
	wg  = &sync.WaitGroup{}
	arr = []int{2, 4, 6, 8, 10}
)

func main() {
	// устанавливаем кол-во элементов в WaitGroup.
	// здесь кол-во горутин равно длине массива
	wg.Add(len(arr))
	sum := 0

	for i := 0; i < len(arr); i++ {
		// в каждой итерации цикла
		// вызываем горутину которая прибавляет к переменной sum квадрат числа i
		go func(i int) {
			//
			mu.Lock()
			sum += arr[i] * arr[i]
			mu.Unlock()
			// сообщаем о том что горутина завершила работу
			wg.Done()
		}(i)
	}
	// дожидаемся выполнения всех горутин
	wg.Wait()
	fmt.Println(sum)
}
