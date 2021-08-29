package domain

// InlineQueryResultCachedGif Represents a link to an animated GIF file stored on the Telegram servers.
// By default, this animated GIF file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with specified content instead of the animation.
type InlineQueryResultCachedGif struct {
	// Type of the result, must be gif.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Caption of the GIF file to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Content of the message to be sent instead of the GIF animation.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// A valid file identifier for the GIF file.
	GifFileId string `json:"gif_file_id"`
	// Title for the result.
	//
	// Optional.
	Title string `json:"title"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
