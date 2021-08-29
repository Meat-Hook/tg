package domain

// VoiceChatParticipantsInvited This object represents a service message about new members invited to a voice chat.
type VoiceChatParticipantsInvited struct {
	// New members that were invited to the voice chat.
	//
	// Optional.
	Users []User `json:"users"`
}
