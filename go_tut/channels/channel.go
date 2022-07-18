package channels

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	c <- link
}

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
	fmt.Println("close channel")
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

func waitGroup() {
	var wg sync.WaitGroup
	chanel := make(chan int)

	// set counter of waitGroup
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

	for i := 0; i < 10; i++ {
		chanel <- i
	}

	close(chanel)

	fmt.Println("wait for goroutines to finish")
	// wait till wg.Done is called 2 times
	wg.Wait()
	fmt.Println("done waiting")
}

func Demo() {
	fmt.Println("-----channels-----")

	simple()

	buffered()

	rangeChan()

	closingChanel()

	waitGroup()

	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
		"http://www.netflex.com",
	}

	c := make(chan string)

	for _, link := range links {
		go checkLink(link, c)
	}

	// fmt.Println(<-c)
	// for i := 0; i < len(links); i++ {
	// 	fmt.Println(<-c)
	// }
	// for {
	// 	fmt.Println(<-c)
	// }

	// this will loop forever
	for l := range c {
		// time.Sleep(5 * time.Second)
		// fmt.Println(l)
		go func(l string) {
			time.Sleep(5 * time.Second)
			checkLink(l, c)
		}(l)
	}

}
