package main

import "log"

var arr = []int{1, 4, 12, 24, 56, 78}

func binarySearch(arr []int, requiredNum int) int {
	// определил границы поиска:
	// начальная граница
	leftPointer := 0
	// конечная граница
	rightPointer := len(arr) - 1

	for leftPointer <= rightPointer {
		// нашел средний элемент между начальной и конечной границей
		midPointer := int((leftPointer + rightPointer) / 2)
		midValue := arr[midPointer]
		// сравнил средний элемент с  искомым значением
		if midValue == requiredNum {
			return midPointer
		} else if midValue < requiredNum {
			leftPointer = midPointer + 1
		} else {
			rightPointer = midPointer - 1
		}
	}
	return -1
}
func main() {

	log.Println(binarySearch(arr, 1))
}
