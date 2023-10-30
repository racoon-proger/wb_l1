package main

import "fmt"

func main() {
	arr := []string{"cat", "cat", "dog", "cat", "tree"}
	m := make(map[string]struct{})
	result := []string{}
	for _, elem := range arr {
		if _, ok := m[elem]; !ok {
			result = append(result, elem)
		}
		m[elem] = struct{}{}
	}
	fmt.Println(result)
}
