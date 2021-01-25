package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello, Alnur")

	var temp int
	temp = 11
	var variable = 10
	fmt.Println(temp + variable)

	var name, age = "Alnur", 21
	fmt.Println(name, age)

	var input string
	fmt.Scanln(&input)
	fmt.Println(input)

	hey("Alnur")
	fmt.Println(heyheyhey("Alnur"))
}

func hey(name string) {
	fmt.Println("Hey, " + name)
}

func heyheyhey(name string) string {
	return "heyheyhey, " + name
}
