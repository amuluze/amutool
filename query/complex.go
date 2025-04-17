package query

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"gorm.io/gorm"
)

type ProtectedObject struct {
	Key      string            `json:"key"`
	Value    uint64            `json:"value"`
	Children []ProtectedObject `json:"children"`
}

type ComplexInArrayQuery struct {
	Operator string            `json:"operator" dc:"操作" eg:"="`
	Value    []ProtectedObject `json:"value"`
}

func (s *ComplexInArrayQuery) QuerySQL(name string, sql string) (string, error) {
	if s == nil {
		return sql, nil
	}
	var protectedObjectIDs []uint64
	applicationIDs := make(map[uint64][]uint64)
	for _, c := range s.Value {
		if len(c.Children) > 0 {
			for _, obj := range c.Children {
				protectedObjectIDs = append(protectedObjectIDs, obj.Value)
				if len(obj.Children) > 0 {
					for _, o := range obj.Children {
						applicationIDs[obj.Value] = append(applicationIDs[obj.Value], o.Value)
					}
				}
			}
		}
	}
	fmt.Printf("pids: %v, aids: %v\n", protectedObjectIDs, applicationIDs)

	var subSql []string
	switch s.Operator {
	case "=":
		if len(protectedObjectIDs) == 1 {
			pCond := fmt.Sprintf("%s = %d", "protect_object_id", protectedObjectIDs[0])
			if len(applicationIDs[protectedObjectIDs[0]]) == 1 {
				pCond += " AND " + fmt.Sprintf("%s = %d", "app_id", applicationIDs[protectedObjectIDs[0]][0])
			} else if len(applicationIDs[protectedObjectIDs[0]]) > 1 {
				aCond := ""
				for _, id := range applicationIDs[protectedObjectIDs[0]] {
					aCond += fmt.Sprintf("'%d',", id)
				}
				pCond += " AND " + fmt.Sprintf("%s IN (%s)", "app_id", strings.Trim(aCond, ","))
			}
			subSql = append(subSql, fmt.Sprintf("(%s)", pCond))
		} else if len(protectedObjectIDs) > 1 {
			cond := ""
			var subCond []string
			for _, id := range protectedObjectIDs {
				pCond := fmt.Sprintf("%s = %d", "protect_object_id", id)
				if len(applicationIDs[id]) == 1 {
					pCond += " AND " + fmt.Sprintf("%s = %d", "app_id", applicationIDs[id][0])
				} else if len(applicationIDs[id]) > 1 {
					aCond := ""
					for _, id := range applicationIDs[id] {
						aCond += fmt.Sprintf("'%d',", id)
					}
					pCond += " AND " + fmt.Sprintf("%s IN (%s)", "app_id", strings.Trim(aCond, ","))
				}
				subCond = append(subCond, fmt.Sprintf("(%s)", pCond))
			}
			cond = "(" + strings.Join(subCond, " OR ") + ")"
			subSql = append(subSql, cond)
		}
	case "!=":
		if len(protectedObjectIDs) == 1 {
			pCond := fmt.Sprintf("%s != %d", "protect_object_id", protectedObjectIDs[0])
			if len(applicationIDs[protectedObjectIDs[0]]) == 1 {
				pCond = fmt.Sprintf("%s = %d", "protect_object_id", protectedObjectIDs[0])
				pCond += " AND " + fmt.Sprintf("%s != %d", "app_id", applicationIDs[protectedObjectIDs[0]][0])
			} else if len(applicationIDs[protectedObjectIDs[0]]) > 1 {
				pCond = fmt.Sprintf("%s = %d", "protect_object_id", protectedObjectIDs[0])
				aCond := ""
				for _, id := range applicationIDs[protectedObjectIDs[0]] {
					aCond += fmt.Sprintf("'%d',", id)
				}
				pCond += " AND " + fmt.Sprintf("%s NOT IN (%s)", "app_id", strings.Trim(aCond, ","))
			}
			subSql = append(subSql, fmt.Sprintf("(%s)", pCond))
		} else if len(protectedObjectIDs) > 1 {
			cond := ""
			var subCond []string
			for _, id := range protectedObjectIDs {
				pCond := fmt.Sprintf("%s != %d", "protect_object_id", id)
				if len(applicationIDs[id]) == 1 {
					pCond = fmt.Sprintf("%s = %d", "protect_object_id", id)
					pCond += " AND " + fmt.Sprintf("%s != %d", "app_id", applicationIDs[id][0])
				} else if len(applicationIDs[id]) > 1 {
					pCond = fmt.Sprintf("%s = %d", "protect_object_id", id)
					aCond := ""
					for _, id := range applicationIDs[id] {
						aCond += fmt.Sprintf("'%d',", id)
					}
					pCond += " AND " + fmt.Sprintf("%s NOT IN (%s)", "app_id", strings.Trim(aCond, ","))
				}
				subCond = append(subCond, fmt.Sprintf("(%s)", pCond))
			}
			cond = "(" + strings.Join(subCond, " AND ") + ")"
			subSql = append(subSql, cond)
		}
	default:
		return "", errors.New("invalid filter operator")
	}
	condition := "(" + strings.Join(subSql, " OR ") + ")"
	if sql == "" {
		return condition, nil
	} else {
		return sql + " AND " + condition, nil
	}
}

