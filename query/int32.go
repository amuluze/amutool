package query

import (
	"encoding/json"
	"fmt"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Int32Query struct {
	Operator string `json:"operator" dc:"操作" eg:"="`
	Value    int32  `json:"value" dc:"字段名" eg:"100"`
}

func (s *Int32Query) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}

	var subSql []string
	switch s.Operator {
	case "=":
		if name == "skynet_rule_id" {
			subSql = append(subSql, fmt.Sprintf(`has(%s, %d)`, "req_skynet_rule_id_list", s.Value))
			subSql = append(subSql, fmt.Sprintf(`has(%s, %d)`, "rsp_skynet_rule_id_list", s.Value))
		}
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

func (s *Int32Query) Marshal() (datatypes.JSON, error) {
	if s == nil {
		return nil, nil
	}
	b, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Int32Query) Unmarshal(value datatypes.JSON) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value, s)
	if err != nil {
		return err
	}
	return nil
}

// GormQuerySQL 使用 gormDB 解析 Int32Query 并生成查询
func (s *Int32Query) GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error) {
	if s == nil {
		return db, nil
	}

	switch s.Operator {
	case "=":
		if name == "skynet_rule_id" {
			return db.Where("has(req_skynet_rule_id_list, ?) OR has(rsp_skynet_rule_id_list, ?)", s.Value, s.Value), nil
		}
		return nil, errors.New("invalid field for this operation")
	default:
		return nil, errors.New("invalid filter operation")
	}
}

type int32ArrayQuery struct {
	Operator string  `json:"operator" dc:"操作" eg:"="`
	Value    []int32 `json:"value" dc:"字段名" eg:"[100, 200]"`
}
type Int32InArrayQuery struct {
	Operator string  `json:"operator" dc:"操作" eg:"="`
	Value    []int32 `json:"value" dc:"字段名" eg:"[100, 200]"`
}

func (s *Int32InArrayQuery) QuerySQL(name string, sql string) (string, error) {
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

func (s *Int32InArrayQuery) Marshal() (datatypes.JSON, error) {
	if s == nil {
		return nil, nil
	}
	b, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *Int32InArrayQuery) Unmarshal(value datatypes.JSON) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value, s)
	if err != nil {
		return err
	}
	return nil
}

// GormQuerySQL 使用 gormDB 解析 Int32InArrayQuery 并生成查询
func (s *Int32InArrayQuery) GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error) {
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

type singleInt64Query struct {
	Operator string `json:"operator" dc:"操作" eg:"="`
	Value    uint64 `json:"value" dc:"字段名" eg:"100"`
}

type Int64Query []singleInt64Query

func (s *Int64Query) QuerySQL(name string, sql string) (string, error) {
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

func (s *Int64Query) Marshal() (datatypes.JSONSlice[string], error) {
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

func (s *Int64Query) Unmarshal(values datatypes.JSONSlice[string]) error {
	if s == nil || len(values) == 0 {
		return nil
	}
	for _, value := range values {
		var saq singleInt64Query
		err := json.Unmarshal([]byte(value), &saq)
		if err != nil {
			return err
		}
		*s = append(*s, saq)
	}
	return nil
}
