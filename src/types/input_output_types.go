package types

// User is a representation of a chat user
type User struct {
	ConnectionID string `json:"connectionId" bson:"connectionId,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
}
