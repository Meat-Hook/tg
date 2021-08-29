package domain

// PhotoSize This object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Photo width.
	Width int `json:"width"`
	// Photo height.
	Height int `json:"height"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
}
