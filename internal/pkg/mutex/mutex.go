package mutex

import (
	"fmt"
	"sync"
)

// Mutexes allow us to lock access to data. This ensures that we can control which goroutines can access certain data at which time.

// Go's standard library provides a built-in implementation of a mutex with the sync.Mutex type and its two methods:

// .Lock()
// .Unlock()
// We can protect a block of code by surrounding it with a call to Lock and Unlock as shown on the protected() method below.
// It's good practice to structure the protected code within a function so that defer can be used to ensure that we never forget to unlock the mutex.

func protected(mu sync.Mutex) {
	mu.Lock()
	defer mu.Unlock()
	// the rest of the function is protected
	// any other calls to `mu.Lock()` will block
}

// It stands for mutual exclusion so only one go routine can access that variable

// MAPS ARE NOT THREAD-SAFE
// Maps are not safe for concurrent use! If you have multiple goroutines accessing the same map, and at least one of them is writing to the map, you must lock your maps with a mutex.
// If you don't lock them go will PANIC

func ExecuteMutext() {
	myFirstMutex()
	rwMutex()
}

type safeCounter struct {
	counts map[string]int
	mu     *sync.Mutex
}

type emailCounter struct {
	email string
	count int
}

func (sc safeCounter) inc(key string) {
	sc.mu.Lock() // If these are commented i get this problem: fatal error: concurrent map writes
	defer sc.mu.Unlock()
	sc.counts[key]++
}

func myFirstMutex() {
	emails := []emailCounter{{email: "test", count: 10}, {email: "test", count: 1}, {email: "test", count: 4}, {email: "test", count: 6}}
	sc := safeCounter{
		counts: make(map[string]int),
		mu:     &sync.Mutex{},
	}

	var wg sync.WaitGroup

	for _, email := range emails {
		for i := 0; i < email.count; i++ {
			wg.Add(1)
			go func(email emailCounter) {
				sc.inc(email.email)
				wg.Done()
			}(email)
		}
	}

	wg.Wait()
	fmt.Println("The safeCounter is", sc)

}

//The sync.RWMutex can help with performance if we have a read-intensive process.
//  Many goroutines can safely read from the map at the same time (multiple Rlock() calls can happen simultaneously).
// However, only one goroutine can hold a Lock() and all RLock()'s will also be excluded.

func rwMutex() {
	m := map[int]int{}

	mu := &sync.RWMutex{}

	var wg sync.WaitGroup

	wg.Add(1)
	go writeLoop(m, mu, &wg)

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readLoop(m, mu, &wg, i)
	}

	wg.Wait()
	fmt.Println("Ended wait")
	// stop program from exiting, must be killed
	// block := make(chan struct{})
	// <-block
}

func writeLoop(m map[int]int, mu *sync.RWMutex, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 10; i++ {
		mu.Lock()
		m[i] = i
		mu.Unlock()
	}

}

func readLoop(m map[int]int, mu *sync.RWMutex, wg *sync.WaitGroup, time int) {
	defer wg.Done()
	mu.RLock()
	for k, v := range m {
		fmt.Println("Time: ", time, k, "-", v)
	}
	mu.RUnlock()

}
