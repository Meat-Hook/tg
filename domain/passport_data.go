package domain

// PassportData Contains information about Telegram Passport data shared with the bot by the user.
type PassportData struct {
	// Encrypted credentials required to decrypt the data.
	Credentials EncryptedCredentials `json:"credentials"`
	// Array with information about documents and other Telegram Passport elements that was shared with the bot.
	Data []EncryptedPassportElement `json:"data"`
}
