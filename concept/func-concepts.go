package concept

import "fmt"

func FuncConcepts() {
	fmt.Println("Total amount: ", calculateTotal(1.1, 1.3, 1.5, 10.0))

	p := Product{
		Name:  "Lenovo Thinkpad X12 Carbon",
		Price: 2000,
		Stock: 4,
	}

	fmt.Printf("Total product amount: %.2f", p.calculate(5)) // Calling a receiver function
	p.reduceStock(2)
	fmt.Println("\n", p)

}

func getUserName() string {
	return "My name is Geoffrey Opiyo"
}
func getUserById(id int) (name string, age int) {
	// DB call
	println(id)
	return "Geoffrey", 29
}

func calculateTotal(products ...float64) float64 {
	totalAmount := 0.0

	for _, price := range products {
		totalAmount += price
	}

	return totalAmount
}

// Receiver Function: Is a special kind of function that can only be performed on a specific type.
type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p *Product) calculate(qty int) float64 {
	return p.Price * float64(qty)
}

func (p *Product) reduceStock(qty int) {
	if p.Stock >= qty {
		p.Stock -= qty
	}
}
