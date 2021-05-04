package domain

// Video this object represents a video file.
type Video struct {
	// FileID identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// FileUniqueID unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width video width as defined by sender.
	Width int `json:"width"`
	// Height height as defined by sender.
	Height int `json:"height"`
	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`
	// Thumb video thumbnail.
	Thumb *PhotoSize `json:"thumb"`
	// FileName original filename as defined by sender.
	FileName string `json:"file_name"`
	// MimeType mime type of a file as defined by sender.
	MimeType string `json:"mime_type"`
	// FileSize file size.
	FileSize int `json:"file_size"`
}
