package main

import (
	"fmt"
	"time"
)

func even(number int) {
	fmt.Println("É um valor par:", number)
}

func odd(number int) {
	fmt.Println("É um valor impar:", number)
}

func main() {
	ints := []int{91, 32, 123, 53, 346, 234}
	number := make(chan int)

	go func() {
		for _, i := range ints {
			time.Sleep(time.Second)
			number <- i
		}
	}()

	go func() {
		for i := range number {
			if i%2 == 0 {
				even(i)
			} else {
				odd(i)
			}
		}
	}()

	for {
	}
}
