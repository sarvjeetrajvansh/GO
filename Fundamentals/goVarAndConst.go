package main

import (
	"fmt"
	"math"
)

const pi float64 = 2.17 // const type var, can be declare anywhere
var radius int = 5
var name string

func add(num1, num2 int) int {
	sum := num1 + num2
	return sum
}

func multiReturn(num1 int, num2 int) (int, string) {
	return (num1 + num2), "Success"
}

func tableGenerator(number int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(number, "*", i, "=", number*i)
	}
	//  same as above with range
	for t := range 11 {
		fmt.Println(number, "*", t, "=", number*t)
	}

}

func counterGenerator(number int) {
	for i := 1; i <= number; i++ {
		fmt.Print(i, "->")
	}
}

func ciCalculator(principal float64, rate float64, time float64) {
	const inflationRate = 6.48
	amount := principal * math.Pow(1+rate/100, time)
	amountPostInflation := amount / math.Pow(1+inflationRate/100, time)

	fmt.Printf("\nThe Amount will be Rs.%.2f after %.2f Years at rate of %.2f for principal of Rs.%.2f \n", amount, time, rate, principal)

	fmt.Printf("Considering inflation, adusted value will be : Rs.%.2f", amountPostInflation)
}

func main() {
	fmt.Println("Learning Const & Var")
	ans := 2 * pi * math.Pow(float64(radius), 2) // short hand var declare & assign
	fmt.Println(ans)
	fmt.Println(add(radius, radius))         //
	fmt.Println(multiReturn(radius, radius)) // accept multiple return
	_, result := multiReturn(radius, radius) // accept only one return
	fmt.Println(result)
	tableGenerator(5)
	// output shows runtime rounding off towards -ve range
	counterGenerator(10)
	ciCalculator(1000, 8.5, 10)
	fmt.Println("\nEnter your Name :")
	fmt.Scanf("%v", &name) // waits for input
	msg(name)              // calling from hello.go in same package
}
