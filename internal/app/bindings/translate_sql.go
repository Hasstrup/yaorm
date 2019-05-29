package bindings

import (
	"github.com/hasstrup/yaorm/internal/app/errors"
	"github.com/hasstrup/yaorm/internal/app/models"
)


const sqlOperationMap = map[string]func(target interface{})((result models.DatabasePayload, errors.DatabaseOpError)){
   "createRow": CreateRowBinding
}

func TranslateSQL(target interface{}, operation string) {
  
}