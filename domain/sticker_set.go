package domain

// StickerSet This object represents a sticker set.
type StickerSet struct {
	// Sticker set name.
	Name string `json:"name"`
	// Sticker set title.
	Title string `json:"title"`
	// True, if the sticker set contains animated stickers.
	IsAnimated bool `json:"is_animated"`
	// True, if the sticker set contains masks.
	ContainsMasks bool `json:"contains_masks"`
	// List of all set stickers.
	Stickers []Sticker `json:"stickers"`
	// Sticker set thumbnail in the WEBP or. TGS format.
	//
	// Optional.
	Thumb *PhotoSize `json:"thumb"`
}
