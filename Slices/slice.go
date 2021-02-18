package main

import "fmt"

func main() {
	names := [5]string{
		"John",
		"Hardey",
		"Amey",
		"Wish",
		"Aley",
	}
	fmt.Println(names)

	a := names[1:3]
	b := names[0:4]
	fmt.Println(a, b)

	b[1] = "XX"
	fmt.Println(a, b)
	fmt.Println(names)
}
