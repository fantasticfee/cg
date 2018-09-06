package movie

const (
	REGULAR    = 0
	NEWRELEASE = 1
	CHILDRENS  = 2
)

type Movie struct {
	Title     string
	PriceCode int
}

func (m *Movie) Movie(title string, priceCode int) {
	m.Title = title
	m.PriceCode = priceCode
}

func (m *Movie) GetPriceCode() int {
	return m.PriceCode
}

func (m *Movie) SetPriceCode(arg int) {
	m.PriceCode = arg
}

func (m *Movie) GetTitle() string {
	return m.Title
}

func (m *Movie) GetCharge(daysRented int) float64 {
	var result float64 = 0
	switch m.GetPriceCode() {
	case REGULAR:
		result += 2
		if daysRented > 2 {
			result = result + float64((daysRented-2))*1.5
		}
	case NEWRELEASE:
		result += float64(daysRented * 3)
	case CHILDRENS:
		result += 1.5
		if daysRented > 3 {
			result += float64((daysRented - 3)) * 1.5
		}
	}
	return result
}

func (m *Movie) GetFrequentRenterPoints(daysRented int) int {
	var resutl int = 0
	resutl++
	if m.GetPriceCode() == NEWRELEASE &&
		daysRented > 1 {
		resutl += 1
	}
	return resutl
}
