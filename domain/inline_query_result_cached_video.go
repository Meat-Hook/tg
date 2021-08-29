package domain

// InlineQueryResultCachedVideo Represents a link to a video file stored on the Telegram servers.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
type InlineQueryResultCachedVideo struct {
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// A valid file identifier for the video file.
	VideoFileId string `json:"video_file_id"`
	// Short description of the result.
	//
	// Optional.
	Description string `json:"description"`
	// Mode for parsing entities in the video caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Caption of the video to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Content of the message to be sent instead of the video.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Type of the result, must be video.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Title for the result.
	Title string `json:"title"`
}
