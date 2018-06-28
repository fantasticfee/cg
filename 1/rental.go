package rental

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
