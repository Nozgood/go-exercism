package farm

import "fmt"

type Animal interface {
	Noise() string
	NumberOfLegs() int 
}

type Cat struct {
	Name string
	Breed string
}

type Spider struct {
	Name string
	Breed string
	Venomous bool
}

func (c *Cat) Noise() string {
	return "Miaou"
}

func (s *Spider) Noise() string {
	return "Silence..."
}

func (c *Cat) NumberOfLegs() int {
	return 4
}

func (s *Spider) NumberOfLegs() int {
	return 8
}

func AllAnimalInfo(a Animal) {
	fmt.Println("Noise : ", a.Noise(), "\n numbers of legs : ", a.NumberOfLegs())
}