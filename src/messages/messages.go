package messages

import "fmt"

func Greet(name string) string {
	return fmt.Sprint("Hello, %v\n", name)
}

func depart(name string) string {
	return fmt.Sprint("Goodbye, %v\n", name)
}
