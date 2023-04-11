// Package clickhouse
// Date: 2023/4/11 14:42
// Author: Amu
// Description:
package clickhouse

import "github.com/jmoiron/sqlx"

type Stmt struct {
	*sqlx.Stmt
}

func (s *Stmt) Append(v ...interface{}) error {
	_, err := s.Exec(v...)
	if err != nil {
		return err
	}
	return nil
}

func (s *Stmt) Close() {
	s.Close()
}
