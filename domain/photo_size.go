package domain

// PhotoSize this object represents one size of a photo or a file / sticker thumbnail.
type PhotoSize struct {
	// FileID identifier for this file, which can be used to download or reuse the file.
	FileID string `json:"file_id"`
	// FileUniqueID unique identifier for this file,
	// which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueID string `json:"file_unique_id"`
	// Width photo width.
	Width int `json:"width"`
	// Height photo height.
	Height int `json:"height"`
	// FileSize file size.
	FileSize int `json:"file_size"`
}
