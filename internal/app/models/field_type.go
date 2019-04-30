package models

/*
	should be the definition of every field in a model
	ideally should be able to yield the
*/
type Field struct {
	DataType   string
	Index      bool
	PrimaryKey bool
	ForeignKey bool
	References Model
	Check      func(Model) bool
	Nullable   bool
	ColumnName string
	Unique     bool
}

func (f *Field) Validate() bool {
	return true
}

// returns a new field.  We should be able to take a hash of the attributes
// and convert into the field directly
func NewField() *Field {
	return &Field{}
}
