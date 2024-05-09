// паттерн graceful shutdown.
// Для реализации данного паттерна воспользуйтесь каналами и оператором select с default-кейсом.
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func square(ch chan int, quit chan os.Signal) {
	x := 1
	for {
		select {
		case <-quit:
			fmt.Println("выхожу из программы")
			return
		default:
			ch <- x * x
			x++
			time.Sleep(1 * time.Second)
		}
	}
}

func main() {
	ch := make(chan int)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM) //Буферизированный канал используется в этом коде для обеспечения того, что сигналы операционной системы (os.Interrupt, syscall.SIGTERM) могут быть корректно обработаны, даже если главная горутина занята и не может немедленно принять сигнал.

	go func() {
		for {
			fmt.Println(<-ch)
		}
	}()

	square(ch, quit)
}
