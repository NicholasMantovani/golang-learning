package structs

import (
	"fmt"
	"time"
)

// Defining my first struct
type Car struct {
	Model    string
	Power    int
	CarMotor Motor
}

type Motor struct {
	Hp, Cylinders, Liters int
}

// This is an anonymus struct
type Book struct {
	Title   string
	PubYear int
	Owner   struct {
		Name string
		Age  int
	}
}

func ExecuteStructs() {

	describeCar(Car{Model: "Alfa Romeo Giulietta", Power: 120})

	describeCarAndMotor()

	inizializeWithoutValues()

	inizializeSomeValues()

	anonymusStruct()

	embeddedStruct()

	executeMethod()
}

func describeCar(car Car) {
	fmt.Printf("\nCar model: %s | Car hp: %v", car.Model, car.Power)
}

func describeCarAndMotor() {
	car := Car{Model: "Alfa Romeo Giulia", Power: 180, CarMotor: Motor{Hp: 180, Cylinders: 4, Liters: 2.0}}
	fmt.Printf("\nMy car is: %v", car)
	describeCar(car)
	describeMotor(car.CarMotor)
}

func describeMotor(motor Motor) {
	fmt.Printf("\nMotor hp: %v | Motor cylinders: %v | Motor liters: %v", motor.Hp, motor.Cylinders, motor.Liters)
}

// if nothing is passed the values are at its defaul zero value
func inizializeWithoutValues() {
	car := Car{}
	fmt.Printf("\nA generic car is: %v", car)

	car.CarMotor.Hp = 180

	fmt.Printf("\nNow it has some horsepower: %v", car)
}

func inizializeSomeValues() {
	car := Car{Model: "Alfa Romeo Stelvio", CarMotor: Motor{Hp: 170}}
	fmt.Printf("\nWe know something about this car: %v", car)
}

// These type of struct are used only if you need the struct in that specific part of code and you will not plan to reuse it in other part of code
func anonymusStruct() {
	dog := struct {
		Name  string
		Years int
	}{
		Name:  "Yoshi",
		Years: 10,
	}
	fmt.Printf("\nThis is a dog: %v", dog)

	book := Book{}

	fmt.Printf("\nThis is an empty book: %v", book)

	book.Owner = struct {
		Name string
		Age  int
	}{Name: "Nicholas", Age: 23}

	fmt.Printf("\nNow it has an owner: %v", book)

	book.Owner = struct {
		Name string
		Age  int
	}{}

	if book.Owner.Name == "" && book.Owner.Age == 0 {
		fmt.Printf("\nNow it doesn't have it anymore: %v", book)
	}

	book.Owner.Name = "Nicholas"

	fmt.Printf("\nNow it has an owner with just a name: %v", book)
}

func embeddedStruct() {

	type Car struct {
		Make  string
		Model string
	}

	type Truck struct {
		// "car" is embedded, so the definition of a
		// "truck" now also additionally contains all
		// of the fields of the car struct
		Car
		BedSize int
	}

	lanesTruck := Truck{
		BedSize: 10,
		Car: Car{
			Make:  "toyota",
			Model: "camry",
		},
	}
	fmt.Println("\n", lanesTruck.BedSize)

	// embedded fields promoted to the top-level
	// instead of lanesTruck.car.make
	fmt.Println(lanesTruck.Make)
	fmt.Println(lanesTruck.Model)

	fmt.Printf("Truck: %v", lanesTruck)
}


// METHODS
type Person struct {
	Name string
	BirthYear int
}

func (p Person) getAge() int {
	return time.Now().Year() - p.BirthYear
}


func executeMethod(){
	me := Person{Name: "Nicholas", BirthYear: 2000}

	fmt.Printf("\nI'm currently: %v years old", me.getAge())
}
