package query

import (
	"fmt"
	"strings"
	"testing"

	"gorm.io/gorm"
)

func TestInt32InArrayQueryMultiToSQL(t *testing.T) {
	saq := Int32InArrayQuery{
		Operator: "=",
		Value:    []int32{123, 32},
	}
	fmt.Println(saq.QuerySQL("src_port", ""))
}

func TestInt32InArrayQueryOneToSQL(t *testing.T) {
	saq := Int32InArrayQuery{
		Operator: "!=",
		Value:    []int32{123},
	}
	fmt.Println(saq.QuerySQL("src_port", ""))
}

func TestInt32InArrayMarshal(t *testing.T) {
	saq := Int32InArrayQuery{
		Operator: "=",
		Value:    []int32{23},
	}
	fmt.Println(saq.Marshal())
}

func TestInt32InArrayUnmarshal(t *testing.T) {
	saq := Int32InArrayQuery{
		Operator: "=",
		Value:    []int32{213, 232},
	}
	req, _ := saq.Marshal()
	ss := Int32InArrayQuery{}
	err := ss.Unmarshal(req)
	if err != nil {
		t.Error(err)
	}
	t.Logf("ss: %#v", ss)
}

func TestInt32InArrayQueryMultiToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()
	saq := Int32InArrayQuery{
		Operator: "=",
		Value:    []int32{123, 345},
	}
	gdb, _ := saq.GormQuerySQL(db.DB, "id")

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println(whereClause)
	fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))
}

func TestInt32InArrayQueryOneToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()
	saq := Int32InArrayQuery{
		Operator: "=",
		Value:    []int32{123},
	}
	gdb, _ := saq.GormQuerySQL(db.DB, "id")

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println(whereClause)
}

func TestInt32InArrayQueryGTToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()
	saq := Int32InArrayQuery{
		Operator: ">",
		Value:    []int32{100},
	}
	gdb, _ := saq.GormQuerySQL(db.DB, "id")

	// 使用DryRun获取SQL语句而不执行查询
	stmt := gdb.Session(&gorm.Session{DryRun: true}).Find(&struct{}{}).Statement

	// 只打印WHERE条件部分
	whereClause := stmt.SQL.String()
	if idx := strings.Index(strings.ToUpper(whereClause), "WHERE"); idx >= 0 {
		whereClause = whereClause[idx+6:]
	}
	fmt.Println(whereClause)
	fmt.Println(db.Dialector.Explain(stmt.SQL.String(), stmt.Vars...))
}

func TestInt32QueryToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()
	sq := Int32Query{
		Operator: "=",
		Value:    123,
	}
	gdb, err := sq.GormQuerySQL(db.DB, "skynet_rule_id")
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
	fmt.Println(whereClause)
}

func TestInt32QueryInvalidFieldToGormSQL(t *testing.T) {
	clickhouse.InitGormDB()
	db := clickhouse.GetGormDB()
	sq := Int32Query{
		Operator: "=",
		Value:    123,
	}
	_, err := sq.GormQuerySQL(db.DB, "other_field")
	if err == nil {
		t.Error("Expected error for invalid field, but got nil")
	} else {
		fmt.Println("Expected error:", err)
	}
}
