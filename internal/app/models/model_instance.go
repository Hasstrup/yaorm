package models 

import (
	"reflect"
	"github.com/hasstrup/yaorm/internal/app/errors"
)

type ModelInstance struct {
	Model
	Mappings []Mapping
	Columns []string
	RowNumber int
	Errors []errors.DatabaseOpError
}

type Mapping struct {
	Field
	Value interface{}
}

func (m *ModelInstance) Save() bool {
   return true
}

func (m *ModelInstance) Valid() bool {
	if (len(m.Mappings) == 0) {
		return false
	}
	for _, mapping := range m.Mappings {
		if typeOf(mapping.Value) != mapping.Field.DataType {
			return false
		}
	}
	return true
}

func (m *ModelInstance) ToJson() {

}

func (m *ModelInstance) BuildAttributes(attrs map[string]interface{}) {
	mappings := []Mapping{}
	for _, field := range m.Fields {
	  mapping := Mapping{
		  Field: field,
		  Value: attrs[field.ColumnName],
	  }
	  mappings = append(mappings, mapping)
  }
  m.Mappings = mappings
}

func typeOf(val interface{}) string {
	return reflect.TypeOf(val).Kind().String()
}
