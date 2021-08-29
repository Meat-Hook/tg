package domain

// ChatInviteLink Represents an invite link for a chat.
type ChatInviteLink struct {
	// The invite link.
	// If the link was created by another chat administrator, then the second part of the link will be replaced with “…”.
	InviteLink string `json:"invite_link"`
	// Creator of the link.
	Creator User `json:"creator"`
	// True, if the link is primary.
	IsPrimary bool `json:"is_primary"`
	// True, if the link is revoked.
	IsRevoked bool `json:"is_revoked"`
	// Point in time (Unix timestamp) when the link will expire or has been expired.
	//
	// Optional.
	ExpireDate int `json:"expire_date"`
	// Maximum number of users that can be members of the chat simultaneously after joining the chat via this invite link; 1-99999.
	//
	// Optional.
	MemberLimit int `json:"member_limit"`
}
