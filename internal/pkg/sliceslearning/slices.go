package sliceslearning

import (
	"fmt"
	"reflect"
)

func ExecuteSlices() {
	declareArrays()
	names := getArraysOfNames()
	fmt.Println("Array of names", names)
	fmt.Println("Slices of names", getSliceOfNames(names))
	fmt.Println("Cast from array to slice", names[:])
	allTypeOfSlices()
	allTwoArray := []int{2, 2, 2, 2}
	fmt.Println("This is the array before modifing it in a separte function", allTwoArray)
	slicesArePassedAsReference(allTwoArray)
	fmt.Println("This is the array after modifing it in a separte function", allTwoArray)
	slicesWithMake()
	lenAndCapAreSafeToUseOnNilValues()
	fmt.Println("\nThis is a variadic function and it will sum 1,2,3 result: ", variadicSum(1, 2, 3))
	fmt.Printf("I will call a variadic function with the spread operator values: %v | result: %v", allTwoArray, variadicSum(allTwoArray...))
	appendSlices()
	executingCostByDay()
	slicesOfSlices()
	createMatrix(2, 2)
	iterateWithRange()
}

func declareArrays() {
	// array of 10 integers
	var arrayInts [10]int

	fmt.Println("Default array print", arrayInts)

	primes := [7]int{2, 3, 5, 7, 11, 13}
	fmt.Println("Default array with values print", primes)

}

func getArraysOfNames() [3]string {
	return [3]string{"Nicholas", "Mantovani", "Armani"}
}

func getSliceOfNames(names [3]string) []string {
	return names[0:1]
}

func allTypeOfSlices() {

	// arrayname[lowIndex:highIndex]
	// arrayname[lowIndex:]
	// arrayname[:highIndex]
	// arrayname[:]

	primes := [7]int{2, 3, 5, 7, 11, 13, 17}
	fmt.Println("Non sliced array", primes)
	fmt.Println("Slice from 1 to 4", primes[1:5])
	fmt.Println("Slice from 4 to end", primes[4:])
	fmt.Println("Slice from start to 6", primes[:7])
	fmt.Println("Slice from start to end", primes[:7])

}

func slicesArePassedAsReference(array []int) {
	array[0] = 1
}

func slicesWithMake() {
	// func make([]T, len, cap) []T
	mySliceWithCap := make([]int, 5, 10)

	// the capacity argument is usually omitted and defaults to the length
	mySliceWithoutCap := make([]int, 5)

	mySliceWithoutMake := []string{"Sium1", "Sium2", "Sium3"}

	fmt.Printf("\nSlice with make and cap. Len: %v | Cap: %v | Values: %v", len(mySliceWithCap), cap(mySliceWithCap), mySliceWithCap)

	fmt.Printf("\nSlice with make no cap. Len: %v | Cap: %v | Values: %v", len(mySliceWithoutCap), cap(mySliceWithoutCap), mySliceWithoutCap)

	fmt.Printf("\nSlice with no make no cap. Len: %v | Cap: %v | Values: %v", len(mySliceWithoutMake), cap(mySliceWithoutMake), mySliceWithoutMake)

}

func lenAndCapAreSafeToUseOnNilValues() {
	var test []string = nil
	if test == nil {
		fmt.Printf("\nThis is a nil slice. Len: %v | Cap: %v | Values: %v", len(test), cap(test), test)
	}
}

func variadicSum(nums ...int) int {
	sum := 0
	if nums != nil {
		for i := 0; i < len(nums); i++ {
			sum = sum + nums[i]
		}
	}
	return sum
}

func appendSlices() {
	firstSlice := []int{2, 2, 2, 2}
	secondSlice := make([]int, 5)

	fmt.Println("\nAppending slices. This is the first slice", firstSlice)
	fmt.Println("Appending slices. This is the second slice", secondSlice)

	secondSlice = append(secondSlice, 2)
	fmt.Println("Second slice after one append", secondSlice)
	secondSlice = append(secondSlice, 3, 4, 5)
	fmt.Println("Second slice after one multiple append", secondSlice)
	secondSlice = append(secondSlice, firstSlice...)
	fmt.Println("Second slice after the first slice append", secondSlice)

}

// BOOT.DEV PROBLEM
// FROM THIS:
//
//	[]cost{
//	    {0, 4.0},
//	    {1, 2.1},
//	    {1, 3.1},
//	    {5, 2.5},
//	}
//
// TO THIS:
//
//	[]float64{
//	    4.0,
//	    5.2,
//	    0.0,
//	    0.0,
//	    0.0,
//	    2.5,
//	}
type cost struct {
	day   int
	value float64
}

func getCostByDay(costs []cost) []float64 {
	costByDay := []float64{}
	for i := 0; i < len(costs); i++ {
		cost := costs[i]

		for cost.day >= len(costByDay) {
			// we must grow the array with some default value
			costByDay = append(costByDay, 0.0)
		}

		costByDay[cost.day] += cost.value
	}
	return costByDay
}

func executingCostByDay() {

	initialSlice := []cost{
		{day: 0, value: 4.0},
		{1, 2.1},
		{1, 3.1},
		{5, 2.5},
	}

	expectedSlice := []float64{
		4.0,
		5.2,
		0.0,
		0.0,
		0.0,
		2.5,
	}
	resultSlice := getCostByDay(initialSlice)
	fmt.Println("This is the inital array", initialSlice)
	fmt.Println("This is the expected array", expectedSlice)
	fmt.Println("This is the result of the function", resultSlice)

	if reflect.DeepEqual(expectedSlice, resultSlice) {
		fmt.Println("CORRECT!!!")
	}
}

func slicesOfSlices() {
	matrix := make([][]int, 0)
	fmt.Println("Created matrix", matrix)
	row := []int{1, 2, 3, 4, 5}
	fmt.Println("Created row", row)
	matrix = append(matrix, [][]int{row}...)
	row = []int{6, 7, 8, 9, 10}
	fmt.Println("Created second row", row)
	matrix = append(matrix, [][]int{row}...)
	fmt.Println("Matrix", matrix)
}

func createMatrix(rows, cols int) {
	matrix := make([][]int, 0)

	fmt.Printf("\nCreating matrix of %vx%v", rows, cols)

	for i := 0; i <= rows; i++ {
		row := make([]int, 0)
		for j := 0; j <= cols; j++ {
			row = append(row, i*j)
		}
		matrix = append(matrix, row)
	}
	fmt.Println("\nCreated matrix", matrix)
}

func iterateWithRange() {
	friends := []string{"Me", "Sium", "Fortnite"}

	fmt.Println("Iterating with forEach")

	for i, friend := range friends {
		fmt.Println(i, friend)
	}

	for _, friend := range friends {
		fmt.Println(friend)
	}
}
