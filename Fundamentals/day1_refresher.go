// day1_refresher.go
package main

import "fmt"

// Const
const PI = 3.14159

func day1Main() {
	//Variables & Types
	var a int = 5
	b := 10                // shorthand declaration
	var c, d = 15, "Go"    // multiple declaration
	var f float64 = PI * 2 // float calculation

	fmt.Println("a:", a, "b:", b, "c:", c, "d:", d, "f:", f)

	//Swap a & b without a temp variable
	a = a + b
	b = a - b
	a = a - b
	fmt.Println("After swap - a:", a, "b:", b)

	// Even or Odd
	if a%2 == 0 {
		fmt.Println(a, "is even")
	} else {
		fmt.Println(a, "is odd")
	}

	// ➕ Sum of first 10 natural numbers
	sum := 0
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("Sum 1 to 10:", sum)

	// ✖️ Multiplication Table for b
	fmt.Println("Table of", b)
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d x %d = %d\n", b, i, b*i)
	}

	//Fibonacci first 10 terms (non-recursive)
	n1, n2 := 0, 1
	fmt.Print("Fibonacci: ", n1, n2)
	for i := 0; i < 8; i++ {
		next := n1 + n2
		fmt.Print(" ", next)
		n1, n2 = n2, next
	}
	fmt.Println()

	//Largest of three numbers
	e := -5
	largest := a
	if c > largest {
		largest = c
	}
	if e > largest {
		largest = e
	}
	fmt.Println("Largest among", a, c, e, "is", largest)
}
