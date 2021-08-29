package domain

// PreCheckoutQuery This object contains information about an incoming pre-checkout query.
// Telegram Passport.
// Telegram Passport is a unified authorization method for services that require personal identification.
// Users can upload their documents once, then instantly share their data with services that require real-world ID (finance, ICOs, etc.
// ).
// Please see the manual for details.
type PreCheckoutQuery struct {
	// Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`
	// Identifier of the shipping option chosen by the user.
	//
	// Optional.
	ShippingOptionId string `json:"shipping_option_id"`
	// Order info provided by the user.
	//
	// Optional.
	OrderInfo *OrderInfo `json:"order_info"`
	// Unique query identifier.
	Id string `json:"id"`
	// User who sent the query.
	From User `json:"from"`
	// Three-letter ISO 4217 currency code.
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.
	// 45 pass amount = 145.
	// See the exp parameter in currencies.
	// json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
}
