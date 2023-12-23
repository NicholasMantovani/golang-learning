package generics

import "fmt"

// WHY GENERICS?
// GENERICS REDUCE REPETITIVE CODE
// You should care about generics because they mean you don’t have to write as much code! It can be frustrating to write the same logic over and over again, just because you have some underlying data types that are slightly different.

// GENERICS ARE USED MORE OFTEN IN LIBRARIES AND PACKAGES
// Generics give Go developers an elegant way to write amazing utility packages. While you will see and use generics in application code, I think it will be much more common to see generics used in libraries and packages. Libraries and packages contain importable code intended to be used in many applications, so it makes sense to write them in a more abstract way. Generics are often the way to do just that!

// WHY DID IT TAKE SO LONG TO GET GENERICS?
// Go places an emphasis on simplicity. In other words, Go has purposefully left out many features to provide its best feature: being simple and easy to work with.

// According to historical data from Go surveys, Go’s lack of generics has always been listed as one of the top three biggest issues with the language. At a certain point, the drawbacks associated with the lack of a feature like generics justify adding complexity to the language.

func ExecuteGenerics() {
	generics()
	genericsOnInterface()
	storeExample()
}

func generics() {

	intSlice := []int{1, 2, 3, 4}
	stringSlice := []string{"nicholas", "mantovani"}

	fmt.Println(splitIntSlice(intSlice))
	fmt.Println(splitStringSlice(stringSlice))
	fmt.Println(splitSlice[int](intSlice))
	fmt.Println(splitSlice[string](stringSlice))
}

// instead of these 2 method
func splitIntSlice(s []int) ([]int, []int) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

func splitStringSlice(s []string) ([]string, []string) {
	mid := len(s) / 2
	return s[:mid], s[mid:]
}

// I can have one method only

func splitSlice[T any](input []T) ([]T, []T) {
	mid := len(input) / 2
	return input[:mid], input[mid:]
}

type ToString interface {
	toString() string
}

type Object struct {
	Name string
	UUID string
}

func (o Object) toString() string {
	return o.Name + " - " + o.UUID
}

type User struct {
	Name    string
	Surname string
}

func (u User) toString() string {
	return u.Name + " - " + u.Surname
}

func printObjectToString[T ToString](input T) {
	fmt.Println("This is the toString of the object", input, input.toString())
}

func genericsOnInterface() {
	o := Object{"Test", "123"}
	u := User{"Nicholas", "Mantovani"}

	printObjectToString[ToString](o)
	printObjectToString[ToString](u)
	printObjectToString[Object](o)
	printObjectToString[User](u)
}

// INTERFACE TYPE LISTS
// When generics were released, a new way of writing interfaces was also released at the same time!

// We can now simply list a bunch of types to get a new interface/constraint.

// Ordered is a type constraint that matches any ordered type.
// An ordered type is one that supports the <, <=, >, and >= operators.
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// PARAMETRIC CONSTRAINTS
// Your interface definitions, which can later be used as constraints, can accept type parameters as well.

// The store interface represents a store that sells products.
// It takes a type parameter P that represents the type of products the store sells.
type store[P product] interface {
	Sell(P)
}

type product interface {
	Price() float64
	Name() string
}

type book struct {
	title  string
	author string
	price  float64
}

func (b book) Price() float64 {
	return b.price
}

func (b book) Name() string {
	return fmt.Sprintf("%s by %s", b.title, b.author)
}

type toy struct {
	name  string
	price float64
}

func (t toy) Price() float64 {
	return t.price
}

func (t toy) Name() string {
	return t.name
}

// The bookStore struct represents a store that sells books.
type bookStore struct {
	booksSold []book
}

// Sell adds a book to the bookStore's inventory.
func (bs *bookStore) Sell(b book) {
	bs.booksSold = append(bs.booksSold, b)
}

// The toyStore struct represents a store that sells toys.
type toyStore struct {
	toysSold []toy
}

// Sell adds a toy to the toyStore's inventory.
func (ts *toyStore) Sell(t toy) {
	ts.toysSold = append(ts.toysSold, t)
}

// sellProducts takes a store and a slice of products and sells
// each product one by one.
func sellProducts[P product](s store[P], products []P) {
	for _, p := range products {
		s.Sell(p)
	}
}

func storeExample() {
	bs := bookStore{
		booksSold: []book{},
	}

	// By passing in "book" as a type parameter, we can use the sellProducts function to sell books in a bookStore
	sellProducts[book](&bs, []book{
		{
			title:  "The Hobbit",
			author: "J.R.R. Tolkien",
			price:  10.0,
		},
		{
			title:  "The Lord of the Rings",
			author: "J.R.R. Tolkien",
			price:  20.0,
		},
	})
	fmt.Println(bs.booksSold)

	// We can then do the same for toys
	ts := toyStore{
		toysSold: []toy{},
	}
	sellProducts[toy](&ts, []toy{
		{
			name:  "Lego",
			price: 10.0,
		},
		{
			name:  "Barbie",
			price: 20.0,
		},
	})
	fmt.Println(ts.toysSold)
}
