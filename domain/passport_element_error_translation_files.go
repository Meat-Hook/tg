package domain

// PassportElementErrorTranslationFiles Represents an issue with the translated version of a document.
// The error is considered resolved when a file with the document translation change.
type PassportElementErrorTranslationFiles struct {
	// Error source, must be translation_files.
	Source string `json:"source"`
	// Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”.
	Type string `json:"type"`
	// List of base64-encoded file hashes.
	FileHashes []string `json:"file_hashes"`
	// Error message.
	Message string `json:"message"`
}
