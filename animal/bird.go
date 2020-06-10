package animal

type Bird struct {

	name string
	color string
	height float
	weight float

}

func (b Bird) Eat(name string) string {
		
	fmt.Println("eating food", name)
	return name
}

func (b Bird) Sleep(sleepTime int) string {
	fmt.Println("bird sleeping time",sleepTime)
	return sleepTime
}

func (b Bird) Sound(sound string) string {
	fmt.Println("bird sound", sound)
	return sound
}