package domain

// ReplyMarkup is InlineKeyboardMarkup or ReplyKeyboardMarkup or ReplyKeyboardRemove or ForceReply.
type ReplyMarkup interface{ isChatMember() }

var (
	_ ReplyMarkup = InlineKeyboardMarkup{}
	_ ReplyMarkup = ReplyKeyboardMarkup{}
	_ ReplyMarkup = ReplyKeyboardRemove{}
	_ ReplyMarkup = ForceReply{}
)
