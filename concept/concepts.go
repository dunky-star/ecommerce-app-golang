package concept

import "fmt"

func MasterConcepts() {
	var age int64
	var height float64
	var firstName string
	var isEmployed bool

	age = 29
	height = 138.9
	firstName = "Geoffrey"
	isEmployed = true

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("First Name: %s\n", firstName)
	fmt.Printf("Emplyed? %v\t", isEmployed)
}
