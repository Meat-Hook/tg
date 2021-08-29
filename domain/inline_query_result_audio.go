package domain

// InlineQueryResultAudio Represents a link to an MP3 audio file.
// By default, this audio file will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the audio.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultAudio struct {
	// Performer.
	//
	// Optional.
	Performer string `json:"performer"`
	// Audio duration in seconds.
	//
	// Optional.
	AudioDuration int `json:"audio_duration"`
	// A valid URL for the audio file.
	AudioUrl string `json:"audio_url"`
	// Title.
	Title string `json:"title"`
	// Caption, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Content of the message to be sent instead of the audio.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Type of the result, must be audio.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Mode for parsing entities in the audio caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
