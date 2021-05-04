package domain

// Audio this object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// FileID identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// FileUniqueID unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Duration of the audio in seconds as defined by sender.
	Duration int64 `json:"duration"`
	// Performer of the audio as defined by sender or by audio tags.
	Performer string `json:"performer"`
	// Title of the audio as defined by sender or by audio tags.
	Title string `json:"title"`
	// FileName original filename as defined by sender.
	FileName string `json:"file_name"`
	// MimeType MIME type of the file as defined by sender.
	MimeType string `json:"mime_type"`
	// FileSize file size.
	FileSize string `json:"file_size"`
}
