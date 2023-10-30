package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		wg  = &sync.WaitGroup{}
		arr = []int{2, 4, 6, 8, 10}
	)
	// устанавливаем кол-во элементов в WaitGroup.
	// здесь кол-во горутин равно длине массива
	wg.Add(len(arr))
	for i := 0; i < len(arr); i++ {
		// в каждой итерации цикла
		// вызываем горутину которая выводит квадрат числа
		go func(i int) {
			fmt.Println(arr[i] * arr[i])
			// сообщаем о том что горутина завершила работу
			wg.Done()
		}(i)
	}
	// дожидаемся выполнения всех горутин
	wg.Wait()
}
