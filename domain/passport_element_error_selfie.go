package domain

// PassportElementErrorSelfie Represents an issue with the selfie with a document.
// The error is considered resolved when the file with the selfie changes.
type PassportElementErrorSelfie struct {
	// Error source, must be selfie.
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”.
	Type string `json:"type"`
	// Base64-encoded hash of the file with the selfie.
	FileHash string `json:"file_hash"`
	// Error message.
	Message string `json:"message"`
}
