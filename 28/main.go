package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg sync.WaitGroup, i int) {
			fmt.Println(i)
			wg.Done()
		}(wg, i)
	}
	wg.Wait()
	fmt.Println("exit")
}

// будет дедлок, потому что мы коприуем wait group, и горутина
// вызывает метод Done() уже у другого wg.
// и wg.Wait() заблочит нас навсегда
