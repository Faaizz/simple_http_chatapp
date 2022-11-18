package types

type Connection struct {
	ConnectionID string `json:"connectionId"`
}

// User is a representation of a chat user
type User struct {
	ConnectionID string `json:"connectionId" bson:"connectionId,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
}

// Message is a representation of a chat message and destination
type Message struct {
	ConnectionID string `json:"connectionId" bson:"connectionId"`
	Message      string `json:"message" bson:"message"`
	Username     string `json:"username" bson:"username,omitempty"`
	FromUsername string `json:"from_username" bson:"from_username"`
	URL          string `json:"url" bson:"url"` // WebSocket connection callback base URL of the form: https://{api-id}.execute-api.us-east-1.amazonaws.com/{stage}
}
