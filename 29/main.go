package main

import "fmt"

func main() {
	n := 0
	if true {
		n := 1
		n++
	}
	fmt.Println(n)
}

// выведется 0  потому что переменная n в блоке if другая
