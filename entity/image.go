package entity

// Image describes the basic image that could be attached to other entities
type Image struct {
	URL string `json:"url" bson:"url" valid:"url"`
}
