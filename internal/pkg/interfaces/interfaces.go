package interfaces

import (
	"fmt"
	"math"
)

// Interface are inherit implicitally, you dont need to explicit tell the compiler to implement an interface (like java)

// This is the interface
type shape interface {
	calculateArea() float64
	calculatePerimeter() float64
}

// this is the object that will implement the interface
type rectangle struct {
	width, height float64
}

// these are the methods of the object that are implementing the interface

func (rec rectangle) calculateArea() float64 {
	return rec.width * rec.height
}

func (rec rectangle) calculatePerimeter() float64 {
	return 2*rec.height + 2*rec.width
}

// this is another object that will implement the interface
type circle struct {
	radius float64
}

func (cir circle) calculateArea() float64 {
	return math.Pi * cir.radius * cir.radius
}

func (cir circle) calculatePerimeter() float64 {
	return 2*math.Pi - cir.radius
}

func ExecuteInterfaces() {
	recangle := rectangle{width: 15, height: 10}
	circle := circle{radius: 10}
	fmt.Print("\nPrinting the area and perimeter of the rectangle")
	printAreaAndPerimeterOfTheShapes(recangle)
	fmt.Print("\nPrinting the area and perimeter of the cicle")
	printAreaAndPerimeterOfTheShapes(circle)
	copier := awsFileCopier{}
	copier.Copy("test.txt", "awsText.txt")
	typeAssertion(recangle)
	typeAssertion(circle)
}

func printAreaAndPerimeterOfTheShapes(s shape) {
	fmt.Printf("\nThe area of the shape is: %v", s.calculateArea())
	fmt.Printf("\nThe perimeter of the shape is: %v", s.calculatePerimeter())
}

// THIS WILL NOT WORK
//func (s shape) printShapesAreaAndPerimeter() {
//	fmt.Printf("\nThe area of the shape is: %v", s.calculateArea())
//	fmt.Printf("\nThe perimeter of the shape is: %v", s.calculatePerimeter())
//}

// A Struct can implement multiple interface
type points interface {
	countAngles() int
}

type cube struct {
	side float64
}

func (cub cube) calculateArea() float64 {
	return cub.side * cub.side
}

func (cub cube) calculatePerimeter() float64 {
	return 4 * cub.side
}

func (cub cube) countAngles() int {
	return 4
}

// NAME INTERFACE ARGUMENTS

// not named
type Copier interface {
	copy(string, string) int
}

// NAMED
type copier interface {
	Copy(sourceFile string, destinationFile string) int
}

type awsFileCopier struct {
	s3Bucket string
}

func (a awsFileCopier) Copy(sourceFile string, destinationFile string) int {
	fmt.Printf("\nIm copyin this file: %s in this file: %s", sourceFile, destinationFile)
	return 1
}

func typeAssertion(s shape) {
	c, ok := s.(circle)
	fmt.Print("\n----------------------------------------")
	fmt.Printf("\nThe shape is: %v", s)
	if ok {
		fmt.Printf("\nC is: %v", c)
		fmt.Printf("\nThe shape is a circle and the radius is: %v", c.radius)
	} else {
		fmt.Printf("\nC is: %v", c)
		// This print 0 because its the default value of the radius
		fmt.Printf("\nThe shape is not a circle but the area is %v", c.calculateArea())
	}
	
}
