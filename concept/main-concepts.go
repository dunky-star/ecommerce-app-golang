package concept

import (
	"fmt"
	"log"
)

func MasterConcepts() {
	// Basic Types: int, float64, string, bool //
	// Composite Types: array, slice, map, struct //
	// Pointer Types: * //

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

	// Utilizing make to create slice
	myDeCourses := make([]string, 2, 10)
	// Insert values
	myDeCourses[0] = "Best course 1"
	myDeCourses[1] = "Best Course 2"
	// Appending values beyond the initial length
	myDeCourses = append(myDeCourses, "Star course 3", "Story course 4", "Pimp course 5")
	fmt.Printf("Available DE courses part 2: %v\n", myDeCourses)

	// Map data structure -> Key, Value pair
	myWishList := make(map[string]string)
	myWishList["first"] = "Lenovo ThinkPad X12"
	myWishList["second"] = "2 million dollars"
	myWishList["third"] = "A beautiful car"
	myWishList["fourth"] = "A plane"
	myWishList["fifth"] = "A beautiful Guitar"
	// To remove an element from the map
	delete(myWishList, "fifth")
	myFirstWish := myWishList["first"]
	log.Println("My first wish list: ", myFirstWish)
	fmt.Printf("My wish list: %v\n", myWishList)

	// Struct Data Type
	type Details struct {
		Description string `json:"description"`
		Images      string `json:"images"`
	}
	type Product struct {
		Name    string           `json:"name"`
		Price   float64          `json:"price"`
		Details `json:"details"` // Nested struct
	}

	//var product Product
	product := Product{
		Name:  "Lenovo Thinkpad X12 Carbon",
		Price: 2000,
		Details: Details{
			Description: "An incredible laptop with real compute power",
			Images:      "http://lenovothinkpad.jpg",
		},
	}

	fmt.Printf("My product: %v", product)

	/* Conditional statement & Pointers in Go
		- if else
	    - switch case
		- select: Select is only applicable when we are working with Goroutines & the channels
	*/
	// What is a Pointer ? -> A pointer is a kind of variable that holds memory address.
	// Geoffrey => Laptop
	// Guest => Geoffrey => Laptop
	geoffrey := "Lenovo Thinkpad Laptop"
	fmt.Println("\nGeoffrey's laptop:", geoffrey)

	//var guest *string
	guest := &geoffrey
	fmt.Println("Guest: ", *guest)

}
