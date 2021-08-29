package domain

// File This object represents a file ready to be downloaded.
// The file can be downloaded via the link https://api.
// telegram.
// org/file/bot<token>/<file_path>.
// It is guaranteed that the link will be valid for at least 1 hour.
// When the link expires, a new one can be requested by calling getFile.
// Maximum file size to download is 20 MB.
type File struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// File size, if known.
	//
	// Optional.
	FileSize int `json:"file_size"`
	// File path.
	// Use https://api.
	// telegram.
	// org/file/bot<token>/<file_path> to get the file.
	//
	// Optional.
	FilePath string `json:"file_path"`
}
