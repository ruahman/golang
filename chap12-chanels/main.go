package main

import (
	"fmt"
	"sync"
)

func simple() {
	// unbuffer chanel
	chanel := make(chan int)

	go func() {
		// finaly add a value to chanel
		chanel <- 312
	}()

	fmt.Println("wait for chanel to get a value")
	x := <-chanel
	fmt.Println("got data from chanel: ", x)
}

func buffered() {
	// give it a slot of 1
	chanel := make(chan int, 1)
	// once a chanel is full it blocks
	chanel <- 111
	// doesnt block because it has a value
	fmt.Println(<-chanel)

	chanel2 := make(chan int, 2)
	chanel2 <- 222
	chanel2 <- 333

	fmt.Println(<-chanel2)
	fmt.Println(<-chanel2)

}

func brokenFor() {
	chanel := make(chan int)
	go func() {
		for {
			fmt.Println(<-chanel)
		}
	}()

	chanel <- 312
	chanel <- 773
	chanel <- 878

	fmt.Println("done")
}

func rangeChan() {
	chanel := make(chan int)

	go func() {
		chanel <- 111
		chanel <- 222
		chanel <- 333
		chanel <- 444
		chanel <- 666
		close(chanel)
	}()

	// loop until chanel close
	for val := range chanel {
		fmt.Println(val)
	}
}

func closingChanel() {
	chanel := make(chan int)
	close(chanel)
	// causes panic
	// chanel <- 666

	// you can read from a close chanel
	val, ok := <-chanel

	fmt.Println(val, ok)

}

func broadCast() {
	var wg sync.WaitGroup
	chanel := make(chan int)

	wg.Add(2)

	go func(c <-chan int) {
		for x := range c {
			fmt.Println("goroutine1: ", x)
		}
		wg.Done()
	}(chanel)

	go func(c <-chan int) {
		for x := range c {
			fmt.Println("goroutine2: ", x)
		}
		wg.Done()
	}(chanel)

	for i := 0; i < 1000; i++ {
		chanel <- i
	}

	close(chanel)

	fmt.Println("wait for goroutines to finish")
	wg.Wait()
	fmt.Println("done waiting")
}

func main() {
	// simple()
	// buffered()
	// brokenFor()
	// rangeChan()
	// closingChanel()
	broadCast()
}
