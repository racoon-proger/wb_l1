package main

import (
	"context"
	"log"
	"os/signal"
	"sync"
	"syscall"
)

type counter struct {
	num int
	sync.Mutex
}

// метод структуры counter, инкрементит поле num
func (c *counter) inc() {
	c.Lock()
	c.num++
	c.Unlock()
}

// метод структуры counter, возвращает значения num
func (c *counter) value() int {
	c.Lock()
	output := c.num
	c.Unlock()
	return output
}
func main() {
	c := &counter{}
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	incrementConcurrently(ctx, c)

	log.Println(c.value())
}

func incrementConcurrently(ctx context.Context, c *counter) {
	wg := sync.WaitGroup{}
	for i := 0; i < 15; i++ {
		select {
		case <-ctx.Done():
			return
		default:
			wg.Add(1)
			go func() {
				defer wg.Done()
				c.inc()
			}()
		}
	}
	wg.Wait()
}
