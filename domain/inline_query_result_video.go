package domain

// InlineQueryResultVideo Represents a link to a page containing an embedded video player or a video file.
// By default, this video file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the video.
// If an InlineQueryResultVideo message contains an embedded video (e.g. YouTube), you must replace its content using input_message_content.
type InlineQueryResultVideo struct {
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Title for the result.
	Title string `json:"title"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Type of the result, must be video.
	Type string `json:"type"`
	// Mime type of the content of video url, “text/html” or “video/mp4”.
	MimeType string `json:"mime_type"`
	// Video width.
	//
	// Optional.
	VideoWidth int `json:"video_width"`
	// Content of the message to be sent instead of the video.
	// This field is required if InlineQueryResultVideo is used to send an HTML-page as a result (e.g., a YouTube video).
	//
	// Optional.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// A valid URL for the embedded video player or video file.
	VideoUrl string
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Video height.
	//
	// Optional.
	VideoHeight int `json:"video_height"`
	// Video duration in seconds.
	//
	// Optional.
	VideoDuration int `json:"video_duration"`
	// URL of the thumbnail (jpeg only) for the video.
	ThumbUrl string `json:"thumb_url"`
	// Caption of the video to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the video caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Short description of the result.
	//
	// Optional.
	Description string `json:"description"`
}
