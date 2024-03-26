package waitgroups

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func worker(id int) {
	fmt.Println("worker", id, "processing job")
	time.Sleep(1 * time.Second)
	fmt.Println("worker", id, "done")
}

func Run() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(id int) {
			worker(id)
			wg.Done()
		}(i)
	}

	// wait till all workers are done
	wg.Wait()
	fmt.Println("all workers done")
}

func RunAtomicCounters() {
	// atomic counters provide low level atomic operations
	// only one goroutine can access the counter at a time
	// this is useful for managing state in concurrent programs
	var counter atomic.Uint64
	var wg sync.WaitGroup

	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < 1000; j++ {
				counter.Add(1)
			}
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("counter", counter.Load())
}

type Container struct {
	m        sync.Mutex
	counters map[string]int
}

func (c *Container) inc(name string) {
	// this makes sure that only one goroutine can access the datastucture at a time
	c.m.Lock()
	// you must remember to unlock the mutex to allow other goroutines to access the datastructure
	defer c.m.Unlock()
	c.counters[name]++
}

func RunMutex() {
	// mutexes provide a higher level of control over access to shared memory
	// if you need to share something more complex than a single value, you can use a mutex

	c := Container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	doIncrement := func(name string, num int) {
		for i := 0; i < num; i++ {
			c.inc(name)
		}
		wg.Done()
	}

	wg.Add(3)
	go doIncrement("a", 1000)
	go doIncrement("b", 1000)
	go doIncrement("a", 1000)

	wg.Wait()
	fmt.Println("counters", c.counters)
}
