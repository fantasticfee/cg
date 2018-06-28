package customer

import (
	"rental"
)
type Customer {
	name string
	rentals []Rental
}

func (c *Customer) Customer (name string) {
	c.name = name
}


func (c *Customer) AddRental (arg Rental) {
	c.rentals = append(c.rentals, arg)
}

func (c *Customer) GetName () string {
	return c.name
}

func (c *Customer) Statement () string {
	totalAmount := 0
	frequentRenterPoints := 0

	result := "Rental Record for " + c.GetName() + "\n"
	while(len(c.rentals) != 0){
		thisAmount = 0;
		for _,each := range c.rentals {
			switch each.GetMovie().GetPriceCode() {
			case Movie.REGULAR:
				thisAmount += 2
				if each.GetDaysRented() > 2 {
					thisAmount += ()
				}
			}
		}

	}
}