package types

// A DBAdapter provides a layer of abstraction for interaction with the underlying database
type DBAdapter interface {
	SetTableName(string)
	CheckExists() error
	PutConn(PutConnInput) error
}
