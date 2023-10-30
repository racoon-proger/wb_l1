package main

import "fmt"

func main() {
	slice := []string{"a", "a"}

	func(slice []string) {
		slice = append(slice, "a")
		slice[0] = "b"
		slice[1] = "b"
		fmt.Print(slice)
	}(slice)
	fmt.Print(slice)
}

// вывод: [b b a] [a a], в анонимную функцию копируется slice,
// с помощью функции append добавляется элемент "а
// и после этого capacity увеличивается вдвое, следовательно массив внутри слайса меняется
// затем меняется значения под индексами 0,1
// в конечном итоге мы имеем 2 разных слайса [a a] и [b b a]
