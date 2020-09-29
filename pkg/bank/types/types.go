package types

// Money represents a monetary amount in minimum units (cents, kopecks, diramas, etc.).
type Money int64

// Categoty represents the category in which the payment was made (cars, pharmacies, restaurants, etc.).
type Categoty string

// Payment presents information about the payment.
type Payment struct {
	ID       int
	Amount   Money
	Categoty Categoty
}
