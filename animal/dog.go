package animal

type Dog struct {

	name string
	color string
	height float
	weight float

}

func (d Dog) Eat(name string) string {
		
	fmt.Println("dog eating food", name)
	return name
}

func (d Dog) Sleep(sleepTime int) string {
	fmt.Println("dog sleeping time", sleepTime)
	return sleepTime
}

func (d Dog) Sound(sound string) string {
	fmt.Println("dog barking sound", sound)
	return sound
}