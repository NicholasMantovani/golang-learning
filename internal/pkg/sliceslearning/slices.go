package sliceslearning

import "fmt"

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
