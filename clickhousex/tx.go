// Package clickhousex
// Date: 2023/4/11 11:42
// Author: Amu
// Description: 批量写入封装
package clickhousex

import "github.com/jmoiron/sqlx"

type Tx struct {
	*sqlx.Tx
}

func (t *Tx) GetStmt(stmtString string) (*Stmt, error) {
	stmt, err := t.Preparex(stmtString)
	if err != nil {
		return nil, err
	}
	return &Stmt{stmt}, nil
}

func (t *Tx) Commit() error {
	return t.Commit()
}

func (t *Tx) Close() {
	t.Close()
}
