package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup
var mu sync.Mutex

func main() {
	a := [5]int{2, 4, 6, 8, 10}
	sum := 0
	for _, v := range a {
		wg.Add(1)

		go func(v int) {
			mu.Lock()
			defer wg.Done()

			sum += v * v
			mu.Unlock()
		}(v)
	}
	wg.Wait()
	fmt.Println(sum)

}
