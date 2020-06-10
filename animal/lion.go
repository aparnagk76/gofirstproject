package animal
import "fmt"

type Lion struct {

	Name string
	Color string
	Height float64
	Weight float64

}

func (l Lion) Eat(name string) string {
		
	fmt.Println("lion eating", name)
	return name
}

func (l Lion) Sleep(sleepTime int) int {
	fmt.Println("lion sleeping time",sleepTime)
	return sleepTime
}

func (l Lion) Sound(sound string) string {
	fmt.Println("lion sound", sound)
	return sound
}