package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "a b c d"
	arr := strings.Fields(str)

	// i это левый указатель, opp правый
	// i и opp начинают с центра и идут в свои стороны
	// меняя элементы местами
	// пример результатов итераций:
	// 1: "a", "b", "c", "d"
	// 2: "a", "c", "b", "d"
	// 3: "d", "c", "b", "a"
	for i := len(arr)/2 - 1; i >= 0; i-- {
		opp := len(arr) - 1 - i
		arr[i], arr[opp] = arr[opp], arr[i]
	}
	// формируем строку из слайса строк с пробелом в качестве разделителя
	result := strings.Join(arr, " ")
	fmt.Println(result)
}
