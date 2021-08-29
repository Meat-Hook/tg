package domain

// InputMediaPhoto Represents a photo to be sent.
type InputMediaPhoto struct {
	// Type of the result, must be a photo.
	Type string `json:"type"`
	// File to send.
	// Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// More info on Sending Files ».
	Media string `json:"media"`
	// Caption of the photo to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the photo caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
}
