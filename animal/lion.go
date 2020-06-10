package animal

type Lion struct {

	name string
	color string
	height float
	weight float

}

func (l Lion) Eat(name string) string {
		
	fmt.Println("eating food", name)
	return name
}

func (l Lion) Sleep(sleepTime int) string {
	fmt.Println("lion sleeping time",sleepTime)
	return sleepTime
}

func (l Lion) Sound(Sound string) string {
	fmt.Println("lion sound", sound)
	return sound
}