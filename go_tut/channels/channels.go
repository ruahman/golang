package channels

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

// sum and send the sum to the channel
func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

// we use channels to let the main go routine know when the other go routines are done
// channels are the only way to communicate between go routines
// ch <- x, send val to channel
// x := <-ch, receive value from channel
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
func simple() {
	// unbuffer chanel
	chanel := make(chan int)

	go func() {
		// finaly add a value to chanel
		chanel <- 312
	}()

	fmt.Println("wait for chanel to get a value")
	// wait till channel returns something
	x := <-chanel
	fmt.Println("got data from chanel: ", x)
}

func buffered() {
	// give it a slot of 1
	chanel := make(chan int, 1)
	// once a chanel is full it blocks
	// if the channel is full it will block until the channel is empty
	chanel <- 111
	// doesnt block because it has a value
	// if channel is empty it will block until the channel has a value
	fmt.Println(<-chanel)

	// buffered channel with 2 slots
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
		// block till value is taken from loop
		chanel <- 222
		chanel <- 333
		chanel <- 444
		chanel <- 666

		// signal that the channel is closed
		// and there will be no more values
		// this is how the loop knows when to stop
		close(chanel)
	}()

	// loop until chanel close
	// v, ok := <-ch, ok is false if chanel is closed
	// this is how the loop knows when to stop
	for val := range chanel {
		fmt.Println(val)
	}
	fmt.Println("close channel")
}

// select statement lets a goroutine wait on multiple communication operations
func selectChans() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	quit := make(chan int)

	go func() {
		// loop until quit is closed
		for {
			// select which channel to read from
			select {
			case x := <-ch1:
				fmt.Println("ch1: ", x)
			case x := <-ch2:
				fmt.Println("ch2: ", x)
			case <-quit:
				fmt.Println("quit")
				return
			default:
				// if no other case is ready
				fmt.Println("default")
				time.Sleep(1 * time.Second)
			}
		}
	}()

	ch1 <- 1
	ch2 <- 2
	ch2 <- 3
	ch1 <- 4
	quit <- 0
}

type SafeCounter struct {
	v int
	// this makes sure that only one go routine can access the counter at a time
	mux sync.Mutex
}

func (c *SafeCounter) inc() {
	// lock it so that only one go routine can access the counter at a time
	c.mux.Lock()
	fmt.Println("increaseCounter: ", c.v)
	c.v++
	// unlock it so that other go routines can access the counter
	c.mux.Unlock()
}

func (c *SafeCounter) read() int {
	// lock it so that only one go routine can access the counter at a time
	c.mux.Lock()
	// unlock it so that other go routines can access the counter
	defer c.mux.Unlock()
	fmt.Println("readCounter: ", c.v)
	return c.v
}

func syncMutex() {
	c := SafeCounter{}
	for i := 0; i < 1000; i++ {
		go func() {
			c.inc()
		}()
	}

	time.Sleep(1 * time.Second)
	fmt.Println(c.read())
}

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

func Run() {
	fmt.Println("-----channels-----")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// wait till both channels have a value
	x, y := <-c, <-c
	fmt.Println(x, y, x+y)

	simple()
	buffered()
	rangeChan()
	selectChans()
	syncMutex()

	message := make(chan string)
	go func() { message <- "ping" }()
	msg := <-message
	fmt.Println(msg)

	// buffered channel, by default channels are unbuffered
	// they will only accept a value if there is a corresponding receive for the value
	// buffered channels can accept a limited number of values without a corresponding receive
	message2 := make(chan string, 2)
	message2 <- "buffered"
	message2 <- "channel"
	fmt.Println(<-message2)
	fmt.Println(<-message2)

	// blocking receive
	worker := func(done chan bool) {
		fmt.Println("working...")
		time.Sleep(1 * time.Second)
		done <- true
	}

	done := make(chan bool, 1)
	go worker(done)
	// this will block until we receive a value from the channel
	<-done
	fmt.Println("we are done")

	// select
	// select statement lets a goroutine wait on multiple communication operations
	// A select blocks until one of its cases can run, then it executes that case.

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		time.Sleep(1 * time.Second)
		c1 <- "one"
	}()
	go func() {
		time.Sleep(2 * time.Second)
		c2 <- "two"
	}()

	// keep on looping until we have received all messages,
	// waitgroup can also be used to wait for multiple go routines to finish
	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
	fmt.Println("we have received all messages")

	// closing channels
	// closing a channel indicates that no more values will be sent on it
	// this can be useful to communicate completion to the channel's receivers

	jobs := make(chan int, 5)
	done = make(chan bool)
	go func() {
		for {
			j, more := <-jobs
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	fmt.Println("waiting for jobs to finish")
	<-done
	fmt.Println("all jobs are done")

	// range over channels
	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue)
	for elem := range queue {
		fmt.Println("range: ", elem)
	}

	// timers, are for when you want to do something once in the future but
	// you want to stop it if you change your mind
	timer1 := time.NewTimer(2 * time.Second)

	// timer comes with a channel that will be notified when the timer is done
	<-timer1.C
	fmt.Println("timer1 done")

	// you can also stop the timer before it fires
	timer2 := time.NewTimer(3 * time.Second)
	go func() {
		<-timer2.C
		fmt.Println("timer2 done")
	}()
	// stop it before it fires
	stop2 := timer2.Stop()

	if stop2 {
		fmt.Println("timer2 stopped")
	}

	// tickers are for when you want to do something repeatedly at regular intervals
	ticker := time.NewTicker(500 * time.Millisecond)
	done = make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				fmt.Println("tick at", t)
			}
		}
	}()
	time.Sleep(1600 * time.Millisecond)
	ticker.Stop()
	// tell go routine to stop
	done <- true
	fmt.Println("ticker stopped")
}

// the main function is a go routine
// the when the main routine finishes the program exits
// so its important to wait for the other go routines to finish
func RunLinks() {
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
}

// jobs is a channel that will receive work
// results is a channel that will send the results of the work
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		time.Sleep(1 * time.Second)
		results <- j * 2
	}
}

func RunWorkerPool() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)

	// create the workers, three in this case
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for a := 1; a <= numJobs; a++ {
		<-results
	}
	fmt.Println("got all results")
}
