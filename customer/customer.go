package Customer

import (
	"cg/rental"
	"cg/movie"
)
type Customer {
	Name string
	Rentals []rental.Rental
}

func (c *Customer) Customer (name string) {
	c.Name = name
}


func (c *Customer) AddRental (arg Rental) {
	c.Rentals = append(c.Rentals, arg)
}

func (c *Customer) GetName () string {
	return c.Name
}

func (c *Customer) Statement () string {
	totalAmount := 0
	frequentRenterPoints := 0

	result := "Rental Record for " + c.GetName() + "\n"
	while(len(c.Rentals) != 0){
		thisAmount = 0;
		for _,each := range c.Rentals {
			switch each.GetMovie().GetPriceCode() {
			case REGULAR:
				thisAmount += 2
				if GetDaysRented() > 2 {
					thisAmount += ()
				}
			case NEW_RELEASE:
				thisAmount += (each.GetDaysRented() - 2) * 1.5
			case CHILDRENS:
				thisAmount += 1.5
				if each.GetDaysRented() > 3 {
					thisAmount += (each.GetDaysRented() - 3) * 1.5
				}
			}
			frequentRenterPoints ++;
			if each.GetMovie().GetPriceCode() == NEW_RELEASE &&
			each.GetDaysRented() > 1 {
				frequentRenterPoints++;
			}
			result += "\t" + each.GetMovie().GetTitle() + "\t" + string(thisAmount) + "\n"
			totalAmount += thisAmount
		}

	}
		result += "Amount owed is" + string(totalAmount) + "\n"
		result += "You earned " + string(frequentRenterPoints) + " frequent renter points";
		return result
}

//第一个修改版本
func (c *Customer) StatementModify1 () string {
	totalAmount := 0

	result := "Rental Record for " + c.GetName() + "\n"
	for len(c.Rentals) != 0 {
		//var thisAmount float64
		for _,each := range c.Rentals {

			//thisAmount = each.GetCharge()

			result += "\t" + each.GetMovie().GetTitle() + "\t" + string(each.GetCharge()) + "\n"
			//totalAmount += thisAmount
		}

	}
		result += "Amount owed is" + string(c.getTotalCharge()) + "\n"
		result += "You earned " + string(c.GetFrequentRenterPoints) + " frequent renter points"
		return result
}

func (c *Customer) getTotalCharge() float64 {
	var result float64 = 0
	
		for _,each := range c.Rentals {
			oneMount := each.GetCharge()
			result += oneMount
	}

	return result
}

func (c *Customer) GetFrequentRenterPoints() float64 {
	frequentRenterPoints := 0
		for _,each := range c.Rentals {
			frequentRenterPoints += each.GetFrequentRenterPoints()
		}
		return frequentRenterPoints
}

func AmountFor(aRental Rental) float64 {
	var result float64 = 0;
			switch aRental.GetMovie().GetPriceCode() {
			case REGULAR:
				result += 2
				if GetDaysRented() > 2 {
					result += ()
				}
			case NEW_RELEASE:
				result += (aRental.GetDaysRented() - 2) * 1.5
			case CHILDRENS:
				result += 1.5
				if aRental.GetDaysRented() > 3 {
					result += (aRental.GetDaysRented() - 3) * 1.5
				}
			}
			return result
}