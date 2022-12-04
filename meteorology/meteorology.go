package meteorology

import "fmt"

type Stringer interface{
	String() string
}

type TemperatureUnit int

const (
	Celsius    TemperatureUnit = 0
	Fahrenheit TemperatureUnit = 1
)

// Add a String method to the TemperatureUnit type	
func (tempUnit TemperatureUnit) String() string {
	/* if tempUnit == Celsius {
		return "°C"
	} else if tempUnit == Fahrenheit {
		return "°F"
	}
	return "" */
	units := []string{"°C", "°F"}
	return units[tempUnit]
}

type Temperature struct {
	Degree int
	Unit   TemperatureUnit
}

// Add a String method to the Temperature type
func (temp Temperature) String() string {
	degreeString := fmt.Sprint(temp.Degree)
	if temp.Unit == Celsius {
		return degreeString + "°C"
	} else if temp.Unit == Fahrenheit {
		return  degreeString + "°F"
	}
	return "" 
}

type SpeedUnit int

const (
	KmPerHour    SpeedUnit = 0
	MilesPerHour SpeedUnit = 1
)

// Add a String method to SpeedUnit
func (speedUnit SpeedUnit) String() string { 
	if speedUnit == KmPerHour {
		return "km/h"
	} else if speedUnit == MilesPerHour {
		return "mph"
	}
	return ""
}

type Speed struct {
	magnitude int
	unit      SpeedUnit
}

// Add a String method to Speed
func (speed Speed) String() string {
	magnitudeString := fmt.Sprint(speed.magnitude)
	speedUnit := speed.unit.String()
	return magnitudeString + " " + speedUnit
}
type MeteorologyData struct {
	location      string
	temperature   Temperature
	windDirection string
	windSpeed     Speed
	humidity      int
}

// Add a String method to MeteorologyData
func (meteo MeteorologyData) String() string {
    temperatureString := meteo.temperature.String()
	windString := meteo.windSpeed.String()
	humidityString := fmt.Sprint(meteo.humidity)

	return meteo.location + ":" + temperatureString + ", Wind " + meteo.windDirection + " at " + windString + ", " + humidityString + "% Humidity"
}