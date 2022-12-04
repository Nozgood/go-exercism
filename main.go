package main

import (
	"fmt"
	"go-exercism/meteorology"
)

func main () {
	testTemperature := meteorology.Temperature {
		Degree: 42,
		Unit: meteorology.Fahrenheit,
	}
	testDisplayTemperature := testTemperature.String()
	
	fmt.Print(testDisplayTemperature)
}