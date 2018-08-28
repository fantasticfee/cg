package cg1

type Rental struct {
	movie      Movie
	daysRented int
}

func (r *Rental) Rental(movie Movie, daysRented int) {
	r.movie = movie
	r.daysRented = daysRented
}

func (r *Rental) GetDaysRented() int {
	return r.daysRented
}

func (r *Rental) GetMovie() Movie {
	return r.movie
}

func (r *Rental)GetCharge() float64 {
	float64 result = 0;
			switch r.GetMovie().GetPriceCode() {
			case REGULAR:
				thisAmount += 2
				if GetDaysRented() > 2 {
					thisAmount += ()
				}
			case NEW_RELEASE:
				thisAmount += (GetDaysRented() - 2) * 1.5
			case CHILDRENS:
				thisAmount += 1.5
				if aRental.GetDaysRented() > 3 {
					thisAmount += (GetDaysRented() - 3) * 1.5
				}
			}
			return result
}

func (r *Rental)GetFrequentRenterPoints() int {
			resutl := 0
			resutl++;
			if each.GetMovie().GetPriceCode() == NEW_RELEASE &&
			each.GetDaysRented() > 1 {
				result++;
			}
			return result
}