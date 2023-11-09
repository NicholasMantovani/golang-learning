package variables

import "fmt"

func ExecuteVariables() {
	howToAssingVariables()
	fmt.Println()
	variablesTypes()
	conditions()
	convertions()
}

func howToAssingVariables() {

	var firstWay int = 0

	secondWay := 0

	const thirdWay int = 0

	insaneWay, insaneWayTwo := 1, "ciao"

	fmt.Printf("first way, %T, %v", firstWay, firstWay)

	fmt.Printf("\nsecond way, %T, %v", secondWay, secondWay)

	fmt.Printf("\nthird way, %T, %v", thirdWay, thirdWay)

	fmt.Printf("\ninsane way, %T, %v", insaneWay, insaneWay)

	fmt.Printf("\ninsane way two, %T, %s", insaneWayTwo, insaneWayTwo)

}

func variablesTypes() {

	var allVariablesTypes string = `
	bool

	string
	
	int  int8  int16  int32  int64
	uint uint8 uint16 uint32 uint64 uintptr
	
	byte // alias for uint8
	
	rune // alias for int32
		 // represents a Unicode code point
	
	float32 float64
	
	complex64 complex128`

	fmt.Printf("all variables types: \n %s", allVariablesTypes)
}

func conditions() {

	val1 := 2.0

	val2 := 3.0

	if val1 == val2 {
		fmt.Println("Shoudln't be printed")
	} else {
		fmt.Println("Should be printed")
	}

	// insane condition

	email := "somemail@itsamail.com"

	if length := len(email); length < 1 {
		fmt.Println("Email is invalid")
	}

	// here length doesn't exist
}

func convertions() {
	x := 1

	y := float64(x)

	fmt.Printf("\nThe type is %T", y)
}
