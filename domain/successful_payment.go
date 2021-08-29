package domain

// SuccessfulPayment This object contains basic information about a successful payment.
type SuccessfulPayment struct {
	// Identifier of the shipping option chosen by the user.
	//
	// Optional.
	ShippingOptionId string `json:"shipping_option_id"`
	// Order info provided by the user.
	//
	// Optional.
	OrderInfo *OrderInfo `json:"order_info"`
	// Telegram payment identifier.
	TelegramPaymentChargeId string `json:"telegram_payment_charge_id"`
	// Provider payment identifier.
	ProviderPaymentChargeId string `json:"provider_payment_charge_id"`
	// Three-letter ISO 4217 currency code.
	Currency string `json:"currency"`
	// Total price in the smallest units of the currency (integer, not float/double).
	// For example, for a price of US$ 1.
	// 45 pass amount = 145.
	// See the exp parameter in currencies.
	// json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	TotalAmount int `json:"total_amount"`
	// Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`
}
