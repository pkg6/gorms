package users

import "gorm.io/gorm"

// AutoMigrate
// 同步用户表结构
func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&User{},
		&UserOauth{},
		&UserExtend{},
	)
}
