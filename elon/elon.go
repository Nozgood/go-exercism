package elon

import "fmt"

type Car struct {
    Battery 		int
    BatteryDrain	int
    Speed			int
    Distance		int
}

func (car *Car) Drive() {
	if car.BatteryDrain < car.Battery {
		car.Battery -= car.BatteryDrain
		car.Distance += car.Speed
	}
}

func (car *Car) DisplayDistance() string {
	stringDistance := fmt.Sprintf("Driven %v meters", car.Distance)
	return stringDistance
}

func (car *Car) DisplayBattery() string {
	percentage := "%"
	stringBattery := fmt.Sprintf("Battery at %v%v", car.Battery, percentage)
	return stringBattery
}

func (car *Car) CanFinish(trackDistance int) bool {
	timeDrive := car.Battery / car.BatteryDrain
	maxDistance := timeDrive * car.Speed
	if maxDistance > trackDistance {
		return true
	} else {
		return false
	}
}
