package animal
import "fmt"

type Bird struct {

	Name string
	Color string
	Height float64
	Weight float64

}

func (b Bird) Eat(name string) string {
		
	fmt.Println("bird eating", name)
	return name
}

func (b Bird) Sleep(sleepTime int) int {
	fmt.Println("bird sleeping time",sleepTime)
	return sleepTime
}

func (b Bird) Sound(sound string) string {
	fmt.Println("bird sound", sound)
	return sound
}