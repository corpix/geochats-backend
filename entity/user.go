package entity

// User describes a user information
type User struct {
	Name string `json:"name" bson:"name"`
	Image
}
