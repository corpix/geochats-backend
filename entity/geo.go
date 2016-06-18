package entity

// Geo represents a geo data
type Geo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

// Point represents a uniquely addressable point on the map
// Eg, the point that has an ID as a primary key and some data associated with it(title, ...)
type Point struct {
	ID    string `json:"id"`
	Title string `json:"title"`
	Geo
}

// Area represents an area on the map
type Area struct {
	Geo
	LatitudeDelta  float64 `json:"latitudeDelta"`
	LongitudeDelta float64 `json:"longitudeDelta"`
}
