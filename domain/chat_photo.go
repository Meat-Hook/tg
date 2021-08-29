package domain

// ChatPhoto This object represents a chat photo.
type ChatPhoto struct {
	// Unique file identifier of small (160x160) chat photo, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	SmallFileUniqueId string `json:"small_file_unique_id"`
	// File identifier of big (640x640) chat photo.
	// This file_id can be used only for photo download and only for as long as the photo is not changed.
	BigFileId string `json:"big_file_id"`
	// Unique file identifier of big (640x640) chat photo, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	BigFileUniqueId string `json:"big_file_unique_id"`
	// File identifier of small (160x160) chat photo.
	// This file_id can be used only for photo download and only for as long as the photo is not changed.
	SmallFileId string `json:"small_file_id"`
}
