package domain

// MessageEntity This object represents one special entity in a text message.
// For example, hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Type of the entity.
	// Can be “mention” (@username), “hashtag” (#hashtag), “cashtag” ($USD), “bot_command” (/start@jobs_bot), “url” (https://telegram.
	// org), “email” (do-not-reply@telegram.
	// org), “phone_number” (+1-212-555-0123), “bold” (bold text), “italic” (italic text), “underline” (underlined text), “strikethrough” (strikethrough text), “code” (monowidth string), “pre” (monowidth block), “text_link” (for clickable text URLs), “text_mention” (for users without usernames).
	Type string `json:"type"`
	// Offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`
	// Length of the entity in UTF-16 code units.
	Length int `json:"length"`
	// For “text_link” only, url that will be opened after user taps on the text.
	//
	// Optional.
	Url string `json:"url"`
	// For “text_mention” only, the mentioned user.
	//
	// Optional.
	User *User `json:"user"`
	// For “pre” only, the programming language of the entity text.
	//
	// Optional.
	Language string `json:"language"`
}
