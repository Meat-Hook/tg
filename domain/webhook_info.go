package domain

// WebhookInfo Contains information about the current status of a webhook.
// Available types.
// All types used in the Bot API responses are represented as JSON-objects.
// It is safe to use 32-bit signed integers for storing all Integer fields unless otherwise noted.
// Optional fields may be not returned when irrelevant.
type WebhookInfo struct {
	// Unix time for the most recent error that happened when trying to deliver an update via webhook.
	//
	// Optional.
	LastErrorDate int `json:"last_error_date"`
	// Error message in human-readable format for the most recent error that happened when trying to deliver an update via webhook.
	//
	// Optional.
	LastErrorMessage string `json:"last_error_message"`
	// Maximum allowed number of simultaneous HTTPS connections to the webhook for update delivery.
	//
	// Optional.
	MaxConnections int `json:"max_connections"`
	// A list of update types the bot is subscribed to.
	// Defaults to all update types except chat_member.
	//
	// Optional.
	AllowedUpdates []string `json:"allowed_updates"`
	// Webhook URL, may be empty if webhook is not set up.
	Url string `json:"url"`
	// True, if a custom certificate was provided for webhook certificate checks.
	HasCustomCertificate bool `json:"has_custom_certificate"`
	// Number of updates awaiting delivery.
	PendingUpdateCount int `json:"pending_update_count"`
	// Currently used webhook IP address.
	//
	// Optional.
	IpAddress string `json:"ip_address"`
}
