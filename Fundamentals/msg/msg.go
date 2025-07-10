package msg

import "fmt"

func Msg(name string) {
	fmt.Println("Welcome to the msg package")
	fmt.Println("This is a simple message function that greets the user.")
	fmt.Printf("\nHello %v ", name)
}
