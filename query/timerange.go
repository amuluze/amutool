// Package query
// Date:   2025/4/14 09:47
// Description:
package query

import (
	"encoding/json"
	"fmt"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type TimeRange struct {
	From int64 `json:"from"`
	To   int64 `json:"to"`
}

type TimeRangeQuery struct {
	Operator string    `json:"operator" dc:"操作" eg:"="`
	Value    TimeRange `json:"value" dc:"字段值" eg:"value"`
}

func (s *TimeRangeQuery) QuerySQL(name string, sql string) (string, error) {
	return sql, nil
}

func (s *TimeRangeQuery) Marshal() (datatypes.JSON, error) {
	if s == nil {
		return nil, nil
	}
	b, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *TimeRangeQuery) Unmarshal(value datatypes.JSON) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value, s)
	if err != nil {
		return err
	}
	return nil
}

func (s *TimeRangeQuery) GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error) {
	if s == nil {
		return db, nil
	}

	switch s.Operator {
	case "=":
		return db.Where(fmt.Sprintf("%s BETWEEN to_timestamp(?) AND to_timestamp(?)", name), s.Value.From, s.Value.To), nil
	case ">":
		return db.Where(fmt.Sprintf("%s > to_timestamp(?)", name), s.Value.From), nil
	case "<":
		return db.Where(fmt.Sprintf("%s < to_timestamp(?)", name), s.Value.To), nil
	default:
		return nil, errors.New("invalid filter operation")
	}
}
