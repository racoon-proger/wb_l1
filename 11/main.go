package main

import "fmt"

func intersection(arrOne, arrTwo []int) []int {
	counter := make(map[int]int)
	result := []int{}
	// посчитали количество появлений элементов первого слайса,
	// в ключе мапы элемент, в значении число его появлений в слайсе
	for _, elem := range arrOne {
		counter[elem]++
	}
	// проверяем есть ли совпадающие элементы в обоих слайсах и если есть
	// записываем их в результирующий слайс
	for _, elem := range arrTwo {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}
	return result
}

func main() {
	arrOne := []int{2, 16, 21, 14, 81}
	arrTwo := []int{4, 75, 21, 64, 2, 35, 16}
	fmt.Println(intersection(arrOne, arrTwo))
}
