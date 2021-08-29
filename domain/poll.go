package domain

// Poll This object contains information about a poll.
type Poll struct {
	// 0-based identifier of the correct answer option.
	// Available only for polls in the quiz mode, which are closed, or was sent (not forwarded) by the bot or to the private chat with the bot.
	//
	// Optional.
	CorrectOptionId int `json:"correct_option_id"`
	// Unique poll identifier.
	Id string `json:"id"`
	// Poll question, 1-300 characters.
	Question string `json:"question"`
	// List of poll options.
	Options []PollOption `json:"options"`
	// Total number of users that voted in the poll.
	TotalVoterCount int `json:"total_voter_count"`
	// True, if the poll is closed.
	IsClosed bool `json:"is_closed"`
	// True, if the poll is anonymous.
	IsAnonymous bool `json:"is_anonymous"`
	// True, if the poll allows multiple answers.
	AllowsMultipleAnswers bool `json:"allows_multiple_answers"`
	// Text that is shown when a user chooses an incorrect answer or taps on the lamp icon in a quiz-style poll, 0-200 characters.
	//
	// Optional.
	Explanation string `json:"explanation"`
	// Point in time (Unix timestamp) when the poll will be automatically closed.
	//
	// Optional.
	CloseDate int `json:"close_date"`
	// Poll type, currently can be “regular” or “quiz”.
	Type string `json:"type"`
	// Special entities like usernames, URLs, bot commands, etc.
	// that appear in the explanation.
	//
	// Optional.
	ExplanationEntities []MessageEntity `json:"explanation_entities"`
	// Amount of time in seconds the poll will be active after creation.
	//
	// Optional.
	OpenPeriod int `json:"open_period"`
}
