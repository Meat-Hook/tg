package domain

// PassportElementErrorFrontSide Represents an issue with the front side of a document.
// The error is considered resolved when the file with the front side of the document changes.
type PassportElementErrorFrontSide struct {
	// The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”.
	Type string `json:"type"`
	// Base64-encoded hash of the file with the front side of the document.
	FileHash string `json:"file_hash"`
	// Error message.
	Message string `json:"message"`
	// Error source, must be front_side.
	Source string `json:"source"`
}
