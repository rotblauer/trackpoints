package trackPoint

import (
	"time"
)

// TrackPoint Stores a snippet of life, love, and location
type TrackPoint struct {
	ID        int64     `json:"id"` //either bolt auto id or unixnano //think nano is better cuz can check for dupery
	Name      string    `json:"name"`
	Lat       float64   `json:"lat"`
	Lng       float64   `json:"long"`
	Accuracy  float64   `json:"accuracy"`
	Elevation float64   `json:"elevation"` //in meters
	Speed     float64   `json:"speed"`     //in kilometers per hour
	Tilt      float64   `json:"tilt"`      //degrees?
	Heading   float64   `json:"heading"`   //in degrees
	HeartRate float64   `json:"heartrate"` // bpm
	Time      time.Time `json:"time"`
	Notes     string    `json:"notes"` //special events of the day
}

//TrackPoints is plural. might implement Len method for Sortabliilty
type TrackPoints []TrackPoint

// TrackPoints will implement all the methods required to satisfy
// the sort.Interface interface
func (slice TrackPoints) Len() int {
	return len(slice)
}
func (slice TrackPoints) Less(i, j int) bool {
	return slice[i].Time.UnixNano() > slice[j].Time.UnixNano() //ids autoincrement //but time is better because it doesn't depend on the order of the array received from the request to bolty
}
func (slice TrackPoints) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
