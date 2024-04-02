package users

import (
	"github.com/pkg6/gorms"
)

// User 用户表
type User struct {
	gorms.IDModel
	gorms.TimeModel
	gorms.SoftModel
	Username        string `gorm:"type:varchar(50);comment:用户名" json:"username"`
	DisplayName     string `gorm:"type:varchar(50);comment:用户显示的昵称" json:"display_name"`
	Email           string `gorm:"type:varchar(255);comment:邮箱地址" json:"email"`
	EmailVerifiedAt int64  `gorm:"comment:邮箱验证时间" json:"email_verified_at"`
	Phone           string `gorm:"varchar(25);comment:手机号" json:"phone"`
	PhoneVerifiedAt int64  `gorm:"comment:手机验证时间" json:"phone_verified_at"`
	Password        string `gorm:"type:varchar(255);comment:密码" json:"-"`
	Avatar          string `gorm:"type:varchar(255);comment:头像地址" json:"avatar"`
	Status          int    `gorm:"type:int(5);comment:用户状态" json:"status"`
	CurrentIP       string `gorm:"type:varchar(100);comment:当前所在ip" json:"current_ip"`
	//授权信息
	UserOauths []*UserOauth `json:"user_oauths"`
	//用户扩展字段
	UserExtends []*UserExtend `json:"user_extends"`
}
