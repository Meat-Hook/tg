package domain

// InlineQueryResultDocument Represents a link to a file.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Currently, only.
// PDF and.
// ZIP files can be sent using this method.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultDocument struct {
	// URL of the thumbnail (jpeg only) for the file.
	//
	// Optional.
	ThumbUrl string `json:"thumb_url"`
	// Thumbnail width.
	//
	// Optional.
	ThumbWidth int `json:"thumb_width"`
	// Mode for parsing entities in the document caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Short description of the result.
	//
	// Optional.
	Description string `json:"description"`
	// Mime type of the content of the file, either “application/pdf” or “application/zip”.
	MimeType string `json:"mime_type"`
	// Content of the message to be sent instead of the file.
	//
	// Optional.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// Thumbnail height.
	//
	// Optional.
	ThumbHeight int `json:"thumb_height"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Type of the result, must be a document.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Title for the result.
	Title string `json:"title"`
	// Caption of the document to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// A valid URL for the file.
	DocumentUrl string `json:"document_url"`
}
