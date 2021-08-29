package domain

// InlineQueryResultCachedMpeg4Gif Represents a link to a video animation (H.
// 264/MPEG-4 AVC video without sound) stored on the Telegram servers.
// By default, this animated MPEG-4 file will be sent by the user with an optional caption.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the animation.
type InlineQueryResultCachedMpeg4Gif struct {
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// A valid file identifier for the MP4 file.
	Mpeg4FileId string `json:"mpeg_4_file_id"`
	// Mode for parsing entities in the caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Content of the message to be sent instead of the video animation.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Type of the result, must be mpeg4_gif.
	Type string `json:"type"`
	// Title for the result.
	//
	// Optional.
	Title string `json:"title"`
	// Caption of the MPEG-4 file to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
