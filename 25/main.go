package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGTERM, syscall.SIGINT)
	defer cancel()

	log.Println("start sleeping")
	sleepContext(ctx, time.Second*10)
	log.Println("finish")
}

// по истечении duration из канала который возвращается
// методом After прилетит элемент, до тех пор мы будем заблокированы
func sleep(duration time.Duration) {
	<-time.After(duration)
}

// канал тикера периодически через каждые duration,
// посылает сигнал, но мы выйдем из функции после первого сигнала,
// так как нам этого достаточно
func sleep2(duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	<-ticker.C
}

func sleepContext(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
	case <-time.After(duration):
	}
}

func sleep2Context(ctx context.Context, duration time.Duration) {
	ticker := time.NewTicker(duration)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
	case <-ticker.C:
	}
}
