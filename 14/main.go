package main

import "fmt"

func checkType(x any) {
	switch x.(type) {
	case int:
		fmt.Println("x is int")
	case string:
		fmt.Println("x is string")
	case bool:
		fmt.Println("x is bool")
	case chan string:
		fmt.Println("x is chan string")
	default:
		fmt.Println("type unknown")
	}
}

func main() {
	checkType("faw")

}
