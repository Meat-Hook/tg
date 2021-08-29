package domain

// MessageAutoDeleteTimerChanged This object represents a service message about a change in auto-delete timer settings.
type MessageAutoDeleteTimerChanged struct {
	// New auto-delete time for messages in the chat.
	MessageAutoDeleteTime int `json:"message_auto_delete_time"`
}
