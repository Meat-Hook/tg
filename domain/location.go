package domain

// Location this object represents a point on the map.
type Location struct {
	// Longitude as defined by sender.
	Longitude float64 `json:"longitude"`
	// Latitude as defined by sender
	Latitude float64 `json:"latitude"`
	// HorizontalAccuracy the radius of uncertainty for the location, measured in meters; 0-1500.
	HorizontalAccuracy float32 `json:"horizontal_accuracy"`
	// LivePeriod time relative to the message sending date, during which the location can be updated, in seconds.
	// For active live locations only.
	LivePeriod int `json:"live_period"`
	// Heading the direction in which user is moving, in degrees; 1-360. For active live locations only.
	Heading int `json:"heading"`
	// ProximityAlertRadius maximum distance for proximity alerts about approaching another chat member, in meters.
	// For sent live locations only.
	ProximityAlertRadius int `json:"proximity_alert_radius"`
}
