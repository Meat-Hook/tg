package domain

// InputContactMessageContent Represents the content of a contact message to be sent as the result of an inline query.
type InputContactMessageContent struct {
	// Contact's phone number.
	PhoneNumber string `json:"phone_number"`
	// Contact's first name.
	FirstName string `json:"first_name"`
	// Contact's last name.
	//
	// Optional.
	LastName string `json:"last_name"`
	// Additional data about the contact in the form of a vCard, 0-2048 bytes.
	//
	// Optional.
	Vcard string `json:"vcard"`
}

func (i InputContactMessageContent) isInputMessageContent() {}
