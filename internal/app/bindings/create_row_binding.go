package bindings

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/hasstrup/yaorm/internal/app/errors"
	"github.com/hasstrup/yaorm/internal/app/models"
)

func createRowBinding(target interface{}) (models.DatabasePayload, errors.DatabaseOpError) {
	instance = target.(models.ModelInstance)
	values = valuesFromSource(instance.Values)
	columns = columnsFromSource(instance.Columns)
	tableName = strings.ToLower(instance.Table)
	sql = fmt.Sprintf("INSERT INTO %v (%s) VALUES(%v)", tableName, columns, values)
}

func columnsFromSource(columns []string) string {
	strings.Join(columns, ",")
}

func valuesFromSource(values []models.Mapping) string {
	keys := []string{}
	for _, val := range values {
		switch t := val.Value.(type) {
		case int:
			append(keys, strconv.Itoa(val.Value.(int)))
		case string:
			append(keys, val.Value.(string))
			// TODO: we need to add more types as we come along
		}
	}
	return strings.Join(keys, ",")
}
