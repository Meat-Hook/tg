package domain

// InlineQueryResultLocation Represents a location on a map.
// By default, the location will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the location.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultLocation struct {
	// Url of the thumbnail for the result.
	//
	// Optional.
	ThumbUrl string `json:"thumb_url"`
	// Period in seconds for which the location can be updated, should be between 60 and 86400.
	//
	// Optional.
	LivePeriod int `json:"live_period"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Location latitude in degrees.
	Latitude float64 `json:"latitude"`
	// Location title.
	Title string `json:"title"`
	// The radius of uncertainty for the location, measured in meters; 0-1500.
	//
	// Optional.
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	// Thumbnail width.
	//
	// Optional.
	ThumbWidth int `json:"thumb_width"`
	// Type of the result, must be location.
	Type string `json:"type"`
	// Unique identifier for this result, 1-64 Bytes.
	Id string `json:"id"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters.
	// Must be between 1 and 100000 if specified.
	//
	// Optional.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
	// Content of the message to be sent instead of the location.
	//
	// Optional.
	InputMessageContent InputMessageContent `json:"input_message_content"`
	// Thumbnail height.
	//
	// Optional.
	ThumbHeight int `json:"thumb_height"`
	// Location longitude in degrees.
	Longitude float64 `json:"longitude"`
	// For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	//
	// Optional.
	Heading int `json:"heading"`
}
