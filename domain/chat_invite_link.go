package domain

// ChatInviteLink represents an invite link for a chat.
type ChatInviteLink struct {
	// InviteLink the invite link.
	// If the link was created by another chat administrator,
	// then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`
	// Creator create of the link.
	Creator User `json:"creator"`
	// IsPrimary true, if the link is primary.
	IsPrimary bool `json:"is_primary"`
	// IsRevoked true, if the link is revoked.
	IsRevoked bool `json:"is_revoked"`
	// ExpireDate point in time (Unix timestamp) when the link will expire or has been expired.
	ExpireDate int64 `json:"expire_date"`
	// MemberLimit maximum number of users that can be members of the chat
	// simultaneously after joining the chat via this invite link; 1-99999.
	MemberLimit int `json:"member_limit"`
}
