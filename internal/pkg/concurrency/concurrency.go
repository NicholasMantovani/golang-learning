package concurrency

import (
	"fmt"
	"time"
)

// CONCURRENCY != PARALLELISM
// In this i'll leran concurrent vs synchronous code
// To build concurrent code you can utilize the GO keyword it basically spin up a goroutine that at a high level is spinning up the function in another thread
// If we use the GO keyword we can't wait for a response.
// More info: https://dave.cheney.net/2014/03/19/channel-axioms

func ExecuteConcurrency() {
	randomOrderSay()
	texioExample()
	// channelsDeadLock()
	channels()
	textioTest()
	bufferedChannels()
	closingChannels()
	rangeChannel()
	selectChannel()
	tickers()
	readandWriteOnlyChannel()
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

// This is not actually random
func randomOrderSay() {
	go say("world") // This execute the function in another thread.
	say("hello")
}

// I need to print the received message after the sent message, if the code is written like this it executes sync
func sendEmail(message string) {
	func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

// I need to print the received message after the sent message, if the code is written like this it executes sync
func sendEmailCorret(message string) {
	go func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email received: '%s'\n", message)
	}()
	fmt.Printf("Email sent: '%s'\n", message)
}

func testSync(message string) {
	sendEmail(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}

func testAsync(message string) {
	sendEmailCorret(message)
	time.Sleep(time.Millisecond * 500)
	fmt.Println("========================")
}
func texioExample() {
	fmt.Println("============SYNC===========")
	testSync("Hello there Stacy!")
	testSync("Hi there John!")
	fmt.Println("============ASYNC===========")
	testAsync("Hello there Stacy!")
	testAsync("Hi there John!")
}

// Channels are use to pass data from go routines
// This will produce a DEADLOCK!!!!
func channelsDeadLock() {
	channel := make(chan int)  // This is an int channel
	channel <- 69              // Add 69 to the channel //This are blocking operation so this if noone is reading the channel it will deadlock
	channelResult := <-channel // i pass the channel value to a variable  //This are blocking operation so this if noone is reading the channel it will deadlock

	fmt.Println("This is the channel result", channelResult)
}

// Channels are use to pass data from go routines
func channels() {
	channel := make(chan int)
	go sixtynine(channel)
	channelResult := <-channel
	fmt.Println("This is the channel result", channelResult)
}

func sixtynine(c chan int) {
	c <- 69
}

func textioTest() {
	numDBs := 2
	fmt.Println("Channels dbs: ", numDBs)
	dbChan := getDatabasesChannel(numDBs)
	waitForAllTheDbs(numDBs, dbChan)
	fmt.Println("All database waited with channels")
}

func waitForAllTheDbs(numDBs int, dbChan chan struct{}) {
	for i := 0; i < numDBs; i++ {
		<-dbChan // Block until a value is in the channel
	}
}

func getDatabasesChannel(numDBs int) chan struct{} {
	ch := make(chan struct{})
	go func() {
		for i := 0; i < numDBs; i++ {
			time.Sleep(time.Second)
			ch <- struct{}{}
			fmt.Printf("Database channel %v is online\n", i+1)
		}
	}()
	return ch
}

func bufferedChannels() {
	var numDBs = 3
	fmt.Println("Buffered Channels dbs: ", numDBs)

	ch := make(chan struct{}, numDBs)
	getDatabasesChannelBuffered(numDBs, ch)
	<-ch // The program goes to here only when the buffered channel is full
	fmt.Println("All database waited with buffered channels")
}

func getDatabasesChannelBuffered(numDBs int, ch chan struct{}) {
	for i := 0; i < numDBs; i++ {
		time.Sleep(time.Second)
		ch <- struct{}{}
		fmt.Printf("Database buffer %v is online\n", i+1)
	}
}

// Channels can be closed
// If you send on a closed channel the code will panic
func closingChannels() {
	numDBs := 4
	fmt.Println("Closing Channels dbs: ", numDBs)

	ch := make(chan struct{})

	go getDatabasesChannelClosed(numDBs, ch)

	for {
		v, ok := <-ch //If the channels is closed ok = false  // if i put _ insead of v the channel will not consume the value and it will be forever pending

		if !ok {
			break
		}
		fmt.Print(v)
	}

	fmt.Println("All database waited with closing channel")

}

func getDatabasesChannelClosed(numDBs int, ch chan struct{}) {
	for i := 0; i < numDBs; i++ {
		time.Sleep(time.Second)
		ch <- struct{}{}
		fmt.Printf("Database closed %v is online\n", i+1)
	}
	close(ch)
}

func rangeChannel() {
	fmt.Println("Fibonacci sequence")
	ch := make(chan int)
	go fibonacci(10, ch)

	for v := range ch {
		fmt.Println(v)
	}
}

func fibonacci(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
		time.Sleep(time.Millisecond * 10)
	}
	close(ch)
}

func selectChannel() {

	numCh, stringCh := getNumsOrStrings()

	for {
		select {
		case n, ok := <-numCh:
			if !ok {
				return
			}
			fmt.Println("Num value", n)

		case s, ok := <-stringCh:
			if !ok {
				return
			}
			fmt.Println("String value", s)
			//default: // This is triggered if when the select is executed and there is no case so no values in the channel
		}
	}
}

func getNumsOrStrings() (chan int, chan string) {

	numCh := make(chan int)
	stringCh := make(chan string)
	go func() {
		for i := 0; i < 10; i++ {
			done := make(chan struct{})
			go func() {
				numCh <- i
				done <- struct{}{}
			}()
			go func() {
				stringCh <- fmt.Sprint(i)
				done <- struct{}{}
			}()
			<-done
			<-done
		}
		close(numCh)
		close(stringCh)
	}()

	return numCh, stringCh
}

// time.Tick() is a standard library function that returns a channel that sends a value on a given interval.
// time.After() sends a value once after the duration has passed.
// time.Sleep() blocks the current goroutine for the specified amount of time.
func tickers() {
	tickCh := time.Tick(100 * time.Millisecond)
	count := 0

	for v := range tickCh {
		if count > 10 {
			break
		}
		count++
		fmt.Println("Tick", v)
	}
}

func readandWriteOnlyChannel() {
	ch := make(chan int)

	func(ch <-chan int) { // This is not a deadlock since i'm casting the channel to a read only channel
		fmt.Println("This is a read only channel")
	}(ch)

	func(ch chan<- int) { // This is not a deadlock since i'm casting the channel to a write only channel
		fmt.Println("This is a write only channel")
	}(ch)
}

// A SEND TO A NIL CHANNEL BLOCKS FOREVER
// var c chan string // c is nil
// c <- "let's get started" // blocks

// A RECEIVE FROM A NIL CHANNEL BLOCKS FOREVER
// var c chan string // c is nil
// fmt.Println(<-c) // blocks

// A SEND TO A CLOSED CHANNEL PANICS
// var c = make(chan int, 100)
// close(c)
// c <- 1 // panic: send on closed channel

// A RECEIVE FROM A CLOSED CHANNEL RETURNS THE ZERO VALUE IMMEDIATELY
// var c = make(chan int, 100)
// close(c)
// fmt.Println(<-c) // 0
