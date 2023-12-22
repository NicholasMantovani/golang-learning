package pointers

import "fmt"

// WHY USE POINTERS: https://google.github.io/styleguide/go/decisions#receiver-type
// The main thing is if you want to manipulate the data inside a function because by default value in function are passed by copy and not by refrence (unlike java objects)
// The second reason is performance, if you have a large struct (maybe with a file inside) creating a copy of that struct is very heavy so you should use pointers.

func ExecutePointers() {
	pointerSyntax()
	pointer()
	quiz1()
	quiz2()
	panicCode()
	testMethod()
}

func pointerSyntax() {
	// This is not so used
	var a *int

	fmt.Printf("This is a pointer: value %v | Type %T", a, a)

	// This is way more used
	myString := "Test"
	myStringPointer := &myString

	fmt.Printf("\nThis is a pointer: value %v | Type %T", myStringPointer, myStringPointer)

	fmt.Printf("\nThis is how to get a value of a pointer: %v", *myStringPointer) // Read myString through the pointer

	*myStringPointer = "new test" //Set myString thorough the pointer

	fmt.Printf("\nThe same syntax can be used to modify the value of the pointer: %v", *myStringPointer)

	fmt.Printf("\nModifing the pointer values change the value of the underlying variable: %v", myString)
}

type Test struct {
	a, b int
}

func pointer() {
	t := Test{a: 1, b: 2}
	fmt.Printf("\nT before calling the method %v", t)
	modifyStruct(&t)
	fmt.Printf("\nT after calling the method %v", t)
}

func modifyStruct(t *Test) {
	t.a *= 10
	t.b *= 10
}

func quiz1() {
	// This is the quiz code
	var x int = 50
	var y *int = &x
	*y = 100

	// Question: What is the value of *y after the code on the left executes?
	fmt.Println("\nThe resul of quiz 1 is ", *y)
}

func quiz2() {
	// This is the quiz code
	var x int = 50
	var y *int = &x
	*y = 100

	// Question: What is the value of x after the code on the left executes?
	fmt.Println("The resul of quiz 2 is ", x)
}

func panicCode() {
	var a *int

	fmt.Printf("This is a pointer zero value: value %v | Type %T\n", a, a)
	// *a = 10 THIS LINE WILL PANIC

	if a != nil {
		fmt.Println("This will never be executed")
	}

	// *a = 10 THIS LINE WILL PANIC

	b := 10
	a = &b // After this line it will no longer panic

	fmt.Printf("This is a pointer after associating a variable to it: value %v | Type %T\n", a, a)
}

// POINTER WITH METHOD

type Car struct {
	model string
}

// This does not require a pointer
func (c Car) getModel() string {
	return c.model
}

// This require a pointer
func (c *Car) setModel(model string) {
	c.model = model
}

func (c Car) setModelWithoutPointer(model string) {
	c.model = model
}

func testMethod() {
	car := Car{"Toyota"}
	fmt.Println("This is the car model", car.getModel())
	car.setModelWithoutPointer("Hyundai")
	fmt.Println("This is the new car model setted without using the pointer", car.getModel())

	car.setModel("Hyundai")
	fmt.Println("This is the new car model setted with the pointer", car.getModel())
}
