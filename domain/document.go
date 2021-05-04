package domain

// Document this object represents a general file (as opposed to photos, voice messages and audio files).
type Document struct {
	// FileID identifier for this file, which can be used to download or reuse the file
	FileID string `json:"file_id"`
	// FileUniqueID unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Thumb document thumbnail as defined by sender.
	Thumb *PhotoSize `json:"thumb,omitempty"`
	// FileName original filename as defined by sender.
	FileName string `json:"file_name"`
	// MimeType MIME type of the file as defined by sender.
	MimeType string `json:"mime_type"`
	// FileSize file size.
	FileSize string `json:"file_size"`
}
