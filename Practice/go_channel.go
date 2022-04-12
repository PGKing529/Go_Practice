package main

// func main() {
// 	go run()
// 	// time.Sleep(time.Second)
// 	for i := 0; i < 100; i++ {
// 		fmt.Println(i)
// 	}
// }

// func run() {
// 	fmt.Println("出动")
// }

import (
	"fmt"
)

// func main() {
// 	var wg sync.WaitGroup
// 	wg.Add(2)
// 	go run_A(&wg)
// 	go run_B(&wg)
// 	wg.Wait()
// }

// func run_A(wg *sync.WaitGroup) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("A")
// 		time.Sleep(100)
// 	}
// 	wg.Done()
// }

// func run_B(wg *sync.WaitGroup) {
// 	for i := 0; i < 10; i++ {
// 		fmt.Println("B")
// 		time.Sleep(100)
// 	}
// 	wg.Done()
// }

// func main() {
// 	c1 := make(chan int)

// 	go func() {
// 		c1 <- 1
// 	}()

// 	fmt.Println(<-c1)
// }

func main() {
	c1 := make(chan int, 5)
	go func() {
		for i := 0; i < 10; i++ {
			c1 <- i
		}
	}()

	for i := 0; i < 10; i++ {
		fmt.Println(<-c1)
	}
}
