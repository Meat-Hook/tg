package domain

// InputMediaVideo Represents a video to be sent.
type InputMediaVideo struct {
	// Caption of the video to be sent, 0-1024 characters after entities parsing.
	//
	// Optional.
	Caption string `json:"caption"`
	// Mode for parsing entities in the video caption.
	// See formatting options for more details.
	//
	// Optional.
	ParseMode string `json:"parse_mode"`
	// Video width.
	//
	// Optional.
	Width int `json:"width"`
	// Video height.
	//
	// Optional.
	Height int `json:"height"`
	// Video duration.
	//
	// Optional.
	Duration int `json:"duration"`
	// Pass True, if the uploaded video is suitable for streaming.
	//
	// Optional.
	SupportsStreaming bool `json:"supports_streaming"`
	// Type of the result, must be video.
	Type string `json:"type"`
	// File to send.
	// Pass a file_id to send a file that exists on the Telegram servers (recommended), pass an HTTP URL for Telegram to get a file from the Internet, or pass “attach://<file_attach_name>” to upload a new one using multipart/form-data under <file_attach_name> name.
	// More info on Sending Files ».
	Media string `json:"media"`
	// Thumbnail of the file sent; can be ignored if thumbnail generation for the file is supported server-side.
	// The thumbnail should be in JPEG format and less than 200 kB in size.
	// A thumbnail's width and height should not exceed 320.
	// Ignored if the file is not uploaded using multipart/form-data.
	// Thumbnails can't be reused and can be only uploaded as a new file, so you can pass “attach://<file_attach_name>” if the thumbnail was uploaded using multipart/form-data under <file_attach_name>.
	// More info on Sending Files ».
	//
	// Optional.
	Thumb string `json:"thumb"`
	// List of special entities that appear in the caption, which can be specified instead of parse_mode.
	//
	// Optional.
	CaptionEntities []MessageEntity `json:"caption_entities"`
}
