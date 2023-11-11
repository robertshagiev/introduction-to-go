package main

import (
	"fmt"
	"sync"
)

// функция для вычисления квадратов чисел

func square(in chan int, out chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for val := range in { //интерация по каналу in, на каждой итерации текущий поток блокируется до следующего значения в канале
		square := val * val
		out <- square
	}
	close(out) // закрытие канала, новых значений поступать не будет. Запись в закрытый канал вызывавет панику

}

func read(out chan int, wg *sync.WaitGroup) { //считываение данных из канала out

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

	wg.Add(3)

	go func(in chan int, wg *sync.WaitGroup) {
		defer wg.Done()
		for _, val := range a { // цикл который передает все значения в канал in
			in <- val
		}
	}(in, &wg)

	go square(in, out, &wg) // горутниа которая возводит значения в квадрат канала in и передает их в канал out

	go read(out, &wg) // горунта которая выводит значения из канала out

	close(in)
	wg.Wait()

}
