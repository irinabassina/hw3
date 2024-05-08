package main

import (
	"fmt"
	"sync"
)

func main() {

	squareCh := make(chan int)
	multiplyCh := make(chan int)

	for {
		fmt.Println("Пожалуйста, введите целое число для вычисления результата или введите любые символы отличные от цифр для завершения программы")
		var i int
		_, err := fmt.Scanf("%d", &i)
		if err != nil {
			return
		}

		var wg sync.WaitGroup

		wg.Add(2)
		go square(squareCh, multiplyCh, &wg)
		go multiplyByTwo(multiplyCh, &wg)

		squareCh <- i
		wg.Wait()
		fmt.Printf("main: все горутины завершилиcь\n")
	}
}

func square(squareCh chan int, multiplyCh chan int, wg *sync.WaitGroup) {
	num := <-squareCh
	fmt.Println("square: квадрат числа = ", num*num)
	wg.Done()
	multiplyCh <- num * num
}

func multiplyByTwo(multiplyCh chan int, wg *sync.WaitGroup) {
	num := <-multiplyCh
	fmt.Println("multiplyBy2: квадрат х 2 = ", num*2)
	wg.Done()
}
