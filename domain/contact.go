package domain

// Contact This object represents a phone contact.
type Contact struct {
	// Contact's phone number.
	PhoneNumber string `json:"phone_number"`
	// Contact's first name.
	FirstName string `json:"first_name"`
	// Contact's last name.
	//
	// Optional.
	LastName string `json:"last_name"`
	// Contact's user identifier in Telegram.
	// This number may have more than 32 significant bits and some programming languages may have difficulty/silent defects in interpreting it.
	// But it has at most 52 significant bits, so a 64-bit integer or double-precision float type are safe for storing this identifier.
	//
	// Optional.
	UserId int `json:"user_id"`
	// Additional data about the contact in the form of a vCard.
	//
	// Optional.
	Vcard string `json:"vcard"`
}
