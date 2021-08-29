package domain

// PassportElementErrorFile Represents an issue with a document scan.
// The error is considered resolved when the file with the document scan changes.
type PassportElementErrorFile struct {
	// Error source, must be a file.
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”.
	Type string `json:"type"`
	// Base64-encoded file hash.
	FileHash string `json:"file_hash"`
	// Error message.
	Message string `json:"message"`
}
