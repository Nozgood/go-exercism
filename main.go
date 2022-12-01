package main

import (
	"fmt"
	"go-exercism/logs"
)

func main () {
	stringTest := "❗ recommended search product 🔍"
	// oldRuneTest := '❗'
	// newRuneTest := '🔍'
	fmt.Println(logs.WithinLimit(stringTest, 29))
}