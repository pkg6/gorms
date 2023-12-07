package gorms

import "gorm.io/gorm"

type IModel interface {
	TableName() string
}

type IDModel struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

type TimeModel struct {
	CreatedAt int `json:"created_at"`
	UpdatedAt int `json:"updated_at"`
}

type SoftDeleteModel struct {
	Deleted gorm.DeletedAt
}
