package domain

// VideoNote This object represents a video message (available in Telegram apps as of v.4.0).
type VideoNote struct {
	// Video thumbnail.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width and height (diameter of the video message) as defined by sender.
	Length int `json:"length"`
	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`
}
