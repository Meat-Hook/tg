package domain

// PassportElementErrorReverseSide Represents an issue with the reverse side of a document.
// The error is considered resolved when the file with reverse side of the document changes.
type PassportElementErrorReverseSide struct {
	// Error source, must be reverse_side.
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “driver_license”, “identity_card”.
	Type string `json:"type"`
	// Base64-encoded hash of the file with the reverse side of the document.
	FileHash string `json:"file_hash"`
	// Error message.
	Message string `json:"message"`
}
