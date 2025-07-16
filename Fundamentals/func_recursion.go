package main

//Factorial with recursion
import (
	"fmt"
	"log"
)

func recursion_test() {
	var num int
	fmt.Println("Enter Number to print Factorial")
	_, err := fmt.Scanf("%d", &num)
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println("Calculation factorial for : ", num)
	fmt.Printf("The Factorial of %d is %d", num, factorial(num))

}
func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
	/*
		factorial of 0 is 1 , and hence it base condition
		factorial of 1 is 1 , 1 * 1 = 1
		factorial of 2 is 2 , 2 * 1 * 1 = 2
		factorial of 3 is 6 , 3 * factorial of 3-1,
		factorial of 4 is 24 , 4 * factorial of 3 or 4-1
	*/
}

// Function factorial is being called for every single value of n till it reduce to zero , hence it's O(n)
