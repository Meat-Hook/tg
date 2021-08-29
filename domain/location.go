package domain

// Location This object represents a point on the map.
type Location struct {
	// Latitude as defined by sender.
	Latitude float64 `json:"latitude"`
	// The radius of uncertainty for the location, measured in meters; 0-1500.
	//
	// Optional.
	HorizontalAccuracy float64 `json:"horizontal_accuracy"`
	// Time relative to the message sending date, during which the location can be updated, in seconds.
	// For active live locations only.
	//
	// Optional.
	LivePeriod int `json:"live_period"`
	// The direction in which user is moving, in degrees; 1-360.
	// For active live locations only.
	//
	// Optional.
	Heading int `json:"heading"`
	// Maximum distance for proximity alerts about approaching another chat member, in meters.
	// For sent live locations only.
	//
	// Optional.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
	// Longitude as defined by sender.
	Longitude float64 `json:"longitude"`
}
