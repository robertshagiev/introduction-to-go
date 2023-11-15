package main

import (
	"fmt"
	"time"
)

// 1
func main() {
	go goroutine()
	time.Sleep(time.Millisecond * 1300)
	fmt.Println("Finish")

}

func goroutine() {
	for i := 0; i < 5; i++ {
		time.Sleep(time.Millisecond * 5)
		fmt.Println(i)
	}
	return
}

//2

// func myGoroutine(stopCh chan struct{}) {
// 	for {
// 		select {
// 		case <-stopCh:
// 			fmt.Println("Stopping goroutine")
// 			return
// 		default:
// 			fmt.Println("Working...")
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }

// func main() {
// 	stopCh := make(chan struct{})

// 	go myGoroutine(stopCh)

// 	time.Sleep(3000 * time.Millisecond)

// 	close(stopCh)

// 	fmt.Println("Exiting main()")
// }

// 3
// func myGoroutine(ctx context.Context) {
// 	for {
// 		select {
// 		case <-ctx.Done():
// 			fmt.Println("Stopping goroutine")
// 			return
// 		default:
// 			fmt.Println("Working...")
// 			time.Sleep(500 * time.Millisecond)
// 		}
// 	}
// }
// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel()

// 	go myGoroutine(ctx)

// 	time.Sleep(3000 * time.Millisecond)

// 	fmt.Println("Exiting main()")
// }

//4

// func myGoroutine() {
// 	for i := 0; i < 5; i++ {
// 		time.Sleep(500 * time.Millisecond)
// 		fmt.Println(i)
// 	}

// 	runtime.Goexit()
// }

// func main() {
// 	go myGoroutine()

// 	time.Sleep(3000 * time.Millisecond)

// 	fmt.Println("Exiting main()")
// }
