package main

import (
	"log"
	"sort"
)

func main() {
	// 1 способ
	arr := []int{2, 1, 66, 43, 13, 11}
	sort.Ints(arr)
	log.Println(arr)

	// 2 способ
	arr = []int{2, 1, 66, 43, 13, 11}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})
	log.Println(arr)

}
