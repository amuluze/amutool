package query

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type singleUInt32Query struct {
	Operator string `json:"operator" dc:"操作" eg:"="`
	Value    uint32 `json:"value" dc:"字段名" eg:"100"`
}

type UInt32Query []singleUInt32Query

func (s *UInt32Query) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}

	var subSql []string
	for _, singleQuery := range *s {
		switch singleQuery.Operator {
		case "=", "!=", ">", "<":
			subSql = append(subSql, fmt.Sprintf(`%s %s %d`, name, singleQuery.Operator, singleQuery.Value))
		default:
			return "", errors.New("invalid filter operation")
		}
	}
	condition := "(" + strings.Join(subSql, " OR ") + ")"
	if sql == "" {
		return condition, nil
	} else {
		return sql + " AND " + condition, nil
	}
}

func (s *UInt32Query) Marshal() (datatypes.JSONSlice[string], error) {
	result := datatypes.JSONSlice[string]{}
	if s == nil {
		return nil, nil
	}
	for _, singleQuery := range *s {
		b, err := json.Marshal(singleQuery)
		if err != nil {
			return result, err
		}
		result = append(result, string(b))
	}
	return result, nil
}

func (s *UInt32Query) Unmarshal(values datatypes.JSONSlice[string]) error {
	if s == nil || len(values) == 0 {
		return nil
	}
	for _, value := range values {
		var saq singleUInt32Query
		err := json.Unmarshal([]byte(value), &saq)
		if err != nil {
			return err
		}
		*s = append(*s, saq)
	}
	return nil
}

type uInt32ArrayQuery struct {
	Operator string   `json:"operator" dc:"操作" eg:"="`
	Value    []uint32 `json:"value" dc:"字段名" eg:"[100, 200]"`
}

type UInt32InArrayQuery struct {
	Operator string   `json:"operator" dc:"操作" eg:"="`
	Value    []uint32 `json:"value" dc:"字段名" eg:"[100, 200]"`
}

func (s *UInt32InArrayQuery) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}

	var subSql []string
	switch s.Operator {
	case "=":
		if len(s.Value) == 1 {
			subSql = append(subSql, fmt.Sprintf("%s = %d", name, s.Value[0]))
		} else {
			condition := ""
			for _, target := range s.Value {
				condition += fmt.Sprintf("%d,", target)
			}
			subSql = append(subSql, fmt.Sprintf("%s IN (%s)", name, strings.Trim(condition, ",")))
		}
	case "!=":
		if len(s.Value) == 1 {
			subSql = append(subSql, fmt.Sprintf("%s != %d", name, s.Value[0]))
		} else {
			condition := ""
			for _, target := range s.Value {
				condition += fmt.Sprintf("%d,", target)
			}
			subSql = append(subSql, fmt.Sprintf("%s NOT IN (%s)", name, strings.Trim(condition, ",")))
		}
	case ">", "<":
		subSql = append(subSql, fmt.Sprintf("%s %s %d", name, s.Operator, s.Value[0]))
	default:
		return "", errors.New("invalid filter operation")
	}
	condition := "(" + strings.Join(subSql, " OR ") + ")"
	if sql == "" {
		return condition, nil
	} else {
		return sql + " AND " + condition, nil
	}
}

func (s *UInt32InArrayQuery) Marshal() (datatypes.JSON, error) {
	if s == nil {
		return nil, nil
	}
	b, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *UInt32InArrayQuery) Unmarshal(value datatypes.JSON) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value, s)
	if err != nil {
		return err
	}
	return nil
}

// GormQuerySQL 使用 gormDB 解析 UInt32InArrayQuery 并生成查询
func (s *UInt32InArrayQuery) GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error) {
	if s == nil || len(s.Value) == 0 {
		return db, nil
	}

	switch s.Operator {
	case "=":
		if len(s.Value) == 1 {
			return db.Where(fmt.Sprintf("%s = ?", name), s.Value[0]), nil
		} else {
			return db.Where(fmt.Sprintf("%s IN ?", name), s.Value), nil
		}
	case "!=":
		if len(s.Value) == 1 {
			return db.Where(fmt.Sprintf("%s != ?", name), s.Value[0]), nil
		} else {
			return db.Where(fmt.Sprintf("%s NOT IN ?", name), s.Value), nil
		}
	case ">":
		return db.Where(fmt.Sprintf("%s > ?", name), s.Value[0]), nil
	case "<":
		return db.Where(fmt.Sprintf("%s < ?", name), s.Value[0]), nil
	default:
		return nil, errors.New("invalid filter operation")
	}
}

type singleUInt64Query struct {
	Operator string `json:"operator" dc:"操作" eg:"="`
	Value    uint64 `json:"value" dc:"字段名" eg:"100"`
}

type UInt64Query []singleUInt64Query

func (s *UInt64Query) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}

	var subSql []string
	for _, singleQuery := range *s {
		switch singleQuery.Operator {
		case "=", "!=", ">", "<":
			subSql = append(subSql, fmt.Sprintf(`%s %s %d`, name, singleQuery.Operator, singleQuery.Value))
		default:
			return "", errors.New("invalid filter operation")
		}
	}
	condition := "(" + strings.Join(subSql, " OR ") + ")"
	if sql == "" {
		return condition, nil
	} else {
		return sql + " AND " + condition, nil
	}
}

func (s *UInt64Query) Marshal() (datatypes.JSONSlice[string], error) {
	result := datatypes.JSONSlice[string]{}
	if s == nil {
		return nil, nil
	}
	for _, singleQuery := range *s {
		b, err := json.Marshal(singleQuery)
		if err != nil {
			return result, err
		}
		result = append(result, string(b))
	}
	return result, nil
}

func (s *UInt64Query) Unmarshal(values datatypes.JSONSlice[string]) error {
	if s == nil || len(values) == 0 {
		return nil
	}
	for _, value := range values {
		var saq singleUInt64Query
		err := json.Unmarshal([]byte(value), &saq)
		if err != nil {
			return err
		}
		*s = append(*s, saq)
	}
	return nil
}
