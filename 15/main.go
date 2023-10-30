package main

import (
	"log"
	"strconv"
	"strings"
)

func main() {
	str := createHugeStringFast(1 << 10)
	log.Println(str)
}

// здесь при конкатенации строка каждый раз пересоздается
// при конкатенации с новым символом
func createHugeStringSlow(length int) string {
	var output string
	for i := 0; i < length; i++ {
		output += strconv.Itoa(i)
	}
	return output
}

// здесь используется слайс байт, и память аллоцируется
// только когда кол-во элементов превышает capacity слайса
func createHugeStringFast(length int) string {
	// создаем билдер строк
	var builder strings.Builder
	for i := 0; i < length; i++ {
		// пишем символ в билдер
		_, _ = builder.WriteString(strconv.Itoa(i))
	}
	// возвращаем результат
	return builder.String()
}
