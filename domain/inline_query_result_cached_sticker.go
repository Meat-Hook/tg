package domain

// InlineQueryResultCachedSticker Represents a link to a sticker stored on the Telegram servers.
// By default, this sticker will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the sticker.
// Note: This will only work in Telegram versions released after 9 April, 2016 for static stickers and after 06 July, 2019 for animated stickers.
// Older clients will ignore them.
type InlineQueryResultCachedSticker struct {
	// Type of the result, must be a sticker.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 bytes.
	Id string `json:"id"`
	// A valid file identifier of the sticker.
	StickerFileId string `json:"sticker_file_id"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Content of the message to be sent instead of the sticker.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
}
