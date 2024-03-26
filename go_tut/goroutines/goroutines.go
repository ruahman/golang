package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

func Demo() {
	fmt.Println("----- goroutines -----")

	urls := []string{
		"https://www.udemy.com/",
		"https://www.google.com/",
		"https://www.facebook.com/",
		"https://www.amazon.com/",
		"https://www.twitter.com/",
	}

	for _, url := range urls {
		// pushes exection on a seperate thread,
		// so that it doesn't block the current thread
		// go thread is a light weight thread
		// they call also be called green threads
		// the run on the same address space so access to shared memory should be synchronized

		// add waitgroup counter to 1
		wg.Add(1)
		go func(x string) {
			time.Sleep(1 * time.Second)
			fmt.Println(x)
			// brings down waitgroup counter
			wg.Done()
		}(url)
	}

	// wait till waitgroup counter is 0
	wg.Wait()
	fmt.Println("end main thread")

	// goroutines are lightweight threads that are managed by go runtime
	go say("world")
	say("hello")

	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)

	// wait till both channels have a value
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)

	// run this synchronously
	f("sychronous")

	// run this as a goroutine
	go f("goroutine")

	go func(msg string) {
		fmt.Println(msg)
	}("going")
	time.Sleep(1 * time.Second)
	fmt.Println("done")
}

func RunStatefullGoroutines() {
	// todo
}
