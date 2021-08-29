package domain

// Video This object represents a video file.
type Video struct {
	// Mime type of file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video width as defined by sender.
	Width int `json:"width"`
	// Video height as defined by sender.
	Height int `json:"height"`
	// Original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name"`
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`
	// Video thumbnail.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
}
