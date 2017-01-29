package trackPoint

import (
	"math"
	"sort"
	"time"
)

// haversin(θ) function
func hsin(theta float64) float64 {
	return math.Pow(math.Sin(theta/2), 2)
}

// Distance function returns the distance (in meters) between two points of
//     a given longitude and latitude relatively accurately (using a spherical
//     approximation of the Earth) through the Haversin Distance Formula for
//     great arc distance on a sphere with accuracy for small distances
//
// point coordinates are supplied in degrees and converted into rad. in the func
//
// distance returned is METERS!!!!!!
// http://en.wikipedia.org/wiki/Haversine_formula
func Distance(lat1, lon1, lat2, lon2 float64) float64 {
	// convert to radians
	// must cast radius as float to multiply later
	var la1, lo1, la2, lo2, r float64
	la1 = lat1 * math.Pi / 180
	lo1 = lon1 * math.Pi / 180
	la2 = lat2 * math.Pi / 180
	lo2 = lon2 * math.Pi / 180

	r = 6378100 // Earth radius in METERS

	// calculate
	h := hsin(la2-la1) + math.Cos(la1)*math.Cos(la2)*hsin(lo2-lo1)

	return 2 * r * math.Asin(math.Sqrt(h))
}

//CatStats
type CatStats struct {
	Speed     DotStats `json:"speed_stats"`
	Elevation DotStats `json:"elevation_stats"`
	Distance  float64  `json:"distance"`
	Count     int      `json:"count"`
}

//DotStats
type DotStats struct {
	Max  float64 `json:"max"`
	Min  float64 `json:"min"`
	Mean float64 `json:"mean"`
}

//ForName returns all points in a given set belonging to a given name.
func (points TrackPoints) ForName(name string) (catpoints TrackPoints) {
	for _, point := range points {
		if point.Name == name {
			catpoints = append(catpoints, point)
		}
	}
	return catpoints
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

//GetUniqueNames yep
func (points TrackPoints) UniqueNames() (names []string) {
	for _, point := range points {
		if !stringInSlice(point.Name, names) {
			names = append(names, point.Name)
		}
	}
	return names
}

//Since return all points from a group after a given time.
func (points TrackPoints) Since(when time.Time) (catpoints TrackPoints) {
	for _, point := range points {
		if point.Time.After(when) {
			catpoints = append(catpoints, point)
		}
	}
	return catpoints
}

//Statistics pull stats for a given set of TrackPoints
func (points TrackPoints) Statistics() (stats CatStats) {

	sort.Sort(points) //set em up (by time)
	stats.Count = len(points)

	stats.Elevation.Mean = 0.0
	stats.Speed.Mean = 0.0

	for i, point := range points {
		if i != 0 {
			stats.Distance = stats.Distance + Distance(point.Lat, point.Lng, points[i-1].Lat, points[i-1].Lng)
		}
		if i == 0 {
			stats.Speed.Min = point.Speed
			stats.Speed.Min = point.Elevation
		}
		if point.Speed > stats.Speed.Max {
			stats.Speed.Max = point.Speed
		}
		if point.Speed < stats.Speed.Min {
			stats.Speed.Min = point.Speed
		}
		if point.Elevation > stats.Elevation.Max {
			stats.Elevation.Max = point.Elevation
		}
		if point.Elevation < stats.Elevation.Min {
			stats.Elevation.Min = point.Elevation
		}

		stats.Elevation.Mean = stats.Elevation.Mean + point.Elevation
		stats.Speed.Mean = stats.Elevation.Mean + point.Speed
	}

	stats.Elevation.Mean = stats.Elevation.Mean / float64(stats.Count)
	stats.Speed.Mean = stats.Speed.Mean / float64(stats.Count)

	return stats

}
