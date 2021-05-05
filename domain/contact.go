package domain

// Contact this object represents a phone contact.
type Contact struct {
	// PhoneNumber contact's phone number.
	PhoneNumber string `json:"phone_number"`
	// FirstName contact's first name.
	FirstName string `json:"first_name"`
	// LastName contact's last name.
	LastName string `json:"last_name"`
	// UserID contact's user identifier in Telegram.
	// This number may have more than 32 significant bits and some programming languages
	// may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision
	// float type are safe for storing this identifier.
	UserID string `json:"user_id"`
	// VCard additional data about the contact in the form of a vCard.
	VCard string `json:"vcard"`
}
