package main

import (
	"go-exercism/farm"
)

func main () {
	cat := farm.Cat{
		Name: "Miaou",
		Breed: "maine coon",
	}
	spider := farm.Spider{
		Name: "Peter Parker",
		Breed: "Human Spider",
		Venomous: false,
	}

	farm.AllAnimalInfo(&spider)
	farm.AllAnimalInfo(&cat)
}