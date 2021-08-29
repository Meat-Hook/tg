package domain

// VoiceChatScheduled This object represents a service message about a voice chat scheduled in the chat.
type VoiceChatScheduled struct {
	// Point in time (Unix timestamp) when the voice chat is supposed to be started by a chat administrator.
	StartDate int `json:"start_date"`
}
