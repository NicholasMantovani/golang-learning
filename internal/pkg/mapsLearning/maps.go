package mapslearning

import "fmt"

// Any value can be passed in maps but not keys https://go.dev/blog/maps for example slices, map and functions cannot be used as keys because they can't be compared with ==
// Effective go for maps https://go.dev/doc/effective_go#maps

func ExecuteMaps() {
	myFirstMap()
	mapCrud()
	dogAges := map[string]int{
		"Zoe":   1,
		"Aston": 5,
	}
	fmt.Println("Before", dogAges)
	mapArePassedByReference(dogAges)
	fmt.Println("After", dogAges)
	executeUserExercise()
	nestedMaps()
}

func myFirstMap() {
	var ages map[string]int
	fmt.Println("Map zero value is NIL", ages)
	ages = make(map[string]int)
	ages["Nicholas"] = 23
	ages["Flora"] = 22
	fmt.Println("Before changes", ages)
	ages["Flora"] = 23
	fmt.Println("After changes", ages)

	dogAges := map[string]int{
		"Zoe":   1,
		"Aston": 5,
	}
	fmt.Println("Map with default values", dogAges)

}

func mapCrud() {
	ages := make(map[string]int)
	fmt.Println("Created map", ages)

	ages["Nicholas"] = 23
	fmt.Println("Added entry to map", ages)

	fmt.Println("Get the value of a map", ages["Nicholas"])

	delete(ages, "Nicholas")
	fmt.Println("Deleted entry of a map", ages)

	elem, ok := ages["Nicholas"]

	fmt.Println("Check if a value is in a map and get its value", elem, ok)
}

func mapArePassedByReference(ages map[string]int) {
	ages["Siummista"] = 102
}

// Count how many time a user is in the list
func executeUserExercise() {
	users := []string{"Elon", "Bezos", "Sium", "Elon", "Sium", "Bezos", "Nicholas"}
	fmt.Println("Users", users)

	fmt.Println("Map of how many time a user is in the Users list", countUser(users))
}

func countUser(users []string) map[string]int {
	output := map[string]int{}

	for _, user := range users {
		if _, ok := output[user]; !ok {
			output[user] = 0
		}

		output[user]++
	}
	return output
}

// Boot.dev example
func nestedMaps() {
	names := []string{
		"Grant", "Eduardo", "Peter", "Matthew", "Matthew", "Matthew", "Peter", "Peter", "Henry", "Parker", "Parker", "Parker", "Collin", "Edward", "", "",
	}
	fmt.Println("All the names are", names)
	output := getNameCounts(names)
	fmt.Println("All the names for the first unique", output)
	for key, value := range output {
		fmt.Printf("\nKey: %v | value: %v", string(key), value)
	}
}

// Complete the getNameCounts function.
// It takes a slice of strings (names) and returns a nested map where the first key is all
// the unique first characters of the names, the second key is all the names themselves, and the value is the count of each name.
func getNameCounts(names []string) map[rune]map[string]int {
	output := map[rune]map[string]int{}
	for _, name := range names {
		if name == "" {
			continue
		}
		firstNameChar := rune(name[0])
		if _, ok := output[firstNameChar]; !ok {
			output[firstNameChar] = map[string]int{}
		}
		// this check could be omitted since the ZERO VALUE of an int is 0
		if _, ok := output[firstNameChar][name]; !ok {
			output[firstNameChar][name] = 0
		}
		output[firstNameChar][name]++
	}
	return output
}
