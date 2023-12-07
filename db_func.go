package gorms

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

var (
	DefaultPageQueryPage        = "page"
	DefaultPageQuerySize        = "size"
	DefaultPageQueryMaxSize     = 100
	DefaultPageQueryDefaultSize = 10
)

// Inc
//字段加 5
// Inc(db.Table("test").Where("id = ?",1), "limit", 5)
func Inc(db *gorm.DB, column string, values ...any) error {
	var value any = 1
	if len(values) > 0 {
		value = values[0]
	}
	return db.UpdateColumn(column, gorm.Expr("`"+column+"` + ?", value)).Error
}

// Dec
//字段减 1
//Dec(db.Table("test").Where("id = ?",1), "limit", 5)
func Dec(db *gorm.DB, column string, values ...any) error {
	var value any = 1
	if len(values) > 0 {
		value = values[0]
	}
	return db.UpdateColumn(column, gorm.Expr("`"+column+"` - ?", value)).Error
}

type PageResponse[T any] struct {
	//当前分页数
	Page int `json:"page" xml:"Page"`
	//当前拉去多少条
	Size int `json:"size" xml:"Size"`
	//总数
	Total int64 `json:"total" xml:"Total"`
	//列表
	Items []T `json:"items" xml:"Items"`
}

// Paginate
//https://gorm.io/zh_CN/docs/scopes.html#%E5%88%86%E9%A1%B5
//db := db.Model(Model{})
//paginate, _ := Paginate[*Model](db, context.Request)
func Paginate[T any](db *gorm.DB, r *http.Request) (*PageResponse[T], error) {
	resp := &PageResponse[T]{}
	q := r.URL.Query()
	page, _ := strconv.Atoi(q.Get(DefaultPageQueryPage))
	if page <= 0 {
		page = 1
	}
	resp.Page = page
	size, _ := strconv.Atoi(q.Get(DefaultPageQuerySize))
	switch {
	case size > DefaultPageQueryMaxSize:
		size = DefaultPageQueryMaxSize
	case size <= 0:
		size = DefaultPageQueryDefaultSize
	}
	resp.Size = size
	var total int64
	_ = db.Count(&total).Error
	resp.Total = total
	var items []T
	err := db.Scopes(func(db *gorm.DB) *gorm.DB {
		if page >= 0 {
			offset := (page - 1) * size
			return db.Offset(offset).Limit(size)
		}
		return db
	}).Find(&items).Error
	resp.Items = items
	return resp, err
}
