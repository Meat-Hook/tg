package domain

// VoiceChatEnded This object represents a service message about a voice chat ended in the chat.
type VoiceChatEnded struct {
	// Voice chat duration; in seconds.
	Duration int `json:"duration"`
}
