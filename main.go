package main

import (
	
	animal "github.com/aparnagk76/gofirstproject/animal"
)

func main() {

	var myAnimal animal.Animal
	myAnimal = animal.Dog{
		Name :"tommy",
		Color: "white",
		Height: 3.5,
		Weight : 5.5,
	}
	myAnimal.Eat("biscuits")
	myAnimal.Sleep(5)
	myAnimal.Sound("bow bow")
	
	myAnimal = animal.Lion{
		Name :"brio",
		Color: "cream",
		Height: 3.5,
		Weight : 5.5,
	}
	myAnimal.Eat("stage")
	myAnimal.Sleep(8)
	myAnimal.Sound("huh huh")
	
	myAnimal = animal.Bird{
		Name :"parrot",
		Color: "red",
		Height: 3.5,
		Weight : 5.5,
	}
	myAnimal.Eat("seeds")
	myAnimal.Sleep(7)
	myAnimal.Sound("la la la")

	myAnimal = animal.Cat{
		Name :"minii",
		Color: "blck",
		Height: 2.5,
		Weight : 1.9,
	}
	myAnimal.Eat("milk")
	myAnimal.Sleep(6)
	myAnimal.Sound("mauw mauw")
}