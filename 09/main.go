package main

import (
	"fmt"
)

var arr = []int{2, 4, 6, 8, 10}

func main() {
	// создали 2 канала
	chanOne := make(chan int)
	chanTwo := make(chan int)

	go func() {
		// с помощью цикла помещаем данные из слайса в канал chanOne
		for i := 0; i < len(arr); i++ {
			chanOne <- arr[i]
		}
		close(chanOne)
	}()

	go func() {
		// с помощью range достаем данные из канала chanOne умножаем
		// их на 2 и помещаем в канал chanTwo
		for num := range chanOne {
			chanTwo <- num * 2
		}
		close(chanTwo)
	}()
	//  достаем и выводим данные из канала chanTwo
	for v := range chanTwo {
		fmt.Printf("value is : %d\n", v)
	}
}
