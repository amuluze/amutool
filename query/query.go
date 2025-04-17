package query

import (
	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type SQLField interface {
	GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error)
	QuerySQL(name string, sql string) (string, error)
	Marshal() (datatypes.JSON, error)
	Unmarshal(datatypes.JSON) error
}
