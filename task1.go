// Реализуйте паттерн-конвейер:  Программа принимает числа из стандартного ввода в бесконечном цикле и передаёт число в горутину.
// Квадрат: горутина высчитывает квадрат этого числа и передаёт в следующую горутину.
// Произведение: следующая горутина умножает квадрат числа на 2.
// При вводе «стоп» выполнение программы останавливается.

package main

import (
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {

	squareCh := make(chan int)
	multiplyCh := make(chan int)

	for {
		i, done := getInput()
		if done {
			return
		}

		var wg sync.WaitGroup

		wg.Add(2)
		go square(squareCh, multiplyCh, &wg)
		go multiplyByTwo(multiplyCh, &wg)

		squareCh <- i
		wg.Wait()
		fmt.Println("main: все горутины завершилиcь\n")
	}

}

func getInput() (int, bool) {
	fmt.Println("Пожалуйста, введите целое число для вычисления результата или введите слово 'стоп' для завершения программы")
	var s string
	_, err := fmt.Scan(&s)
	if err != nil {
		panic(err)
	}
	if strings.ToLower(s) == "стоп" {
		fmt.Println("Завершаю работу")
		return 0, true
	}
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err, "Пожалуйста, введите корректное число")
	}
	return i, false
}

func square(squareCh chan int, multiplyCh chan int, wg *sync.WaitGroup) {
	num := <-squareCh
	fmt.Println("square: квадрат числа = ", num*num)
	defer wg.Done()
	multiplyCh <- num * num
}

func multiplyByTwo(multiplyCh chan int, wg *sync.WaitGroup) {
	num := <-multiplyCh
	fmt.Println("multiplyBy2: квадрат х 2 = ", num*2)
	defer wg.Done()
}
