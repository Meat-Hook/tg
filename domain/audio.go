package domain

// Audio This object represents an audio file to be treated as music by the Telegram clients.
type Audio struct {
	// Performer of the audio as defined by sender or by audio tags.
	//
	// Optional.
	Performer string `json:"performer"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Duration of the audio in seconds as defined by sender.
	Duration int `json:"duration"`
	// Title of the audio as defined by sender or by audio tags.
	//
	// Optional.
	Title string `json:"title"`
	// Original filename as defined by sender.
	//
	// Optional.
	FileName string `json:"file_name"`
	// MIME type of the file as defined by sender.
	//
	// Optional.
	MimeType string `json:"mime_type"`
	// Thumbnail of the album cover to which the music file belongs.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
}
