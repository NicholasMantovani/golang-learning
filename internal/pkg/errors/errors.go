package errors

import (
	"fmt"
	"strconv"
)

// THIS IS THE BUILT IN ERROR INTERFACE
// type error interface {
//   Error() string
//}

func ExecuteErrors() {
	executeFunctionWithError("42b")
	executeFunctionWithError("42")
	executeFunctionWithErrorHandled("52b")
	executeFunctionWithErrorHandled("52b")
}

func executeFunctionWithError(value string) {
	val, err := functionWithPossibileError(value)
	if err != nil {
		fmt.Println("Couldn't convert:", err)
		return
	}
	fmt.Println("Converted:", val)
}

func executeFunctionWithErrorHandled(value string) {
	val := functionWithPossibileErrorHandled(value)
	fmt.Println("Converted: ", val)
}


func functionWithPossibileError(value string) (int, error) {
	return strconv.Atoi(value)
}

func functionWithPossibileErrorHandled(value string) int {
	converted, err := strconv.Atoi(value)
	if err != nil {
		fmt.Println("Couldn't convert:", err)
		return 0
	}
	return converted
}