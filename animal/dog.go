package animal
 import "fmt"

type Dog struct {

	Name string
	Color string
	Height float64
	Weight float64

}

func (d Dog) Eat(name string) string {
		
	fmt.Println("dog eating food", name)
	return name
}

func (d Dog) Sleep(sleepTime int) int {
	fmt.Println("dog sleeping time", sleepTime)
	return sleepTime
}

func (d Dog) Sound(sound string) string {
	fmt.Println("dog barking sound", sound)
	return sound
}