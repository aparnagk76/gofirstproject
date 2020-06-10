package animal
import "fmt"

type Cat struct {

	Name string
	Color string
	Height float64
	Weight float64

}

func (c Cat) Eat(name string) string {
		
	fmt.Println("cat eating food", name)
	return name
}

func (c Cat) Sleep(sleepTime int) int {
	fmt.Println("cat sleeping time",sleepTime)
	return sleepTime
}

func (c Cat) Sound(sound string) string {
	fmt.Println("cat sound", sound)
	return sound
}