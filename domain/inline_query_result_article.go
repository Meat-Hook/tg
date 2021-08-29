package domain

// InlineQueryResultArticle Represents a link to an article or web page.
type InlineQueryResultArticle struct {
	// Url of the thumbnail for the result.
	//
	// Optional.
	ThumbUrl string `json:"thumb_url"`
	// Type of the result, must be article.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes.
	Id string `json:"id"`
	// Title of the result.
	Title string `json:"title"`
	// Content of the message to be sent.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// URL of the result.
	//
	// Optional.
	Url string `json:"url"`
	// Pass True, if you don't want the URL to be shown in the message.
	//
	// Optional.
	HideUrl bool `json:"hide_url"`
	// Short description of the result.
	//
	// Optional.
	Description string `json:"description"`
	// Thumbnail width.
	//
	// Optional.
	ThumbWidth int `json:"thumb_width"`
	// Thumbnail height.
	//
	// Optional.
	ThumbHeight int `json:"thumb_height"`
}
