package main

import (
	"fmt"
	"sync"
)

func main() {
	var (
		n      = 200
		mx     = sync.Mutex{}
		record = make(map[int]int, n)
		wg     = sync.WaitGroup{}
	)
	// установили счетчик запускаемых горутин
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func(i int) {
			// заблокировали переменную record, что бы производить
			//  запись в нее могла только одна горутина
			mx.Lock()
			record[i] = i
			// разблокировали переменную record
			mx.Unlock()
			wg.Done()
		}(i)
	}
	// ожидаем завершения всех горутин
	wg.Wait()

	fmt.Println(len(record))
}
