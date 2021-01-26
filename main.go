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

	if 5 > 6 {
		fmt.Println("five is greater than six")
	} else {
		fmt.Println("five is smaller than six")
	}

	var i int

	for i = 1; i < 3; i++ {
		fmt.Println(i)
	}

	var names = [2]string{"Alnur the first", "Alnur the second"}
	fmt.Println(names)

	var namess = []string{}
	namess = append(namess, "Alnur")
	namess = append(namess, "Alnur2")
	namess = append(namess, "Alnur3")
	fmt.Println(namess)

	var value string
	for i, value = range namess {
		fmt.Println(i, value)
	}

	var department = map[string]string{
		"Head of department": "Alnur",
		"Slave 1":            "Gani",
	}
	fmt.Println(department)
	fmt.Println(department["Head of department"])

	var emptydepartment = make(map[string]string)
	emptydepartment["New head"] = "Aleka"
	fmt.Println(emptydepartment)
}

func hey(name string) {
	fmt.Println("Hey, " + name)
}

func heyheyhey(name string) string {
	return "heyheyhey, " + name
}
