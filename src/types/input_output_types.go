package types

// User is a representation of a chat user
type User struct {
	ConnectionID string `json:"connectionId" bson:"connectionId,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
}

// Message is a representation of a chat message and destination
type Message struct {
	ConnectionID string `json:"connectionId" bson:"connectionId"`
	Message      string `json:"message" bson:"message"`
	FromUsername string `json:"from_username" bson:"from_username"`
}
