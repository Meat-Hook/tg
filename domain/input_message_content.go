package domain

var (
	_ InputMessageContent = InputTextMessageContent{}
	_ InputMessageContent = InputLocationMessageContent{}
	_ InputMessageContent = InputVenueMessageContent{}
	_ InputMessageContent = InputContactMessageContent{}
	_ InputMessageContent = InputInvoiceMessageContent{}
)

type InputMessageContent interface{ isInputMessageContent() }
