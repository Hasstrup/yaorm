package models

/*
	this model should implement a fields method, that should out of the box yield all the types
	that are available
*/
type Model struct {
	Fields []Field
}

type DatabasePayload struct {
}

// generates a new model
func NewModel(fields map[string]string) *Model {
	model := &Model{}
	return model
}

// returns the database payload
func NewDatabasePayload() DatabasePayload {
	return DatabasePayload{}
}

// creates a new model and persists to db
func (m Model) Create() DatabasePayload {
	return DatabasePayload{}
}

func (m Model) BulkCreate() []DatabasePayload {
	return []DatabasePayload{}
}

func (m Model) New() *ModelInstance {
	instance := &ModelInstance{}
	instance.Model = m
	return instance
}
