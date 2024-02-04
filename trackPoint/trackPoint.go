package trackPoint

import (
	"net/http"
	"time"
)

// TrackPoint Stores a snippet of life, love, and location
type TrackPoint struct {
	Uuid            string    `json:"uuid"`
	PushToken       string    `json:"pushToken"`
	Version         string    `json:"version"`
	ID              int64     `json:"id"` // either bolt auto id or unixnano //think nano is better cuz can check for dupery
	Name            string    `json:"name"`
	Lat             float64   `json:"lat"`
	Lng             float64   `json:"long"`
	Accuracy        float64   `json:"accuracy"`       // horizontal, in meters
	VAccuracy       float64   `json:"vAccuracy"`      // vertical, in meteres
	Elevation       float64   `json:"elevation"`      // in meters
	Speed           float64   `json:"speed"`          // in kilometers per hour
	SpeedAccuracy   float64   `json:"speed_accuracy"` // in meters per second
	Tilt            float64   `json:"tilt"`           // degrees?
	Heading         float64   `json:"heading"`        // in degrees
	HeadingAccuracy float64   `json:"heading_accuracy"`
	HeartRate       float64   `json:"heartrate"` // bpm
	Time            time.Time `json:"time"`
	Floor           int       `json:"floor"` // building floor if available
	Notes           string    `json:"notes"` // special events of the day
	COVerified      bool      `json:"COVerified"`
	RemoteAddr      string    `json:"remoteaddr"`
}

func (p *TrackPoint) String() string {
	return p.Name + " " + p.Time.Round(time.Second).String()
}

// TrackPoints is plural. might implement Len method for Sortabliilty
type TrackPoints []*TrackPoint

// TPs has comm
type TPs []*TrackPoint

func (tps TrackPoints) Verified() {
	for i, _ := range tps {
		tps[i].COVerified = true
	}
}

func (tps TrackPoints) Unverified(r *http.Request) {
	for i, _ := range tps {
		tps[i].COVerified = false
		tps[i].RemoteAddr = r.RemoteAddr
	}
}

// TrackPoints will implement all the methods required to satisfy
// the sort.Interface interface
func (slice TPs) Len() int {
	return len(slice)
}
func (slice TPs) Less(i, j int) bool {
	return slice[i].Time.UnixNano() > slice[j].Time.UnixNano() // ids autoincrement //but time is better because it doesn't depend on the order of the array received from the request to bolty
}
func (slice TPs) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}

// TrackPoints will implement all the methods required to satisfy
// the sort.Interface interface
func (slice TrackPoints) Len() int {
	return len(slice)
}
func (slice TrackPoints) Less(i, j int) bool {
	return slice[i].Time.UnixNano() > slice[j].Time.UnixNano() // ids autoincrement //but time is better because it doesn't depend on the order of the array received from the request to bolty
}
func (slice TrackPoints) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
