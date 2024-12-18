package main

import (
	"fmt"
	"maps"
)

func main() {
	//creating map

	// map[keyType]ReturnType
	m := make(map[string]string)

	m["name"] = "Hello World"
	m["description"] = "This is world"

	// get a element
	fmt.Println("map values", m)
	fmt.Println("map name", m["name"])

	// if key does not exist in map it returns zeros value
	fmt.Println("zeros value", m["title"])

	// get map length
	fmt.Println("length", len(m))

	// delete value from map

	delete(m, "description")

	fmt.Println("map values", m)

	// empty the map

	clear(m)

	fmt.Println("map values", m)

	// define map in init

	newMap := map[string]int{"phone": 123, "age": 25}

	fmt.Println("newMap values", newMap)

	// value return with value and flag

	phone, found := newMap["phone"]

	if found {
		fmt.Println("Phone number:", phone)
	} else {
		fmt.Println("Phone number not found")
	}

	newMap2 := map[string]int{"phone": 22, "age": 25}

	// check difference between map
	fmt.Println(maps.Equal(newMap, newMap2))
}
