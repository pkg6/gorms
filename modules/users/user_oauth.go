package users

import "github.com/pkg6/gorms"

// UserOauth  用户授权表
type UserOauth struct {
	gorms.IDModel
	gorms.TimeModel
	gorms.SoftModel
	UserID     uint   `gorm:"comment:用户ID" json:"user_id"`
	User       *User  `json:"user"`
	OauthType  string `gorm:"type:varchar(50);comment:第三方登陆类型" json:"oauth_type"`
	OauthID    string `gorm:"type:varchar(255);comment:第三方ID" json:"oauth_id"`
	UnionID    string `gorm:"type:varchar(255);comment:主体ID" json:"unionid"`
	Credential string `gorm:"type:varchar(255);comment:密码凭证" json:"credential"`
}
