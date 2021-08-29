package domain

// ShippingQuery This object contains information about an incoming shipping query.
type ShippingQuery struct {
	// User who sent the query.
	From User `json:"from"`
	// Bot specified invoice payload.
	InvoicePayload string `json:"invoice_payload"`
	// User specified shipping address.
	ShippingAddress ShippingAddress `json:"shipping_address"`
	// Unique query identifier.
	Id string `json:"id"`
}
