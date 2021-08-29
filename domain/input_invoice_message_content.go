package domain

// InputInvoiceMessageContent Represents the content of an invoice message to be sent as the result of an inline query.
type InputInvoiceMessageContent struct {
	// Pass True, if you require the user's full name to complete the order.
	//
	// Optional.
	NeedName bool `json:"need_name"`
	// Pass True, if you require the user's phone number to complete the order.
	//
	// Optional.
	NeedPhoneNumber bool `json:"need_phone_number"`
	// A JSON-serialized object for data about the invoice, which will be shared with the payment provider.
	// A detailed description of the required fields should be provided by the payment provider.
	//
	// Optional.
	ProviderData string `json:"provider_data"`
	// Photo height.
	//
	// Optional.
	PhotoHeight int `json:"photo_height"`
	// Three-letter ISO 4217 currency code, see more on currencies.
	Currency string `json:"currency"`
	// Price breakdown, a JSON-serialized list of components (e.g. product price, tax, discount, delivery cost, delivery tax, bonus, etc).
	Prices []LabeledPrice `json:"prices"`
	// The maximum accepted amount for tips in the smallest units of the currency (integer, not float/double).
	// For example, for a maximum tip of US$ 1.
	// 45 pass max_tip_amount = 145.
	// See the exp parameter in currencies.
	// json, it shows the number of digits past the decimal point for each currency (2 for the majority of currencies).
	// Defaults to 0.
	//
	// Optional.
	MaxTipAmount int `json:"max_tip_amount"`
	// Photo width.
	//
	// Optional.
	PhotoWidth int `json:"photo_width"`
	// Pass True, if user's phone number should be sent to provider.
	//
	// Optional.
	SendPhoneNumberToProvider bool `json:"send_phone_number_to_provider"`
	// Pass True, if the final price depends on the shipping method.
	//
	// Optional.
	IsFlexible bool `json:"is_flexible"`
	// Product name, 1-32 characters.
	Title string `json:"title"`
	// Payment provider token, obtained via Botfather.
	ProviderToken string `json:"provider_token"`
	// Photo size.
	//
	// Optional.
	PhotoSize int `json:"photo_size"`
	// Pass True, if you require the user's email address to complete the order.
	//
	// Optional.
	NeedEmail bool `json:"need_email"`
	// A JSON-serialized array of suggested amounts of tip in the smallest units of the currency (integer, not float/double).
	// At most 4 suggested tip amounts can be specified.
	// The suggested tip amounts must be positive, passed in a strictly increased order and must not exceed max_tip_amount.
	//
	// Optional.
	SuggestedTipAmounts []int `json:"suggested_tip_amounts"`
	// URL of the product photo for the invoice.
	// Can be a photo of the goods or a marketing image for a service.
	// People like it better when they see what they are paying for.
	//
	// Optional.
	PhotoUrl string `json:"photo_url"`
	// Pass True, if you require the user's shipping address to complete the order.
	//
	// Optional.
	NeedShippingAddress bool `json:"need_shipping_address"`
	// Pass True, if user's email address should be sent to provider.
	//
	// Optional.
	SendEmailToProvider bool `json:"send_email_to_provider"`
	// Product description, 1-255 characters.
	Description string `json:"description"`
	// Bot-defined invoice payload, 1-128 bytes.
	// This will not be displayed to the user, use for your internal processes.
	Payload string `json:"payload"`
}

func (i InputInvoiceMessageContent) isInputMessageContent() {}
