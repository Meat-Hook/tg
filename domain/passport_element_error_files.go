package domain

// PassportElementErrorFiles Represents an issue with a list of scans.
// The error is considered resolved when the list of files containing the scans changes.
type PassportElementErrorFiles struct {
	// List of base64-encoded file hashes.
	FileHashes []string `json:"file_hashes"`
	// Error message.
	Message string `json:"message"`
	// Error source, must be files.
	Source string `json:"source"`
	// The section of the user's Telegram Passport which has the issue, one of “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”.
	Type string `json:"type"`
}
