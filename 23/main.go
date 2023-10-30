package main

import "log"

func main() {
	arr := []int{1, 2, 3, 4, 5, 6, 7}
	arr = remove(arr, 1)
	log.Println(arr)
	arr = removeWithoutOrder(arr, 1)
	log.Println(arr)
}

// здесь мы берем два сабслайса и аппендим второй к первому,
// при этом откидываем элемент который нужно удалить
func remove(arr []int, index int) []int {
	return append(arr[:index], arr[index+1:]...)
}

// здесь мы заменяем элемент который надо удалить на последний
// и делаем сабслайс отбрасывая последний элемент
func removeWithoutOrder(arr []int, index int) []int {
	arr[index] = arr[len(arr)-1]
	return arr[:len(arr)-1]
}
