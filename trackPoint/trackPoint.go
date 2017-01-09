package trackPoint

import (
	"google.golang.org/appengine"
	"time"
)

// Stores a snippet of life, love, and location
type TrackPoint struct {
	Name      string             `json:"name"`
	LatLong   appengine.GeoPoint `json:"latLong"`
	Elevation float64            `json:"elevation"` //in meters
	Speed     float64            `json:"speed"`//in kilometers per hour
	Tilt      float64            `json:"tilt"` //degrees?
	Heading   float64            `json:"heading"` //in degrees
	HeartRate float64            `json:"heartrate"` // bpm
	Time      time.Time          `json:"time"`
	Notes     string             `json:"notes"` //special events of the day
}
