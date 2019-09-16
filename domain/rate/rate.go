package rate

const DateFormat = "2006-01-02"

type Rate interface {
	Date() string
	Currency() string
	BassCurrency() string
	Value() float64
}
