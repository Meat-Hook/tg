package domain

// InlineQueryResultPhoto Represents a link to a photo.
// By default, this photo will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultPhoto struct {
	// Title for the result.
	//
	// Optional.
	Title string `json:"title"`
	// Short description of the result.
	//
	// Optional.
	Description string `json:"description"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// URL of the thumbnail for the photo.
	ThumbUrl string `json:"thumb_url"`
	// Width of the photo.
	//
	// Optional.
	PhotoWidth int `json:"photo_width"`
	// A valid URL of the photo.
	// Photo must be in jpeg format.
	// Photo size must not exceed 5MB.
	PhotoUrl string `json:"photo_url"`
	// Height of the photo.
	//
	// Optional.
	PhotoHeight int `json:"photo_height"`
	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the photo caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Content of the message to be sent instead of the photo.
	//
	// Optional.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// Type of the result, must be a photo.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
}
