package main

import (
	"fmt"
	"time"
)

func main() {
	go routine_A()
	routine_B()
}

func routine_A() {
	for i := 0; i < 10; i++ {
		fmt.Printf("1")
		time.Sleep(time.Second)
	}
}

func routine_B() {
	for i := 0; i < 10; i++ {
		fmt.Printf("2")
		time.Sleep(time.Second)
	}
}
