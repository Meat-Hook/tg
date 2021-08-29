package domain

// InlineQueryResultVoice Represents a link to a voice recording in an.
// OGG container encoded with OPUS.
// By default, this voice recording will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the the voice message.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultVoice struct {
	// Recording duration in seconds.
	//
	// Optional.
	VoiceDuration int `json:"voice_duration"`
	// Content of the message to be sent instead of the voice recording.
	//
	// Optional.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// A valid URL for the voice recording.
	VoiceUrl string `json:"voice_url"`
	// Caption, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the voice message caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
	// Type of the result, must be voice.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// Recording title.
	Title string `json:"title"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
}
