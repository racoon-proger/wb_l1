package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type configuration struct {
	NumWorkers int `envconfig:"NUM_WORKERS" default:"10"`
}

func main() {
	var cfg configuration
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal(err)
	}

	// создаем контекст, который отменится при нажатии ctrl+c
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	// создаем канал
	var newChan = make(chan int)
	// defer'ом закрываем канал, чтобы завершился цикл в горутинах
	defer close(newChan)

	// кол-во запускаемых горутин-воркеров равно значению поля NumWorkers
	for i := 0; i < cfg.NumWorkers; i++ {
		go func(i int) {
			// range будет работать до тех пор пока канал не пуст и не закрыт
			for data := range newChan {
				log.Printf("i am worker number: %d. data: %d", i, data)
			}
			log.Printf("worker %d finished", i)
		}(i)
	}

	var num int
	for {
		select {
		// проверяем, завершен ли контекст
		case <-ctx.Done():
			log.Println("program finished")
			return
		default:
			newChan <- num
			num++
			time.Sleep(time.Millisecond * 500)
		}
	}
}
