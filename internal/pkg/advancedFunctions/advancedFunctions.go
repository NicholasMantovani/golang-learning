package advancedfunctions

import "fmt"

func ExecuteAdvancedFunction() {
	firstClassFunction()
	curryingFunctions()
	deferFunction()
	closuresFunction()
	anonymousFunction()
}

// aggregate applies the given math function to the first 3 inputs
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
	return arithmetic(arithmetic(a, b), c)
}

func sum(a, b int) int {
	return a + b
}

func multiply(a, b int) int {
	return a * b
}

// This means that the programming language treats a function like any other type,
// for example you can pass a function in a function as a parameter:
// this are useful for example as a click callback, or for some callback in http
func firstClassFunction() {
	fmt.Println("Executing sum on inputs", 2, 2, 3)
	fmt.Println("Result: ", aggregate(2, 2, 3, sum))

	fmt.Println("Executing multiply on inputs", 2, 2, 3)
	fmt.Println("Result: ", aggregate(2, 2, 3, multiply))

}

// This is a currying function in this case it maps a function that takes 2 input and transform it
// into one that accepts only one input
func selfMath(mathFunc func(int, int) int) func(int) int {
	return func(x int) int {
		return mathFunc(x, x)
	}
}

// a function that takes a function as input and returns a function
// this is mostly used in middleware for http handling
func curryingFunctions() {
	squareFunc := selfMath(multiply)
	doubleFunc := selfMath(sum)

	fmt.Printf("\nSquare function %T", squareFunc)
	fmt.Printf("\nDouble function %T", squareFunc)

	fmt.Println("\nSquare function result", squareFunc(5))
	// prints 25

	fmt.Println("Double function result", doubleFunc(5))
	// prints 10
}

func deferredFunction() {
	fmt.Println("This is a deferred function")
}

func fileSystemExample() string {
	fmt.Println("Im opening a file...")

	defer deferredFunction()

	fmt.Println("Im done with the file, the function will exit")

	return "Exited"
}

// Usally used when you must close connections, file ecc
func deferFunction() {
	fmt.Println("Defer example")
	fmt.Println("The function is: ", fileSystemExample())
}


func concatter() func(string) (string, int) {
	//This variable will be inizialised just once but the value will be updated every time
	doc := ""
	countTimesCalled := 0
	return func(word string) (string, int) {
		doc += word + " "
		countTimesCalled ++
		return doc, countTimesCalled
	}
}

func closuresFunction() {
	harryPotterAggregator := concatter()
	harryPotterAggregator("Mr.")
	harryPotterAggregator("and")
	harryPotterAggregator("Mrs.")
	harryPotterAggregator("Dursley")
	harryPotterAggregator("of")
	harryPotterAggregator("number")
	harryPotterAggregator("four,")
	harryPotterAggregator("Privet")

	fmt.Println(harryPotterAggregator("Drive"))
	// Mr. and Mrs. Dursley of number four, Privet Drive
}


// doMath accepts a function that converts one int into another
// and a slice of ints. It returns a slice of ints that have been
// converted by the passed in function.
func doMath(f func(int) int, nums []int) []int {
	var results []int
	for _, n := range nums {
		results = append(results, f(n))
	}
	return results
}

func anonymousFunction() {
	nums := []int{1, 2, 3, 4, 5}
	
    // Here we define an anonymous function that doubles an int
    // and pass it to doMath
	allNumsDoubled := doMath(func(x int) int {
	    return x + x
	}, nums)
	
	fmt.Println(allNumsDoubled)
    // prints:
    // [2 4 6 8 10]
}