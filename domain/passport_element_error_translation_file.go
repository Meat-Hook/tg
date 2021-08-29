package domain

// PassportElementErrorTranslationFile Represents an issue with one of the files that constitute the translation of a document.
// The error is considered resolved when the file changes.
type PassportElementErrorTranslationFile struct {
	// Error source, must be translation_file.
	Source string `json:"source"`
	// Type of element of the user's Telegram Passport which has the issue, one of “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”.
	Type string `json:"type"`
	// Base64-encoded file hash.
	FileHash string `json:"file_hash"`
	// Error message.
	Message string `json:"message"`
}
