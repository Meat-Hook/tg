package domain

// ChosenInlineResult Represents a result of an inline query that was chosen by the user and sent to their chat partner.
// Note: It is necessary to enable inline feedback via @Botfather in order to receive these objects in updates.
// Payments.
// Your bot can accept payments from Telegram users.
// Please see the introduction to payments for more details on the process and how to set up payments for your bot.
// Please note that users will need Telegram v.4.0 or higher to use payments (released on May 18, 2017).
type ChosenInlineResult struct {
	// The unique identifier for the result that was chosen.
	ResultId string `json:"result_id"`
	// The user that chose the result.
	From User `json:"from"`
	// Sender location, only for bots that require user location.
	//
	// Optional.
	Location *Location `json:"location"`
	// Identifier of the sent inline message.
	// Available only if there is an inline keyboard attached to the message.
	// Will be also received in callback queries and can be used to edit the message.
	//
	// Optional.
	InlineMessageId string `json:"inline_message_id"`
	// The query that was used to obtain the result.
	Query string `json:"query"`
}
