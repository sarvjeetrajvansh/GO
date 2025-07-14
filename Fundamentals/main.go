package main

import (
	"Fundamentals/Fundamentals/msg"
	"fmt"
	"log"
)

func main() {
	var input string

	greeter()
	goVarAndConst()
	day1Main()

	fmt.Print("\nEnter your Name: ")
	_, err := fmt.Scanln(&input) // waits for input

	if err != nil {
		log.Fatal("Error:", err)
	}

	msg.Msg(input) // call function from msg package

	arraySlice()
}
