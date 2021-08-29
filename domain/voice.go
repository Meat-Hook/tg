package domain

// Voice This object represents a voice note.
type Voice struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`
	// MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
}
