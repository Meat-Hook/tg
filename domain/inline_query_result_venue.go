package domain

// InlineQueryResultVenue Represents a venue.
// By default, the venue will be sent by the user.
// Alternatively, you can use input_message_content to send a message with the specified content instead of the venue.
// Note: This will only work in Telegram versions released after 9 April, 2016.
// Older clients will ignore them.
type InlineQueryResultVenue struct {
	// Type of the result, must be venue.
	Type string `json:"type"`
	// Address of the venue.
	Address string `json:"address"`
	// Google Places type of the venue.
	// (See supported types).
	//
	// Optional.
	GooglePlaceType string `json:"google_place_type"`
	// Inline keyboard attached to the message.
	//
	// Optional.
	ReplyMarkup *InlineKeyboardMarkup `json:"reply_markup"`
	// Unique identifier for this result, 1-64 Bytes.
	Id string `json:"id"`
	// Latitude of the venue location in degrees.
	Latitude float64 `json:"latitude"`
	// Longitude of the venue location in degrees.
	Longitude float64 `json:"longitude"`
	// Google Places identifier of the venue.
	//
	// Optional.
	GooglePlaceId string `json:"google_place_id"`
	// Content of the message to be sent instead of the venue.
	//
	// Optional.
	InputMessageContent *InputMessageContent `json:"input_message_content"`
	// Url of the thumbnail for the result.
	//
	// Optional.
	ThumbUrl string `json:"thumb_url"`
	// Thumbnail width.
	//
	// Optional.
	ThumbWidth int `json:"thumb_width"`
	// Title of the venue.
	Title string `json:"title"`
	// Foursquare identifier of the venue if known.
	//
	// Optional.
	FoursquareId string `json:"foursquare_id"`
	// Foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”).
	//
	// Optional.
	FoursquareType string `json:"foursquare_type"`
	// Thumbnail height.
	//
	// Optional.
	ThumbHeight int `json:"thumb_height"`
}
