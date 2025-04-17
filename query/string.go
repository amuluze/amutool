package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type singleStringQuery struct {
	Operator string `json:"operator" dc:"操作" eg:"="`
	Value    string `json:"value" dc:"字段值" eg:"value"`
}

type StringQuery []singleStringQuery

func (s *StringQuery) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}

	var subSql []string
	for _, singleQuery := range *s {
		switch singleQuery.Operator {
		case "=", "!=":
			subSql = append(subSql, fmt.Sprintf(`%s %s '%s'`, name, singleQuery.Operator, EscapeStringValue([]byte(singleQuery.Value))))
		case "<>": // 包含 position 用户查找子串在字符串中的位置
			subSql = append(subSql, fmt.Sprintf(`positionCaseInsensitive(%s, '%s') > 0`, name, EscapeStringValue([]byte(singleQuery.Value))))
		case "><": // 不包含
			subSql = append(subSql, fmt.Sprintf(`positionCaseInsensitive(%s, '%s') = 0`, name, EscapeStringValue([]byte(singleQuery.Value))))
		case "regex": // 正则匹配
			subSql = append(subSql, fmt.Sprintf(`match(%s, '%s')`, name, EscapeStringValue([]byte(singleQuery.Value))))
		case "not regex": // 正则不匹配
			subSql = append(subSql, fmt.Sprintf(`NOT match(%s, '%s')`, name, EscapeStringValue([]byte(singleQuery.Value))))
		case "net contained or equal": // 属于网段
			if _, ok := NetContainedFields[name]; !ok {
				return "", errors.New(fmt.Sprintf("%s cannot use operation '%s'", name, singleQuery.Operator))
			}
			ipSegment, err := NewIPSegment(singleQuery.Value)
			if err != nil {
				return "", err
			}
			if err := ipSegment.Build(); err != nil {
				return "", err
			}
			subSql = append(subSql, fmt.Sprintf("IPv4StringToNum('%s') <= IPv4StringToNum(%s) <= IPv4StringToNum('%s')", ipSegment.Begin, name, ipSegment.End))
		case "not net contained or equal": // 不属于网段
			if _, ok := NetContainedFields[name]; !ok {
				return "", errors.New(fmt.Sprintf("%s cannot use operation '%s'", name, singleQuery.Operator))
			}
			ipSegment, err := NewIPSegment(singleQuery.Value)
			if err != nil {
				return "", err
			}
			if err := ipSegment.Build(); err != nil {
				return "", err
			}
			subSql = append(subSql, fmt.Sprintf("IPv4StringToNum(%s) <= IPv4StringToNum('%s') || IPv4StringToNum(%s) >= IPv4StringToNum('%s')", name, ipSegment.Begin, name, ipSegment.End))
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

func (s *StringQuery) Marshal() (datatypes.JSONSlice[string], error) {
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

func (s *StringQuery) Unmarshal(values datatypes.JSONSlice[string]) error {
	if s == nil || len(values) == 0 {
		return nil
	}
	for _, value := range values {
		var sq singleStringQuery
		err := json.Unmarshal([]byte(value), &sq)
		if err != nil {
			return err
		}
		*s = append(*s, sq)
	}
	return nil
}

type stringArrayQuery struct {
	Operator string   `json:"operator" dc:"操作" eg:"="`
	Value    []string `json:"value" dc:"字段值" eg:"value"`
}

type StringInArrayQuery struct {
	Operator string   `json:"operator" dc:"操作" eg:"="`
	Value    []string `json:"value" dc:"字段值" eg:"value"`
}

func (s *StringInArrayQuery) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}

	var subSql []string
	switch s.Operator {
	case "=":
		if len(s.Value) == 1 {
			subSql = append(subSql, fmt.Sprintf(`%s = '%s'`, name, s.Value[0]))
		} else {
			cond := ""
			for _, target := range s.Value {
				cond += fmt.Sprintf("'%s',", target)
			}
			subSql = append(subSql, fmt.Sprintf(`%s IN (%s)`, name, strings.Trim(cond, ",")))
		}
	case "!=":
		if len(s.Value) == 1 {
			subSql = append(subSql, fmt.Sprintf(`%s != '%s'`, name, s.Value[0]))
		} else {
			cond := ""
			for _, target := range s.Value {
				cond += fmt.Sprintf("'%s',", target)
			}
			subSql = append(subSql, fmt.Sprintf(`%s NOT IN (%s)`, name, strings.Trim(cond, ",")))
		}
	case "<>":
		for _, target := range s.Value {
			subSql = append(subSql, fmt.Sprintf(`position(%s, '%s') > 0`, name, EscapeStringValue([]byte(target))))
		}
	case "><":
		for _, target := range s.Value {
			subSql = append(subSql, fmt.Sprintf(`not position(%s, '%s') > 0`, name, EscapeStringValue([]byte(target))))
		}
	case "regex":
		for _, target := range s.Value {
			subSql = append(subSql, fmt.Sprintf(`match(%s, '%s')`, name, EscapeStringValue([]byte(target))))
		}
	case "not regex":
		for _, target := range s.Value {
			subSql = append(subSql, fmt.Sprintf(`not match(%s, '%s')`, name, EscapeStringValue([]byte(target))))
		}
	case "net contained or equal":
		for _, target := range s.Value {
			if _, ok := NetContainedFields[name]; !ok {
				return "", errors.New(fmt.Sprintf("%s cannot use operation '%s'", name, s.Operator))
			}

			ipSegment, err := NewIPSegment(target)
			if err != nil {
				return "", err
			}
			if err := ipSegment.Build(); err != nil {
				return "", err
			}
			if strings.Contains(target, ":") {
				subSql = append(subSql, fmt.Sprintf("%s AND %s AND lessOrEquals(IPv6StringToNum('%s'), IPv6StringToNum(trim(%s))) AND greaterOrEquals(IPv6StringToNum('%s'), IPv6StringToNum(trim(%s)))", fmt.Sprintf("%s IS NOT NULL", name), fmt.Sprintf("isIPv6String(%s)", name), ipSegment.Begin, name, ipSegment.End, name))
			} else {
				subSql = append(subSql, fmt.Sprintf("%s AND %s AND lessOrEquals(IPv4StringToNum('%s'), IPv4StringToNum(trim(%s))) AND greaterOrEquals(IPv4StringToNum('%s'), IPv4StringToNum(trim(%s)))", fmt.Sprintf("%s IS NOT NULL", name), fmt.Sprintf("isIPv4String(%s)", name), ipSegment.Begin, name, ipSegment.End, name))
			}
		}
	case "not net contained or equal":
		for _, target := range s.Value {
			if _, ok := NetContainedFields[name]; !ok {
				return "", errors.New(fmt.Sprintf("%s cannot use operation '%s'", name, s.Operator))
			}

			ipSegment, err := NewIPSegment(target)
			if err != nil {
				return "", err
			}
			if err := ipSegment.Build(); err != nil {
				return "", err
			}
			if strings.Contains(target, ":") {
				subSql = append(subSql, fmt.Sprintf("%s AND %s AND (lessOrEquals(IPv6StringToNum(trim(%s)), IPv6StringToNum('%s')) OR greaterOrEquals(IPv6StringToNum(trim(%s)), IPv6StringToNum('%s')))", fmt.Sprintf("%s IS NOT NULL", name), fmt.Sprintf("isIPv6String(%s)", name), name, ipSegment.Begin, name, ipSegment.End))
			} else {
				subSql = append(subSql, fmt.Sprintf("%s AND %s AND (lessOrEquals(IPv4StringToNum(trim(%s)), IPv4StringToNum('%s')) OR greaterOrEquals(IPv4StringToNum(trim(%s)), IPv4StringToNum('%s')))", fmt.Sprintf("%s IS NOT NULL", name), fmt.Sprintf("isIPv4String(%s)", name), name, ipSegment.Begin, name, ipSegment.End))
			}
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

func (s *StringInArrayQuery) Marshal() (datatypes.JSON, error) {
	if s == nil {
		return nil, nil
	}
	b, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *StringInArrayQuery) Unmarshal(value datatypes.JSON) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value, s)
	if err != nil {
		return err
	}
	return nil
}

// GormQuerySQL 使用 gormDB 解析 StringInArrayQuery 并生成查询
func (s *StringInArrayQuery) GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error) {
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
	case "<>": // 包含
		query := db
		for _, target := range s.Value {
			if query.Dialector.Name() == "postgres" {
				escapedTarget := "%" + strings.Trim(target, " ") + "%"
				query = query.Where(fmt.Sprintf("%s LIKE ?", name), escapedTarget)
			}
			if query.Dialector.Name() == "clickhouse" {
				escapedTarget := EscapeStringValue([]byte(target))
				query = query.Where(fmt.Sprintf("position(%s, ?) > 0", name), escapedTarget)
			}
		}
		return query, nil
	case "><": // 不包含
		query := db
		for _, target := range s.Value {
			if query.Dialector.Name() == "postgres" {
				escapedTarget := "%" + strings.Trim(target, " ") + "%"
				query = query.Where(fmt.Sprintf("%s NOT LIKE ?", name), escapedTarget)
			}
			if query.Dialector.Name() == "clickhouse" {
				escapedTarget := EscapeStringValue([]byte(target))
				query = query.Where(fmt.Sprintf("not position(%s, ?) > 0", name), escapedTarget)
			}
		}
		return query, nil
	case "regex": // 正则匹配
		query := db
		for _, target := range s.Value {
			if query.Dialector.Name() == "postgres" {
				escapedTarget := target
				query = query.Where(fmt.Sprintf("%s ~ ?", name), escapedTarget)
			}
			if query.Dialector.Name() == "clickhouse" {
				escapedTarget := EscapeStringValue([]byte(target))
				query = query.Where(fmt.Sprintf("match(%s, ?)", name), escapedTarget)
			}
		}
		return query, nil
	case "not regex": // 正则不匹配
		query := db
		for _, target := range s.Value {
			if query.Dialector.Name() == "postgres" {
				escapedTarget := target
				query = query.Where(fmt.Sprintf("%s !~ ?", name), escapedTarget)
			}
			if query.Dialector.Name() == "clickhouse" {
				escapedTarget := EscapeStringValue([]byte(target))
				query = query.Where(fmt.Sprintf("not match(%s, ?)", name), escapedTarget)
			}
		}
		return query, nil
	case "net contained or equal": // 属于网段
		if _, ok := NetContainedFields[name]; !ok {
			return nil, fmt.Errorf("%s cannot use operation '%s'", name, s.Operator)
		}

		query := db
		for _, target := range s.Value {
			if query.Dialector.Name() == "postgres" {
				ipSegment, err := NewIPSegment(target)
				if err != nil {
					return nil, err
				}
				if err := ipSegment.Build(); err != nil {
					return nil, err
				}

				if strings.Contains(target, ":") {
					// IPv6
					condition := fmt.Sprintf("%s IS NOT NULL AND %s::inet >= ?::inet AND %s::inet <= ?::inet",
						name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				} else {
					// IPv4
					condition := fmt.Sprintf("%s IS NOT NULL AND %s::inet >= ?::inet AND %s::inet <= ?::inet",
						name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				}
			}
			if query.Dialector.Name() == "clickhouse" {
				ipSegment, err := NewIPSegment(target)
				if err != nil {
					return nil, err
				}
				if err := ipSegment.Build(); err != nil {
					return nil, err
				}

				if strings.Contains(target, ":") {
					// IPv6
					condition := fmt.Sprintf("%s IS NOT NULL AND isIPv6String(%s) AND lessOrEquals(IPv6StringToNum(?), IPv6StringToNum(trim(%s))) AND greaterOrEquals(IPv6StringToNum(?), IPv6StringToNum(trim(%s)))",
						name, name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				} else {
					// IPv4
					condition := fmt.Sprintf("%s IS NOT NULL AND isIPv4String(%s) AND lessOrEquals(IPv4StringToNum(?), IPv4StringToNum(trim(%s))) AND greaterOrEquals(IPv4StringToNum(?), IPv4StringToNum(trim(%s)))",
						name, name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				}
			}
		}
		return query, nil
	case "not net contained or equal": // 不属于网段
		if _, ok := NetContainedFields[name]; !ok {
			return nil, errors.New(fmt.Sprintf("%s cannot use operation '%s'", name, s.Operator))
		}

		query := db
		for _, target := range s.Value {
			if query.Dialector.Name() == "postgres" {
				ipSegment, err := NewIPSegment(target)
				if err != nil {
					return nil, err
				}
				if err := ipSegment.Build(); err != nil {
					return nil, err
				}

				if strings.Contains(target, ":") {
					// IPv6
					condition := fmt.Sprintf("%s IS NOT NULL AND (%s::inet < ?::inet OR %s::inet > ?::inet)",
						name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				} else {
					// IPv4
					condition := fmt.Sprintf("%s IS NOT NULL AND (%s::inet < ?::inet OR %s::inet > ?::inet)",
						name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				}
			}
			if query.Dialector.Name() == "clickhouse" {
				ipSegment, err := NewIPSegment(target)
				if err != nil {
					return nil, err
				}
				if err := ipSegment.Build(); err != nil {
					return nil, err
				}

				if strings.Contains(target, ":") {
					// IPv6
					condition := fmt.Sprintf("%s IS NOT NULL AND isIPv6String(%s) AND (lessOrEquals(IPv6StringToNum(trim(%s)), IPv6StringToNum(?)) OR greaterOrEquals(IPv6StringToNum(trim(%s)), IPv6StringToNum(?)))",
						name, name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				} else {
					// IPv4
					condition := fmt.Sprintf("%s IS NOT NULL AND isIPv4String(%s) AND (lessOrEquals(IPv4StringToNum(trim(%s)), IPv4StringToNum(?)) OR greaterOrEquals(IPv4StringToNum(trim(%s)), IPv4StringToNum(?)))",
						name, name, name, name)
					query = query.Where(condition, ipSegment.Begin, ipSegment.End)
				}
			}
		}
		return query, nil
	default:
		return nil, errors.New("invalid filter operation")
	}
}
