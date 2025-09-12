package main

type Coffee struct {
	name     string
	quantity int
	price    float64
}

func NewCoffee(name string, quantity int, price float64) *Coffee {
	return &Coffee{
		name:     name,
		quantity: quantity,
		price:    price,
	}
}
