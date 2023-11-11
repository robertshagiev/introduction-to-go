package main

import (
	"fmt"
	"sync"
)

func square(in <-chan int, out chan<- int) { //функция для вычисления квадратов чисел
	for val := range in { //интерация по каналу in, на каждой итерации текущий поток блокируется до следующего значения в канале
		square := val * val
		out <- square
	}
	close(out) // закрытие канала, новых значений поступать не будет. Запись в закрытый канал вызывавет панику
	//закрываем для того, чтобы цикл по каналу прекратил считать значения
}

func read(out <-chan int, wg *sync.WaitGroup) { //считываение данных из канала out
	defer wg.Done()
	for val := range out { //интерация по каналу out и вывод его переменных
		fmt.Println(val)
	}
}

func main() {
	var wg sync.WaitGroup

	in := make(chan int, 3)
	out := make(chan int, 3)

	a := [3]int{1, 2, 3}

	go square(in, out) // горутниа которая возводит значения в квадрат канала in и передает их в канал out
	wg.Add(1)
	go read(out, &wg) // горунта которая выводит значения из канала out

	for _, val := range a { // цикл который передает все значения в канал in
		in <- val
	}
	close(in)
	wg.Wait()
}
