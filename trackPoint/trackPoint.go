package trackPoint

import (
	"github.com/deet/simpleline"
	"time"
)

// TrackPoint Stores a snippet of life, love, and location
type TrackPoint struct {
	Uuid      string    `json:"uuid"`
	PushToken string    `json:"pushToken"`
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

// Vector deepens
func (p *TrackPoint) Vector() []float64 {
	return []float64{p.Lat, p.Lng}
}

// Scale does deepry
func (p *TrackPoint) Scale(val float64) simpleline.Point {
	np := TrackPoint{}
	np.Lat = p.Lat * val
	np.Lng = p.Lng * val
	return &np
}

//Subtract add depee
func (p *TrackPoint) Subtract(p2 simpleline.Point) simpleline.Point {
	p2Cast := p2.(*TrackPoint)
	np := TrackPoint{}
	np.Lat = p.Lng - p2Cast.Lat
	return &np
}

// Zero is more deep
func (p *TrackPoint) Zero() simpleline.Point {
	np := TrackPoint{}
	return &np
}

//TrackPoints is plural. might implement Len method for Sortabliilty
type TrackPoints []TrackPoint

//TPs has comm
type TPs []*TrackPoint

// TrackPoints will implement all the methods required to satisfy
// the sort.Interface interface
func (slice TPs) Len() int {
	return len(slice)
}
func (slice TPs) Less(i, j int) bool {
	return slice[i].Time.UnixNano() > slice[j].Time.UnixNano() //ids autoincrement //but time is better because it doesn't depend on the order of the array received from the request to bolty
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
	return slice[i].Time.UnixNano() > slice[j].Time.UnixNano() //ids autoincrement //but time is better because it doesn't depend on the order of the array received from the request to bolty
}
func (slice TrackPoints) Swap(i, j int) {
	slice[i], slice[j] = slice[j], slice[i]
}
