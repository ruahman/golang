package goroutines

import (
	"fmt"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

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

	// however
}
