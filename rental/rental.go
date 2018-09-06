package rental

import (
	"cg/movie"
)

type Rental struct {
	Movie      movie.Movie
	DaysRented int
}

func (r *Rental) Rental(movie movie.Movie, daysRented int) {
	r.Movie = movie
	r.DaysRented = daysRented
}

func (r *Rental) GetDaysRented() int {
	return r.DaysRented
}

func (r *Rental) GetMovie() movie.Movie {
	return r.Movie
}

//old 版本
func (r *Rental) GetCharge() float64 {
	var result float64 = 0
	switch r.Movie.GetPriceCode() {
	case movie.REGULAR:
		result += 2
		if r.DaysRented > 2 {
			result = result + float64((r.DaysRented-2))*1.5
		}
	case movie.NEWRELEASE:
		result += float64(r.DaysRented * 3)
	case movie.CHILDRENS:
		result += 1.5
		if r.DaysRented > 3 {
			result += float64((r.DaysRented - 3)) * 1.5
		}
	}
	return result
}

//版本改动
func (r *Rental) GetChargeNew() float64 {
	return r.Movie.GetCharge(r.DaysRented)
}

func (r *Rental) GetFrequentRenterPoints() int {
	var resutl int = 0
	resutl++
	if r.Movie.GetPriceCode() == movie.NEWRELEASE &&
		r.GetDaysRented() > 1 {
		resutl += 1
	}
	return resutl
}

//new
func (r *Rental) GetFrequentRenterPointsNew() int {
	return r.Movie.GetFrequentRenterPoints(r.DaysRented)
}
