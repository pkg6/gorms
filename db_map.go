package gorms

import (
	"gorm.io/gorm"
	"log"
	"os"
	"reflect"
	"sync"
	"time"
)

const (
	dbMapPtrTag    = "gorm"
	DBMapENVMainDB = "main_db"
)

type DBCallback func(name string, db *gorm.DB) error

// DBCallbackOnErrorOnNotFound
//https://github.com/go-gorm/gorm/issues/3789
func DBCallbackOnErrorOnNotFound() DBCallback {
	return func(name string, db *gorm.DB) error {
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
	return func(name string, db *gorm.DB) error {
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
func DefaultDBMap() *DBMap {
	dbmap := NewDBMap()
	dbmap.OnRegisterCallback(DBCallbackOnErrorOnNotFound())
	return dbmap
}

func NewDBMap() *DBMap {
	return &DBMap{
		dbNames:          []string{},
		mapDB:            map[string]*gorm.DB{},
		lock:             &sync.Mutex{},
		registerCallback: []DBCallback{},
	}
}

type DBMap struct {
	dbNames          []string
	mapDB            map[string]*gorm.DB
	lock             *sync.Mutex
	registerCallback []DBCallback
}

// DB 根据定义环境变量读取db
func (g *DBMap) DB() *gorm.DB {
	if name := os.Getenv(DBMapENVMainDB); name != "" {
		return g.MustGet(name)
	}
	return nil
}

// RegisterByPtrConfig
// exp:
// type DBMapConfig struct {
//	Test *mysql.Config `gorm:"test"`
//}
func (g *DBMap) RegisterByPtrConfig(ptr any) error {
	v := reflect.ValueOf(ptr)
	t := reflect.TypeOf(ptr)
	if t.Kind() == reflect.Ptr {
		for i := 0; i < v.Elem().NumField(); i++ {
			e := v.Elem().Field(i)
			if !e.IsZero() {
				if config, ok := e.Interface().(IConfig); ok {
					name := t.Elem().Field(i).Tag.Get(dbMapPtrTag)
					if name == "" {
						name = t.Elem().Field(i).Name
					}
					if err := g.RegisterByConfig(config, name); err != nil {
						return err
					}
				}
			}
		}
		return nil
	}
	return NewErrMapDB("", ErrDBMapStruct)
}

// RegisterByConfig 根据config进行注册config
// exp
// mysql.Config
func (g *DBMap) RegisterByConfig(config IConfig, names ...string) error {
	db, err := config.DB()
	if err != nil {
		return err
	}
	name := config.GetName()
	if len(names) > 0 {
		name = names[0]
	}
	return g.Register(name, db)
}

// RegisterByNameDBConfig 批量注册db
func (g *DBMap) RegisterByNameDBConfig(config INameDBConfig) error {
	dbs, err := config.NameDB()
	if err != nil {
		return err
	}
	for name, db := range dbs {
		err = g.Register(name, db)
		if err != nil {
			return err
		}
	}
	return nil
}

// OnRegisterCallback 注入在注册的回掉函数
func (g *DBMap) OnRegisterCallback(callback DBCallback) *DBMap {
	g.registerCallback = append(g.registerCallback, callback)
	return g
}

// Register 注册db
func (g *DBMap) Register(name string, db *gorm.DB) error {
	g.lock.Lock()
	defer g.lock.Unlock()

	if _, ok := g.mapDB[name]; ok {
		log.Printf("[GORMS] [WARNING] %s already exists and will overwrite the connection with the original DB", name)
	}
	for _, callback := range g.registerCallback {
		if err := callback(name, db); err != nil {
			return err
		}
	}
	g.mapDB[name] = db
	g.dbNames = append(g.dbNames, name)
	return nil
}

// MustGet 根据name强制读取db，否则就panic
func (g *DBMap) MustGet(name string) *gorm.DB {
	db, err := g.Get(name)
	if err != nil {
		panic(err)
	}
	return db
}

// Get 根据name 获取db
func (g *DBMap) Get(name string) (*gorm.DB, error) {
	g.lock.Lock()
	defer g.lock.Unlock()
	db, ok := g.mapDB[name]
	if !ok {
		return nil, NewErrMapDB(name, ErrDBMapGetNotFind)
	}
	return db, nil
}

// Exist 判断namedb 是否存在
func (g *DBMap) Exist(name string) bool {
	g.lock.Lock()
	defer g.lock.Unlock()
	_, ok := g.mapDB[name]
	return ok
}

// GetDBNames 获取所有注册的数据别名
func (g *DBMap) GetDBNames() []string {
	return g.dbNames
}
