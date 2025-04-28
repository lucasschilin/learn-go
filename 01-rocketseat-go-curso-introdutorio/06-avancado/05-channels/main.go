package main

import (
	"fmt"
	"time"
)

func producer(c1 chan int, c2 chan string) {
	for i := 1; i <= 5; i++ {
		c1 <- i * 10
	}

	c2 <- "Lucas Schlindwein"

	close(c1)
	close(c2)
}

func consumer1(c1 chan int) {
	for val := range c1 {
		fmt.Println(val)
	}
}

func consumer2(c2 chan string) {
	fmt.Println("Nome:", <-c2)
}

func main() {
	var channel1 chan int = make(chan int)
	channel2 := make(chan string)
	go producer(channel1, channel2)
	go consumer1(channel1)
	go consumer2(channel2)

	time.Sleep(1 * time.Second)

}
