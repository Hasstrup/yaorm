package models 

import "github.com/hasstrup/yaorm/internal/app/errors"

type ModelInstance struct {
	Model
	Errors []errors.DatabaseOpError
}

func (m *ModelInstance) Save() bool {
   return true
}

func (m *ModelInstance) Valid() bool {
	return true
}
