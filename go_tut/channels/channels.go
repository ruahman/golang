package channels

import (
	"fmt"
	"net/http"
	"time"
)

// we use channels to let the main go routine know when the other go routines are done
// channels are the only way to communicate between go routines
func checkLink(link string, c chan string) {
	// as soon as we block on this line of code, the go runtime will start looking for other go routines to run
	_, err := http.Get(link)
	if err != nil {
		fmt.Println(link, "might be down!")
		// send link to the channel
		c <- link
		return
	}

	fmt.Println(link, "is up!")
	// send the link to the channel
	c <- link
}

// simple channel example.
// func simple() {
// 	// unbuffer chanel
// 	chanel := make(chan int)
//
// 	go func() {
// 		// finaly add a value to chanel
// 		chanel <- 312
// 	}()
//
// 	fmt.Println("wait for chanel to get a value")
// 	// wait till channel returns something
// 	x := <-chanel
// 	fmt.Println("got data from chanel: ", x)
// }

// func buffered() {
// 	// give it a slot of 1
// 	chanel := make(chan int, 1)
// 	// once a chanel is full it blocks
// 	chanel <- 111
// 	// doesnt block because it has a value
// 	fmt.Println(<-chanel)
//
// 	chanel2 := make(chan int, 2)
// 	chanel2 <- 222
// 	chanel2 <- 333
//
// 	fmt.Println(<-chanel2)
// 	fmt.Println(<-chanel2)
// }

// func rangeChan() {
// 	chanel := make(chan int)
//
// 	go func() {
// 		chanel <- 111
// 		chanel <- 222
// 		chanel <- 333
// 		chanel <- 444
// 		chanel <- 666
// 		close(chanel)
// 	}()
//
// 	// loop until chanel close
// 	// v, ok := <-ch, ok is false if chanel is closed
// 	// this is how the loop knows when to stop
// 	for val := range chanel {
// 		fmt.Println(val)
// 	}
// 	fmt.Println("close channel")
// }

// func closingChanel() {
// 	chanel := make(chan int)
// 	close(chanel)
// 	// causes panic
// 	// chanel <- 666
//
// 	// you can read from a close chanel
// 	val, ok := <-chanel
//
// 	fmt.Println(val, ok)
// }

// func waitGroup() {
// 	var wg sync.WaitGroup
// 	chanel := make(chan int)
//
// 	// set counter of waitGroup
// 	wg.Add(2)
//
// 	go func(c <-chan int) {
// 		for x := range c {
// 			fmt.Println("goroutine1: ", x)
// 		}
// 		wg.Done()
//
//
// 	go func(c <-chan int) {
// 		for x := range c {
// 			fmt.Println("goroutine2: ", x)
// 		}
// 		wg.Done()
// 	}(chanel)
//
// 	for i := 0; i < 10; i++ {
// 		chanel <- i
// 	}
//
// 	close(chanel)
//
// 	fmt.Println("wait for goroutines to finish")
// 	// wait till wg.Done is called 2 times
// 	wg.Wait()
// 	fmt.Println("done waiting")
// }

// func fibonacci(c, quit chan int) {
// 	x, y := 0, 1
// 	for {
// 		// select mulitle channels
// 		select {
// 		case c <- x:
// 			x, y = y, x+y
// 		case <-quit:
// 			fmt.Println("quit")
// 			return
// 		}
// 	}
// }

// the main function is a go routine
// the when the main routine finishes the program exits
// so its important to wait for the other go routines to finish
func Run() {
	fmt.Println("-----channels-----")

	links := []string{
		"http://www.google.com",
		"http://www.amazon.com",
		"http://www.facebook.com",
		"http://www.netflex.com",
	}

	// make a channel so that go routines can communicate with the main go routine
	c := make(chan string)

	for _, link := range links {
		// run this function in a brand new go routine
		// the go scheduler will decide which core to run this go routine on
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

	// simple()
	//
	// buffered()
	//
	// rangeChan()
	//
	// closingChanel()
	//
	// waitGroup()

	// con := make(chan int)
	// quit := make(chan int)
	// go func() {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(<-c)
	// 	}
	// 	quit <- 0
	// }()
	// fibonacci(con, quit)
}
