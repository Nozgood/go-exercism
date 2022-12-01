package main

import (
	"fmt"
	"go-exercism/elon"
)
func main () {
	testCar := elon.Car{
		Battery: 100,
		BatteryDrain: 20,
		Speed: 10,
		Distance: 0,
	}
	fmt.Printf("%v \n%v", testCar.DisplayDistance(), testCar.DisplayBattery())
}