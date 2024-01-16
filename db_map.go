package gorms

import (
	"gorm.io/gorm"
	"reflect"
	"sync"
	"time"
)

const (
	MapDBDBName    = "_gorm_db"
	MapDBStructTag = "gorm"
)

type DBCallback func(db *gorm.DB) error

// DBCallbackOnErrorOnNotFound
//https://github.com/go-gorm/gorm/issues/3789
func DBCallbackOnErrorOnNotFound() DBCallback {
	return func(db *gorm.DB) error {
		return db.
			Callback().
			Query().
			Before("gorm:query").
			Register("disable_raise_record_not_found", func(db *gorm.DB) {
				db.Statement.RaiseErrorOnNotFound = false
			})
	}
}

// DBCallbackOnPool
//https://gorm.io/zh_CN/docs/generic_interface.html#%E8%BF%9E%E6%8E%A5%E6%B1%A0
func DBCallbackOnPool(maxIdleConns, maxOpenConns int, d time.Duration) DBCallback {
	return func(db *gorm.DB) error {
		sqlDB, err := db.DB()
		if err != nil {
			return err
		}
		// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
		sqlDB.SetMaxIdleConns(maxIdleConns)
		// SetMaxOpenConns 设置打开数据库连接的最大数量。
		sqlDB.SetMaxOpenConns(maxOpenConns)
		// SetConnMaxLifetime 设置了连接可复用的最大时间。
		sqlDB.SetConnMaxLifetime(d)
		return nil
	}
}

type DBMap struct {
	names            []string
	mapDB            map[string]*gorm.DB
	lock             *sync.Mutex
	registerCallback []DBCallback
}

func DefaultDBMap() *DBMap {
	return &DBMap{
		names: []string{},
		mapDB: map[string]*gorm.DB{},
		lock:  &sync.Mutex{},
		registerCallback: []DBCallback{
			DBCallbackOnErrorOnNotFound(),
		},
	}
}

func NewDBMap() *DBMap {
	return &DBMap{
		names:            []string{},
		mapDB:            map[string]*gorm.DB{},
		lock:             &sync.Mutex{},
		registerCallback: []DBCallback{},
	}
}

// AddStruct
// exp: Config
func (g *DBMap) AddStruct(s any) error {
	v := reflect.ValueOf(s)
	t := reflect.TypeOf(s)
	if t.Kind() == reflect.Ptr {
		for i := 0; i < v.Elem().NumField(); i++ {
			e := v.Elem().Field(i)
			if !e.IsZero() {
				if config, ok := e.Interface().(IConfig); ok {
					name := t.Elem().Field(i).Tag.Get(MapDBStructTag)
					if name == "" {
						name = t.Elem().Field(i).Name
					}
					db, err := config.DB()
					if err != nil {
						return err
					}
					if err = g.Register(name, db); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
	return NewErrMapDB("", ErrDBMapStruct)
}

func (g *DBMap) AddConfig(config IConfig) error {
	db, err := config.DB()
	if err != nil {
		return err
	}
	return g.Register(config.GetName(), db)
}

func (g *DBMap) OnRegisterCallback(callback DBCallback) *DBMap {
	g.registerCallback = append(g.registerCallback, callback)
	return g
}

func (g *DBMap) RegisterDB(db *gorm.DB) error {
	return g.Register(MapDBDBName, db)
}

func (g *DBMap) Register(name string, db *gorm.DB) error {
	if g.Exist(name) {
		return NewErrMapDB(name, ErrDBMapAddExist)
	}
	g.lock.Lock()
	defer g.lock.Unlock()
	for _, callback := range g.registerCallback {
		if err := callback(db); err != nil {
			return err
		}
	}
	g.mapDB[name] = db
	g.names = append(g.names, name)
	return nil
}

func (g *DBMap) Exist(name string) bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	_, ok := g.mapDB[name]
	return ok
}

func (g *DBMap) DB() *gorm.DB {
	return g.MustGet(MapDBDBName)
}

func (g *DBMap) MustGet(name string) *gorm.DB {
	db, err := g.Get(name)
	if err != nil {
		panic(err)
	}
	return db
}

func (g *DBMap) Get(name string) (*gorm.DB, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	db, ok := g.mapDB[name]
	if !ok {
		return nil, NewErrMapDB(name, ErrDBMapGetNotFind)
	}
	return db, nil
}

func (g *DBMap) GetNames() []string {
	return g.names
}
