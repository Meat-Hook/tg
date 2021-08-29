package domain

// Animation This object represents an animation file (GIF or H.
// 264/MPEG-4 AVC video without sound).
type Animation struct {
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Video height as defined by sender.
	Height int `json:"height"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
	// Original animation filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name"`
	// MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type"`
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Video width as defined by sender.
	Width int `json:"width"`
	// Duration of the video in seconds as defined by sender.
	Duration int `json:"duration"`
	// Animation thumbnail as defined by sender.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
}
