package movie

const (
	REGULAR    = 0
	NEWRELEASE = 1
	CHILDRENS  = 2
)

type Movie struct {
	title     string
	priceCode int
}

func (m *Movie) Movie(title string, priceCode int) {
	m.title = title
	m.priceCode = priceCode
}

func (m *Movie) GetPriceCode() int {
	return m.priceCode
}

func (m *Movie) SetPriceCode(arg int) {
	m.priceCode = arg
}

func (m *Movie) GetTitle() string {
	return m.title
}
