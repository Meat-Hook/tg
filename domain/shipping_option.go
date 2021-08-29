package domain

// ShippingOption This object represents one shipping option.
type ShippingOption struct {
	// List of price portions.
	Prices []LabeledPrice `json:"prices"`
	// Shipping option identifier.
	Id string `json:"id"`
	// Option title.
	Title string `json:"title"`
}
