package domain

// InlineQueryResultGif Represents a link to an animated GIF file.
// By default, this animated GIF file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultGif struct {
	// Height of the GIF.
	//
	// Optional.
	GifHeight int `json:"gif_height"`
	// Duration of the GIF.
	//
	// Optional.
	GifDuration int `json:"gif_duration"`
	// Title for the result.
	//
	// Optional.
	Title string `json:"title"`
	// Content of the message to be sent instead of the GIF animation.
	//
	// Optional.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// Width of the GIF.
	//
	// Optional.
	GifWidth int `json:"gif_width"`
	// Caption of the GIF file to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Type of the result, must be gif.
	Type string `json:"type"`
	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result.
	ThumbUrl string `json:"thumb_url"`
	// Mode for parsing entities in the caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// A valid URL for the GIF file.
	// File size must not exceed 1MB.
	GifUrl string `json:"gif_url"`
	// MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”.
	// Defaults to “image/jpeg”.
	//
	// Optional.
	ThumbMimeType string `json:"thumb_mime_type"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
