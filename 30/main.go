package main

import "fmt"

func someAction(v []int8, b int8) {
	/*
		type sliceHeader struct {
			len int = 5
			cap int = 5
			ptr *[]int8{1, 2, 3, 4, 5}
		}
	*/
	v[0] = 100
	/*
		type sliceHeader struct {
			len int = 5
			cap int = 5
			ptr *[]int8{100, 2, 3, 4, 5}
		}
	*/
	v = append(v, b)
	/*
		type sliceHeader struct {
			len int = 6
			cap int = 10
			ptr *[]int8{100, 2, 3, 4, 5, 6}
		}
	*/
}

func main() {
	// var a = []int8{1, 2, 3, 4, 5}
	a := make([]int8, 0, 5)
	/*
		type sliceHeader struct {
			len int = 0
			cap int = 5
			ptr *[]int8{}
		}
	*/
	a = append(a, 1)
	a = append(a, 2)
	a = append(a, 3)
	a = append(a, 4)
	a = append(a, 5)
	/*
		type sliceHeader struct {
			len int = 5
			cap int = 5
			ptr *[]int8{1, 2, 3, 4, 5}
		}
	*/

	someAction(a, 6)
	/*
		type sliceHeader struct {
			len int = 5
			cap int = 5
			ptr *[]int8{100, 2, 3, 4, 5}
		}
	*/

	fmt.Println(a)
}

// вывод: [100 2 3 4 5] потому что в функции someAction
// мы поменяли первый элемент в массиве на который указывает слайс "a"
// потом при аппенде, создался новый массив но уже в копии слайса,
// поэтому это не повлияло на сам слайс "a"
