package domain

// InlineQueryResultCachedDocument Represents a link to a file stored on the Telegram servers.
// By default, this file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the file.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultCachedDocument struct {
	// Title for the result.
	Title string `json:"title"`
	// Short description of the result.
	//
	// Optional.
	Description string `json:"description"`
	// Caption of the document to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Type of the result, must be a document.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Content of the message to be sent instead of the file.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// A valid file identifier for the file.
	DocumentFileId string `json:"document_file_id"`
	// Mode for parsing entities in the document caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
}
