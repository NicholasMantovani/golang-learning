package loops

import "fmt"

func ExecuteLoops() {
	basicForLoop()
	// foreverRunningLoop()
	// foreverRunningLoopTwo()
	maxMessages := calculateMaxMessages(5.0, 0.1)
	fmt.Println("The maxMessage is", maxMessages)
	whileLoop()
	continueForLoop()
	breakForLoop()
	printPrimes(100)
}

func basicForLoop() {
	for i := 0; i < 2; i++ {
		fmt.Println(i)
	}
}

// ANY PART OF THE LOOP CAN BE OMITTED
func foreverRunningLoop() {
	for i := 0; ; i++ {
		fmt.Println(i)
	}
}

func foreverRunningLoopTwo() {
	for {
		fmt.Println("Im runnig")
	}
}

func calculateMaxMessages(budget float64, costPerMessage float64) int {
	totalCost := 0.0
	for i := 0; ; i++ {
		totalCost += costPerMessage
		if totalCost >= budget {
			return i
		}
	}
}

func whileLoop() {
	// Count to 10
	count := 0

	for count < 10 {
		fmt.Println(count)
		count++
	}
}


// CONTINUE
// THIS WILL SKIP THE CURRENT LOOP AND GO TO THE NEXT ITERATION
// THIS WILL PRINT ALL NON EVEN NUMBER
func continueForLoop() {
	fmt.Println("Continue for loop")
	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println(i)
	}
}

// BREAK
// THIS WILL STOP THE LOOP
// THIS WILL PRINT ALL NUMBER UNTIL 4
func breakForLoop() {
	fmt.Println("Break for loop")
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}
}


// PRINT PRIME NUMBERS

func printPrimes(max int) {
	for n := 2; n < max+1; n++ {
		if n == 2 {
			fmt.Println(n)
			continue
		}
		if n % 2 == 0 {
			continue
		}
		isPrime := true
		for i := 3; i * i < n + 1; i++ {
			if n % i == 0 {
				isPrime = false
				break
			}
		}
		if !isPrime {
			continue
		} else {
			fmt.Println(n)
		}
	}
}