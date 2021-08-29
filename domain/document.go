package domain

// Document This object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Document thumbnail as defined by sender.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
	// Original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name"`
	// MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type"`
}
