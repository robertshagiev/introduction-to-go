package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		time.Sleep(time.Microsecond) // sleep 1 microsecond
		message <- "Hello word"

	}()

	fmt.Println(<-message)

}
