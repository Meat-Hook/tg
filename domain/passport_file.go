package domain

// PassportFile This object represents a file uploaded to Telegram Passport.
// Currently, all Telegram Passport files are in JPEG format when decrypted and don't exceed 10MB.
type PassportFile struct {
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// File size.
	FileSize int `json:"file_size"`
	// Unix time when the file was uploaded.
	FileDate int `json:"file_date"`
}
