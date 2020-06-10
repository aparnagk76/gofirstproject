package animal

type Animal interface{
	Eat(string) string
	Sleep(int) int
	Sound(string) string
}