func (s *ComplexInArrayQuery) Marshal() (datatypes.JSON, error) {
	if s == nil {
		return nil, nil
	}
	b, err := json.Marshal(*s)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (s *ComplexInArrayQuery) Unmarshal(value datatypes.JSON) error {
	if value == nil {
		return nil
	}
	err := json.Unmarshal(value, &s)
	if err != nil {
		return err
	}
	return nil
}

// GormQuerySQL 使用 gormDB 解析 ComplexInArrayQuery 并生成查询
func (s *ComplexInArrayQuery) GormQuerySQL(db *gorm.DB, name string) (*gorm.DB, error) {
	if s == nil {
		return db, nil
	}

	var protectedObjectIDs []uint64
	applicationIDs := make(map[uint64][]uint64)
	for _, c := range s.Value {
		if len(c.Children) > 0 {
			for _, obj := range c.Children {
				protectedObjectIDs = append(protectedObjectIDs, obj.Value)
				if len(obj.Children) > 0 {
					for _, o := range obj.Children {
						applicationIDs[obj.Value] = append(applicationIDs[obj.Value], o.Value)
					}
				}
			}
		}
	}

	query := db
	switch s.Operator {
	case "=":
		if len(protectedObjectIDs) == 1 {
			query = query.Where("protect_object_id = ?", protectedObjectIDs[0])
			if len(applicationIDs[protectedObjectIDs[0]]) == 1 {
				query = query.Where("app_id = ?", applicationIDs[protectedObjectIDs[0]][0])
			} else if len(applicationIDs[protectedObjectIDs[0]]) > 1 {
				query = query.Where("app_id IN ?", applicationIDs[protectedObjectIDs[0]])
			}
		} else if len(protectedObjectIDs) > 1 {
			// 创建OR子查询
			orCondition := db.Session(&gorm.Session{})
			for _, id := range protectedObjectIDs {
				subQuery := db.Session(&gorm.Session{}).Where("protect_object_id = ?", id)
				if len(applicationIDs[id]) == 1 {
					subQuery = subQuery.Where("app_id = ?", applicationIDs[id][0])
				} else if len(applicationIDs[id]) > 1 {
					subQuery = subQuery.Where("app_id IN ?", applicationIDs[id])
				}
				orCondition = orCondition.Or(subQuery)
			}
			query = query.Where(orCondition)
		}
	case "!=":
		if len(protectedObjectIDs) == 1 {
			if len(applicationIDs[protectedObjectIDs[0]]) == 0 {
				query = query.Where("protect_object_id != ?", protectedObjectIDs[0])
			} else if len(applicationIDs[protectedObjectIDs[0]]) == 1 {
				// 如果保护对象相同但应用ID不同
				query = query.Where(
					db.Where("protect_object_id = ?", protectedObjectIDs[0]).
						Where("app_id != ?", applicationIDs[protectedObjectIDs[0]][0]),
				)
			} else if len(applicationIDs[protectedObjectIDs[0]]) > 1 {
				// 如果保护对象相同但应用ID不在列表中
				query = query.Where(
					db.Where("protect_object_id = ?", protectedObjectIDs[0]).
						Where("app_id NOT IN ?", applicationIDs[protectedObjectIDs[0]]),
				)
			}
		} else if len(protectedObjectIDs) > 1 {
			// 对多个保护对象，所有条件需要用AND连接
			for _, id := range protectedObjectIDs {
				if len(applicationIDs[id]) == 0 {
					query = query.Where("protect_object_id != ?", id)
				} else if len(applicationIDs[id]) == 1 {
					subCond := db.Session(&gorm.Session{}).
						Where("protect_object_id = ?", id).
						Where("app_id != ?", applicationIDs[id][0])
					query = query.Where(subCond)
				} else if len(applicationIDs[id]) > 1 {
					subCond := db.Session(&gorm.Session{}).
						Where("protect_object_id = ?", id).
						Where("app_id NOT IN ?", applicationIDs[id])
					query = query.Where(subCond)
				}
			}
		}
	default:
		return nil, errors.New("invalid filter operator")
	}

	return query, nil
}
