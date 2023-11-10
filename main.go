package main

import (
	"fmt"
)

func main() {

	intCh := make(chan int)

	go func() {

		intCh <- 5 // отправка данных в канал
	}() // вызов горутины
	fmt.Println(<-intCh) // получение данных из канала

}
