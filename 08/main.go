package main

import "fmt"

func inverseByte(num, index int) int {
	// создаем маску путем сдвига
	mask := 1 << index
	// применяем операцию логического или,
	// которая возвращает 1 только если биты разные
	return num ^ mask
}
func main() {
	num := 5
	fmt.Printf("%b\n", num)
	num = inverseByte(num, 2)
	fmt.Printf("%b\n", num)

}
