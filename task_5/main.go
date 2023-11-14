package main

import (
	"fmt"
	"time"
)

func main() {
	var n int
	ch := make(chan int)
	fmt.Println("Wtire number:")
	fmt.Scan(&n)
	go write(ch)
	go read(ch)

	<-time.After(time.Second * time.Duration(n))

	fmt.Println("time is over")

}

func write(ch chan int) {

	for {
		ch <- 5
		fmt.Println("Write in channal")
	}
}
func read(ch chan int) {

	for {
		<-ch
		fmt.Println("Read from channal")
	}
}
