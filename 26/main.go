package main

import (
	"log"
	"strings"
)

func checkUniqueChars(input string) bool {
	m := make(map[string]int)
	for _, elem := range input {
		// приводим элементы к нижнему регистру и кладем их в ключ мапы
		m[strings.ToLower(string(elem))]++
		// если елемент встречается больше одного раза возвращаем false
		if m[strings.ToLower(string(elem))] > 1 {
			return false
		}
	}
	return true
}

func main() {
	var str string = "aLlon"
	log.Println(checkUniqueChars(str))
}
