package main

import (
	"fmt"

	advancedfunctions "github.com/NicholasMantovani/go-learning/internal/pkg/advancedFunctions"
	cesarcipher "github.com/NicholasMantovani/go-learning/internal/pkg/cesarCipher"
	"github.com/NicholasMantovani/go-learning/internal/pkg/errors"
	"github.com/NicholasMantovani/go-learning/internal/pkg/functions"
	"github.com/NicholasMantovani/go-learning/internal/pkg/interfaces"
	"github.com/NicholasMantovani/go-learning/internal/pkg/loops"
	"github.com/NicholasMantovani/go-learning/internal/pkg/mapsLearning"
	"github.com/NicholasMantovani/go-learning/internal/pkg/sliceslearning"
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
	fmt.Println("\n---------------LOOPS-----------")
	loops.ExecuteLoops()
	fmt.Println("\n---------------SLICES-----------")
	sliceslearning.ExecuteSlices()
	fmt.Println("\n---------------MAPS-------------")
	mapslearning.ExecuteMaps()
	fmt.Println("\n---------------ADVANCED_FUNCTIONS------")
	advancedfunctions.ExecuteAdvancedFunction()




	fmt.Println("\n---------------Test---------------")
	cesarcipher.ExecuteCesarCipher()
}
