package types

// ConnOutput is a representation of connection data obtained from DB
type ConnOutput struct {
	ConnID   string `json:"connectionId"`
	Username string `json:"username"`
}

// ConnInput is a representation of connection data required to create a connection
type ConnInput struct {
	Username     string `json:"username"`
	ConnectionID string `json:"connectionId"`
}

type PutConnInput struct {
	ConnectionID string `json:"connectionId" bson:"connectionId,omitempty"`
	Username     string `json:"username" bson:"username,omitempty"`
}
