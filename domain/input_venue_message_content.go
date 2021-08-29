package domain

// InputVenueMessageContent Represents the content of a venue message to be sent as the result of an inline query.
type InputVenueMessageContent struct {
	// Google Places identifier of the venue.
	//
	// Optional.
	GooglePlaceId string `json:"google_place_id"`
	// Google Places type of the venue.
	// (See supported types).
	//
	// Optional.
	GooglePlaceType string `json:"google_place_type"`
	// Latitude of the venue in degrees.
	Latitude float64 `json:"latitude"`
	// Longitude of the venue in degrees.
	Longitude float64 `json:"longitude"`
	// Name of the venue.
	Title string `json:"title"`
	// Address of the venue.
	Address string `json:"address"`
	// Foursquare identifier of the venue, if known.
	//
	// Optional.
	FoursquareId string `json:"foursquare_id"`
	// Foursquare type of the venue, if known.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”).
	//
	// Optional.
	FoursquareType string `json:"foursquare_type"`
}

func (i InputVenueMessageContent) isInputMessageContent() {}
