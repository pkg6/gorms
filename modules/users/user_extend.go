package users

import "github.com/pkg6/gorms"

// UserExtend 用户基本信息扩展表
type UserExtend struct {
	gorms.IDModel
	gorms.TimeModel
	gorms.SoftModel
	UserID uint   `gorm:"comment:用户ID" json:"user_id"`
	User   *User  `json:"user"`
	Field  string `gorm:"type:varchar(50);comment:扩展字段" json:"field"`
	Value  string `gorm:"type:varchar(255);comment:扩展字段值" json:"value"`
}
