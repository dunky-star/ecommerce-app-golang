package concept

import "fmt"

func MasterConcepts() {
	// Basic Types: int, float64, string, bool
	// Composite Types: array, slice, map, struct
	// Pointer Types: *
	//var age int64
	//var height float64
	//var firstName string
	//var isEmployed bool

	age := 29
	height := 138.9
	firstName := "Geoffrey"
	isEmployed := true

	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Height: %.2f\n", height)
	fmt.Printf("First Name: %s\n", firstName)
	fmt.Printf("Emplyed? %v\t", isEmployed)

	// Array Data structure - Fixed size
	var myFamily [4]string
	myFamily[0] = "Arma"
	myFamily[1] = "Nyeko"
	myFamily[2] = "Robinson"
	myFamily[3] = "Mel"
	fmt.Printf("\nMy family: %v\n", myFamily)

	// Alternatively
	hisFamily := [3]string{"John", "Christine", "Palmo"}
	fmt.Printf("His family: %v\n", hisFamily)

	// Multi-dimensional array
	myCourses := [3][2]string{
		{"Go", "Java"},
		{"AWS", "GCP"},
		{"CDK", "Pulumi"},
	}
	fmt.Printf("Available courses: %v\n", myCourses)

	// Slice Data structure - Dynamic size
	var myFriends []string
	myFriends = append(myFriends, "Mike", "Deo", "Hill")
	fmt.Printf("\nMy friends: %v\n", myFriends)

	mySliceCourses := [][]string{
		{"Python", "C++"},
		{"Terrafoam", "Ansible"},
		{"NodeJS", "Angular"},
	}

	mySliceCourses = append(mySliceCourses, []string{"IAC", "Cloud Formation"})

	fmt.Printf("Available courses part 2: %v\n", mySliceCourses)

}
