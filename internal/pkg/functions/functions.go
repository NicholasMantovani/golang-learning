package functions

import (
	"errors"
	"fmt"
	"strconv"
)

func ExecuteFunction() {
	if sum(1, 2) == sumOnlyOneType(1, 2) {
		fmt.Println("\ncorrect")
	}

	callBackExample(2, logAndReturnCallback)
	copyVariableByValueExplained()
	copyVariableByValueCorrected()
	callFunctionWithMultipleReturnValue()

	a, b := namedReturnValues()

	if c, d := notNamedReturnValues(); a == c && b == d {
		fmt.Println("\nIts correct")
	}

	fmt.Println(divide(10, 0))

}

func sum(x int, y int) int {
	return x + y
}

func sumOnlyOneType(x, y int) int {
	return x + y
}

func callBackExample(x int, callback func(int) int) int {

	result := x * 2

	fmt.Println("This is executed before")
	return callback(result)
}

func logAndReturnCallback(a int) int {
	fmt.Println("This is executed after " + strconv.Itoa(a))
	return a
}

// Go copy variable by value, it doesn't copy its pointer to the memory

func copyVariableByValueExplained() {
	x := 5
	increment(x)

	fmt.Println(x)
	// still prints 5,
	// because the increment function
	// received a copy of x
}

func increment(x int) {
	x++
}

func copyVariableByValueCorrected() {
	x := 5
	x = incrementCorrected(x)

	fmt.Println(x)
	// will print 6
}

func incrementCorrected(x int) int {
	a := x + 1
	return a
}

func callFunctionWithMultipleReturnValue() {
	a, b := multipleReturnValue(1)

	fmt.Printf("\nFirst value is a %T", a)
	fmt.Printf("\nSecond value is a %T", b)

	// Ignore some return

	c, _ := multipleReturnValue(1)
	fmt.Printf("\nFirst value is a %T and second value is ignored", c)
}

func multipleReturnValue(x int) (int, string) {
	return x + 1, "multiple returns"
}

// This code is the same as
func namedReturnValues() (length, width int) {

	// length and width are inizialized with their "zero" values

	return // automatically return length and width. THIS TYPE OF RETURN SHOULD ONLY BE USED IN SHORT FUNCTIONS
}

// This code here
func notNamedReturnValues() (int, int) {
	var x int
	var y int
	return x, y
}

// Early returns is cosidered better than multiple if/if-else
func divide(dividend, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("Can't divide by zero")
	}
	return dividend / divisor, nil
}
