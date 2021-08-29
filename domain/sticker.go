package domain

// Sticker This object represents a sticker.
type Sticker struct {
	// Unique identifier for this file, which is supposed to be the same over time and for different bots.
	// Can't be used to download or reuse the file.
	FileUniqueId string `json:"file_unique_id"`
	// Sticker height.
	Height int `json:"height"`
	// True, if the sticker is animated.
	IsAnimated bool `json:"is_animated"`
	// Sticker thumbnail in the WEBP or JPG format.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
	// Emoji associated with the sticker.
	//
	// Optional.
	Emoji string `json:"emoji"`
	// Name of the sticker set to which the sticker belongs.
	//
	// Optional.
	SetName string `json:"set_name"`
	// For mask stickers, the position where the mask should be placed.
	//
	// Optional.
	MaskPosition *MaskPosition `json:"mask_position"`
	// Identifier for this file, which can be used to download or reuse the file.
	FileId string `json:"file_id"`
	// Sticker width.
	Width int `json:"width"`
	// File size.
	//
	// Optional.
	FileSize int `json:"file_size"`
}
