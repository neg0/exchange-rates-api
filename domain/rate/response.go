package rate

// According to DDD, client should not know about our main Domain logic
// That's why I rather marshall JSON response in a separate struct
// We don't want to expose of domain logic to the client
type Response struct {
	Rates map[string]float64 `json:"rates"`
	Base  string             `json:"base"`
	Date  string             `json:"date"`
}
