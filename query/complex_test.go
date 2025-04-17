package query

import (
	"fmt"
	"skyview-go/common/clickhouse"
	"strings"
	"testing"

	"gorm.io/gorm"
)

func TestComplexObject(t *testing.T) {
	protectedObject1 := ProtectedObject{Key: "object", Value: 1, Children: []ProtectedObject{}}
	protectedObject2 := ProtectedObject{Key: "object", Value: 2, Children: []ProtectedObject{}}
	complexInArrayQuery := ComplexInArrayQuery{
		Operator: "=",
		Value: []ProtectedObject{{
			Key:   "group",
			Value: 1,
			Children: []ProtectedObject{
				protectedObject1, protectedObject2,
			},
		}},
	}
	//marshal, _ := json.Marshal(complexInArrayQuery)
	//fmt.Printf("%#v\n", string(marshal))
	sql, err := complexInArrayQuery.QuerySQL("", "")
	if err != nil {
		t.Errorf("Error in querying protected object: %s", err.Error())
	}
	t.Logf("sql: %s", sql)
}

func TestComplexObjectNot(t *testing.T) {
	protectedObject1 := ProtectedObject{Key: "object", Value: 1, Children: []ProtectedObject{}}
	protectedObject2 := ProtectedObject{Key: "object", Value: 2, Children: []ProtectedObject{}}
	complexInArrayQuery := ComplexInArrayQuery{
		Operator: "!=",
		Value:    []ProtectedObject{protectedObject1, protectedObject2},
	}
	//marshal, _ := json.Marshal(complexInArrayQuery)
	//fmt.Printf("%#v\n", string(marshal))
	sql, err := complexInArrayQuery.QuerySQL("", "")
	if err != nil {
		t.Errorf("Error in querying protected object: %s", err.Error())
	}
	t.Logf("sql: %s", sql)
}

func TestComplexInArrayQuerySingleToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()

	// 构建测试数据 - 单个保护对象，单个应用
	var children []ProtectedObject
	children = append(children, ProtectedObject{
		Key:   "app",
		Value: 456,
	})

	var objects []ProtectedObject
	objects = append(objects, ProtectedObject{
		Key:      "object",
		Value:    123,
		Children: children,
	})

	var sites []ProtectedObject
	sites = append(sites, ProtectedObject{
		Key:      "site",
		Value:    1,
		Children: objects,
	})

	saq := ComplexInArrayQuery{
		Operator: "=",
		Value:    sites,
	}

	gdb, err := saq.GormQuerySQL(db.DB, "")
	if err != nil {
		t.Error(err)
		return
	}

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println("单个保护对象，单个应用:", whereClause)
	fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))
}

func TestComplexInArrayQueryMultiAppToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()

	// 构建测试数据 - 单个保护对象，多个应用
	var children []ProtectedObject
	children = append(children, ProtectedObject{
		Key:   "app1",
		Value: 456,
	})
	children = append(children, ProtectedObject{
		Key:   "app2",
		Value: 789,
	})

	var objects []ProtectedObject
	objects = append(objects, ProtectedObject{
		Key:      "object",
		Value:    123,
		Children: children,
	})

	var sites []ProtectedObject
	sites = append(sites, ProtectedObject{
		Key:      "site",
		Value:    1,
		Children: objects,
	})

	saq := ComplexInArrayQuery{
		Operator: "=",
		Value:    sites,
	}

	gdb, err := saq.GormQuerySQL(db.DB, "")
	if err != nil {
		t.Error(err)
		return
	}

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println("单个保护对象，多个应用:", whereClause)
	fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))
}

func TestComplexInArrayQueryMultiObjectToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()

	// 构建测试数据 - 多个保护对象
	var children1 []ProtectedObject
	children1 = append(children1, ProtectedObject{
		Key:   "app1",
		Value: 456,
	})

	var children2 []ProtectedObject
	children2 = append(children2, ProtectedObject{
		Key:   "app2",
		Value: 789,
	})

	var objects []ProtectedObject
	objects = append(objects, ProtectedObject{
		Key:      "object1",
		Value:    123,
		Children: children1,
	})
	objects = append(objects, ProtectedObject{
		Key:      "object2",
		Value:    234,
		Children: children2,
	})

	var sites []ProtectedObject
	sites = append(sites, ProtectedObject{
		Key:      "site",
		Value:    1,
		Children: objects,
	})

	saq := ComplexInArrayQuery{
		Operator: "=",
		Value:    sites,
	}

	gdb, err := saq.GormQuerySQL(db.DB, "")
	if err != nil {
		t.Error(err)
		return
	}

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println("多个保护对象:", whereClause)
	fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))
}

func TestComplexInArrayQueryNotEqualToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()

	// 构建测试数据 - 用于!=操作符测试
	var children []ProtectedObject
	children = append(children, ProtectedObject{
		Key:   "app",
		Value: 456,
	})

	var objects []ProtectedObject
	objects = append(objects, ProtectedObject{
		Key:      "object",
		Value:    123,
		Children: children,
	})

	var sites []ProtectedObject
	sites = append(sites, ProtectedObject{
		Key:      "site",
		Value:    1,
		Children: objects,
	})

	saq := ComplexInArrayQuery{
		Operator: "!=",
		Value:    sites,
	}

	gdb, err := saq.GormQuerySQL(db.DB, "")
	if err != nil {
		t.Error(err)
		return
	}

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println("不等于条件:", whereClause)
}
