package animal

type Cat struct {

	name string
	color string
	height float
	weight float

}

func (c Cat) Eat(name string) string {
		
	fmt.Println("cat eating food", name)
	return name
}

func (c Cat) Sleep(sleepTime int) string {
	fmt.Println("cat sleeping time",sleepTime)
	return sleepTime
}

func (c Cat) Sound(sound string) string {
	fmt.Println("cat sound", sound)
	return sound
}