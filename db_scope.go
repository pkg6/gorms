package gorms

import (
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

// ScopesPage
//db.Scopes(Paginate(r)).Find(&users)
func ScopesPage(page, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page >= 0 {
			offset := (page - 1) * size
			return db.Offset(offset).Limit(size)
		}
		return db
	}
}

// ScopesWhereQueryArgs
//args = NewSQLArgs()
//db.Scopes(ScopesWhereQueryArgs(args)).Find(&users)
func ScopesWhereQueryArgs(args *QueryArgs) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		query, arg := args.WhereQueryArgs()
		if query != "" {
			return db.Where(query, arg...)
		}
		return db
	}
}

var (
	DefaultPageQueryPage        = "page"
	DefaultPageQuerySize        = "size"
	DefaultPageQueryMaxSize     = 100
	DefaultPageQueryDefaultSize = 10
)

type PageResponse[T any] struct {
	//当前分页数
	CurrentPage int `json:"current_page" xml:"CurrentPage"`
	//当前拉去多少条
	CurrentSize int `json:"current_size" xml:"CurrentSize"`
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
	resp.CurrentPage = page
	size, _ := strconv.Atoi(q.Get(DefaultPageQuerySize))
	switch {
	case size > DefaultPageQueryMaxSize:
		size = DefaultPageQueryMaxSize
	case size <= 0:
		size = DefaultPageQueryDefaultSize
	}
	resp.CurrentSize = size
	var total int64
	_ = db.Count(&total).Error
	resp.Total = total
	var items []T
	err := db.Scopes(ScopesPage(resp.CurrentPage, resp.CurrentSize)).Find(&items).Error
	resp.Items = items
	return resp, err
}
