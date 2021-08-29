package domain

// InlineQueryResultContact Represents a contact with a phone number.
// By default, this contact will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the contact.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultContact struct {
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Thumbnail width.
	//
	// Optional.
	ThumbWidth int `json:"thumb_width"`
	// Thumbnail height.
	//
	// Optional.
	ThumbHeight int `json:"thumb_height"`
	// Type of the result, must be contact.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes.
	Id string `json:"id"`
	// Contact's last name.
	//
	// Optional.
	LastName string `json:"last_name"`
	// Content of the message to be sent instead of the contact.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Url of the thumbnail for the result.
	//
	// Optional.
	ThumbUrl string `json:"thumb_url"`
	// Contact's phone number.
	PhoneNumber string `json:"phone_number"`
	// Contact's first name.
	FirstName string `json:"first_name"`
	// Additional data about the contact in the form of a vCard, 0-2048 bytes.
	//
	// Optional.
	Vcard string `json:"vcard"`
}
