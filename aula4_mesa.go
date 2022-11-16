package main

import "fmt"

type Product struct {
	ID          int
	Name        string
	Price       float64
	Description string
	Category    string
}

var (
	products = []Product{
		{
			ID:          1,
			Name:        "Cachaça",
			Price:       55.50,
			Description: "Da boa",
			Category:    "Bebida alcoólica",
		},
		{
			ID:          2,
			Name:        "Smirnoff",
			Price:       60.80,
			Description: "Russa",
			Category:    "Bebida alcoólica",
		},
	}
)

func (p *Product) save() {
	products = append(products, *p)
}

func (p *Product) getAll() {
	fmt.Println(products)
}

func getById(id int) Product {
	var result Product

	for _, p := range products {
		if p.ID == id {
			result = p
		}
	}
	return result
}

func main() {
	p1 := Product{
		ID:          3,
		Name:        "Macarrão",
		Price:       5,
		Description: "instantâneo",
		Category:    "massa",
	}

	p1.save()

	p1.getAll()

	fmt.Println(getById(1))
}
