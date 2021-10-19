package main

import (
	"fmt"
)

func simpleSelect() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	exitCh := make(chan interface{})

	go func() {
		ch1 <- 444
		ch1 <- 333
		ch1 <- 444
		ch1 <- 444
		ch1 <- 444
		ch1 <- 444
		ch1 <- 444
		ch2 <- 999
		ch2 <- 999
		ch2 <- 999
		ch2 <- 999
		ch2 <- 999
		ch2 <- 999
		exitCh <- true
	}()

	for {
		// wait till get someting from a chanel
		select {
		case i := <-ch1:
			fmt.Println("ch1: ", i)
		case i := <-ch2:

			fmt.Println("ch2: ", i)
		case <-exitCh:
			return
		}
	}
}

func main() {
	simpleSelect()

}
