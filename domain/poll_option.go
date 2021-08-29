package domain

// PollOption This object contains information about one answer option in a poll.
type PollOption struct {
	// Option text, 1-100 characters.
	Text string `json:"text"`
	// Number of users that voted for this option.
	VoterCount int `json:"voter_count"`
}
