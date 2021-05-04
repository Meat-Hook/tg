package domain

// KeyboardButton this object represents one button of the reply keyboard.
// For simple text buttons String can be used instead of this object to specify text of the button.
// Optional fields request_contact, request_location, and request_poll are mutually exclusive.
type KeyboardButton struct {
	// Text of the button.
	// If none of the optional fields are used, it will be sent as a message when the button is pressed.
	Text string `json:"text"`
	// RequestContact if True, the user's phone number will be sent as a contact when the button is pressed.
	// Available in private chats only.
	RequestContact bool `json:"request_contact"`
	// RequestLocation if True, the user's current location will be sent when the button is pressed.
	// Available in private chats only.
	RequestLocation bool `json:"request_location"`
	// RequestPoll if specified, the user will be asked to create a poll and send it to the bot when the button is pressed.
	// Available in private chats only
	RequestPoll *KeyboardButtonPollType `json:"request_poll,omitempty"`
}
