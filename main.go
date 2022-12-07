package main

import (
	"fmt"
	"go-exercism/census"
)

func main () {
	name := "Matthew Sanabria"
	age := 29
	address := map[string]string{"street": "Main St."}

	test := census.NewResident(name, age, address)
	test.Delete()
	fmt.Print(test)
}