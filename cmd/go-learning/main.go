package main

import (
	"fmt"

	"github.com/NicholasMantovani/go-learning/internal/pkg/errors"
	"github.com/NicholasMantovani/go-learning/internal/pkg/functions"
	"github.com/NicholasMantovani/go-learning/internal/pkg/interfaces"
	"github.com/NicholasMantovani/go-learning/internal/pkg/structs"
	"github.com/NicholasMantovani/go-learning/internal/pkg/variables"
)

func main() {
	fmt.Println("\n---------------VARIABLES-----------")
	variables.ExecuteVariables()
	fmt.Println("\n---------------FUNCTIONS-----------")
	functions.ExecuteFunction()
	fmt.Println("\n---------------STRUCTS-----------")
	structs.ExecuteStructs()
	fmt.Println("\n---------------INTERFACES-----------")
	interfaces.ExecuteInterfaces()
	fmt.Println("\n---------------ERRORS-----------")
	errors.ExecuteErrors()
}
