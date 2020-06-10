package animal

type Animal interface{
	Eat(string) string
	Sleep(int) string
	Sound(string) string
}