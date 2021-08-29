package domain

// InlineQueryResultCachedPhoto Represents a link to a photo stored on the Telegram servers.
// By default, this photo will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the photo.
type InlineQueryResultCachedPhoto struct {
	// Type of the result, must be a photo.
	Type string `json:"type"`
	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the photo caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
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
	// Content of the message to be sent instead of the photo.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// A valid file identifier of the photo.
	PhotoFileId string `json:"photo_file_id"`
	// Title for the result.
	//
	// Optional.
	Title string `json:"title"`
}
