package domain

// InputTextMessageContent Represents the content of a text message to be sent as the result of an inline query.
type InputTextMessageContent struct {
	// Disables link previews for links in the scent message.
	//
	// Optional.
	DisableWebPagePreview bool `json:"disable_web_page_preview"`
	// Text of the message to be sent, 1-4096 characters.
	MessageText string `json:"message_text"`
	// Mode for parsing entities in the message text.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in message text, which can be specified instead of parse_mode.
	//
	// Optional.
	Entities []MessageEntity `json:"entities"`
}

func (i InputTextMessageContent) isInputMessageContent() {}
