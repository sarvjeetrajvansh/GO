package main

import (
	"fmt"
	"math"
)

const pi float64 = 2.17 // const type var, can be declare anywhere

var radius int = 5

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
}

func counterGenerator(number int) {
	for i := 1; i <= number; i++ {
		fmt.Print(i, "->")
	}
}

func main() {
	fmt.Println("Learning Const & Var")
	ans := 2 * pi * math.Pow(float64(radius), 2) // short hand var declare & assign
	fmt.Println(ans)
	fmt.Println(add(radius, radius))         //
	fmt.Println(multiReturn(radius, radius)) // accept multiple return
	_, result := multiReturn(radius, radius) // accept only one return
	fmt.Println(result)
	tableGenerator(1000)
	tableGenerator(1000000000000000000) // output shows runtime rounding off towards -ve range
	counterGenerator(10)
}
