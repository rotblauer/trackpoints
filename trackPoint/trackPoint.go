package trackPoint

import (
	"strconv"
	"time"
)

// TrackPoint Stores a snippet of life, love, and location
type TrackPoint struct {
	ID        string    `json:"id"` //either bolt auto id or unixnano
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
	ii, _ := strconv.Atoi(slice[i].ID) //assusmes ID is unixnano which puts them in chrono order
	jj, _ := strconv.Atoi(slice[j].ID)
	return ii < jj
}
func (slice TrackPoints) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
