package main

import (
	"fmt"
	"go-exercism/parsinglogs"
)

func main () {
	testValidLine := parsinglogs.IsValidLine("[BOB] Any old text")
	fmt.Print(testValidLine)
}