package domain

// InlineQueryResultMpeg4Gif Represents a link to a video animation (H.
// 264/MPEG-4 AVC video without sound).
// By default, this animated MPEG-4 file will be sent by the user with optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultMpeg4Gif struct {
	// Type of the result, must be mpeg4_gif.
	Type string `json:"type"`
	// MIME type of the thumbnail, must be one of “image/jpeg”, “image/gif”, or “video/mp4”.
	// Defaults to “image/jpeg”.
	//
	// Optional.
	ThumbMimeType string `json:"thumb_mime_type"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Video width.
	//
	// Optional.
	Mpeg4Width int `json:"mpeg4_width"`
	// Video height.
	//
	// Optional.
	Mpeg4Height int `json:"mpeg4_height"`
	// URL of the static (JPEG or GIF) or animated (MPEG4) thumbnail for the result.
	ThumbUrl string `json:"thumb_url"`
	// Title for the result.
	//
	// Optional.
	Title string `json:"title"`
	// Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Content of the message to be sent instead of the video animation.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// A valid URL for the MP4 file.
	// File size must not exceed 1MB.
	Mpeg4Url string `json:"mpeg4_url"`
	// Video duration.
	//
	// Optional.
	Mpeg4Duration int `json:"mpeg_4_duration"`
	// Mode for parsing entities in the caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
