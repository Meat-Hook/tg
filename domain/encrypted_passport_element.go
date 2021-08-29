package domain

// EncryptedPassportElement Contains information about documents or other Telegram Passport elements shared with the bot by the user.
type EncryptedPassportElement struct {
	// Base64-encoded encrypted Telegram Passport element data provided by the user, available for “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport” and “address” types.
	// Can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Data string `json:"data"`
	// User's verified phone number, available only for “phone_number” type.
	//
	// Optional.
	PhoneNumber string `json:"phone_number"`
	// User's verified email address, available only for “email” type.
	//
	// Optional.
	Email string `json:"email"`
	// Array of encrypted files with documents provided by the user, available for “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Files []PassportFile `json:"files"`
	// Encrypted file with the front side of the document, provided by the user.
	// Available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	FrontSide *PassportFile `json:"front_side"`
	// Encrypted file with the selfie of the user holding a document, provided by the user; available for “passport”, “driver_license”, “identity_card” and “internal_passport”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Selfie *PassportFile `json:"selfie"`
	// Array of encrypted files with translated versions of documents provided by the user.
	// Available if requested for “passport”, “driver_license”, “identity_card”, “internal_passport”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration” and “temporary_registration” types.
	// Files can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	Translation []PassportFile `json:"translation"`
	// Element type.
	// One of “personal_details”, “passport”, “driver_license”, “identity_card”, “internal_passport”, “address”, “utility_bill”, “bank_statement”, “rental_agreement”, “passport_registration”, “temporary_registration”, “phone_number”, “email”.
	Type string `json:"type"`
	// Base64-encoded element hash for using in PassportElementErrorUnspecified.
	Hash string `json:"hash"`
	// Encrypted file with the reverse side of the document, provided by the user.
	// Available for “driver_license” and “identity_card”.
	// The file can be decrypted and verified using the accompanying EncryptedCredentials.
	//
	// Optional.
	ReverseSide *PassportFile `json:"reverse_side"`
}
