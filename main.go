package main

import (
	"fmt"
	animal "github.com/aparnagk76/gofirstproject/animal"
)

func main() {

	var myAnimal animal.Animal = animal.Dog{
		name :"tommy",
		color: "white",
		height: 3.5,
		weight : 5.5,
	}
	myAnimal.Eat("dog biscuits")
}