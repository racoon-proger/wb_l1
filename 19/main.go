package main

import (
	"log"
	"strings"
)

func reverse(str string) (string, error) {
	builder := strings.Builder{}
	// приводим строку к слайсу рун
	arr := []rune(str)
	// идем по слайсу справа налево
	for i := len(arr) - 1; i >= 0; i-- {
		// пишем в билдер руну
		_, err := builder.WriteRune(arr[i])
		if err != nil {
			return "", err
		}
	}
	// возвращаем результат
	return builder.String(), nil
}

func main() {
	var str = "alan"
	result, err := reverse(str)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(result)
}
