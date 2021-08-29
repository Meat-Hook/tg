package domain

// Venue This object represents a venue.
type Venue struct {
	// Venue location.
	// Can't be a live location.
	Location Location `json:"location"`
	// Name of the venue.
	Title string `json:"title"`
	// Address of the venue.
	Address string `json:"address"`
	// Foursquare identifier of the venue.
	//
	// Optional.
	FoursquareId string `json:"foursquare_id"`
	// Foursquare type of the venue.
	// (For example, “arts_entertainment/default”, “arts_entertainment/aquarium” or “food/icecream”).
	//
	// Optional.
	FoursquareType string `json:"foursquare_type"`
	// Google Places identifier of the venue.
	//
	// Optional.
	GooglePlaceId string `json:"google_place_id"`
	// Google Places type of the venue.
	// (See supported types).
	//
	// Optional.
	GooglePlaceType string `json:"google_place_type"`
}
