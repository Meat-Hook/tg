package domain

// Invoice This object contains basic information about an invoice.
type Invoice struct {
	// Unique bot deep-linking parameter that can be used to generate this invoice.
	StartParameter string `json:"start_parameter"`
	// Three-letter ISO 4217 currency code.
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.
	// 45 pass amount = 145.
	// See the exp parameter in currencies.
	// json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Product name.
	Title string `json:"title"`
	// Product description.
	Description string `json:"description"`
}
