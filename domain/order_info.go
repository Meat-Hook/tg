package domain

// OrderInfo This object represents information about an order.
type OrderInfo struct {
	// User name.
	//
	// Optional.
	Name string `json:"name"`
	// User's phone number.
	//
	// Optional.
	PhoneNumber string `json:"phone_number"`
	// User email.
	//
	// Optional.
	Email string `json:"email"`
	// User shipping address.
	//
	// Optional.
	ShippingAddress *ShippingAddress `json:"shipping_address"`
}
