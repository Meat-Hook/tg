package domain

// InputLocationMessageContent Represents the content of a location message to be sent as the result of an inline query.
type InputLocationMessageContent struct {
	// The radius of uncertainty for the location, measured in meters; 0-1500.
	//
	// Optional.
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	// Period in seconds for which the location can be updated, should be between 60 and 86400.
	//
	// Optional.
	LivePeriod int `json:"live_period"`
	// For live locations, a direction in which the user is moving, in degrees.
	// Must be between 1 and 360 if specified.
	//
	// Optional.
	Heading int `json:"heading"`
	// For live locations, a maximum distance for proximity alerts about approaching another chat member, in meters.
	// Must be between 1 and 100000 if specified.
	//
	// Optional.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
	// Latitude of the location in degrees.
	Latitude float64 `json:"latitude"`
	// Longitude of the location in degrees.
	Longitude float64 `json:"longitude"`
}

func (i InputLocationMessageContent) isInputMessageContent() {}
