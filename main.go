package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(time.Microsecond)
		message <- "Hello word"

	}()

	fmt.Println(<-message)

}
