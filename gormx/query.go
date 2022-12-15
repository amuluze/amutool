// Package gormx
// Date: 2022/9/30 17:34
// Author: Amu
// Description:
package gormx

import (
	"gorm.io/gorm"
)

type ListResult struct {
	List       interface{}       `json:"list"`
	Pagination *PaginationResult `json:"pagination,omitempty"`
}

type PaginationResult struct {
	Total int64 `json:"total"` // 数据总量
	Limit int   `json:"limit"` // 每页返回的数量
	Page  int   `json:"page"`  // 页数
}

type PaginationParam struct {
	Limit int `json:"limit"`
	Page  int `json:"page"`
}

func (a PaginationParam) GetLimit() int {
	return a.Limit
}

func (a PaginationParam) GetPage() int {
	return a.Page
}

func FindOne(db *gorm.DB, out interface{}) (bool, error) {
	result := db.First(out)
	if err := result.Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func FindPage(db *gorm.DB, pp PaginationParam, out interface{}) (int64, error) {
	var count int64
	err := db.Count(&count).Error
	if err != nil {
		return 0, err
	} else if count == 0 {
		return count, nil
	}

	limit, page := pp.GetLimit(), pp.GetPage()
	if limit > 0 && page > 0 {
		db = db.Offset((page - 1) * limit).Limit(limit)
	} else if page > 0 {
		db = db.Limit(limit)
	}

	err = db.Find(out).Error
	return count, err
}

func WrapPageQuery(db *gorm.DB, pp PaginationParam, out interface{}) (*PaginationResult, error) {
	total, err := FindPage(db, pp, out)
	if err != nil {
		return nil, err
	}

	return &PaginationResult{
		Total: total,
		Limit: pp.GetLimit(),
		Page:  pp.GetPage(),
	}, nil
}
