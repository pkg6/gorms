package gorms

import "gorm.io/gorm"

type IModel interface {
	TableName() string
}

type IDModel struct {
	ID uint `json:"id" xml:"ID" gorm:"primaryKey"`
}

type TimeModel struct {
	CreatedAt int64 `json:"created_at" xml:"CreatedAt"`
	UpdatedAt int64 `json:"updated_at" xml:"UpdatedAt"`
}

type SoftModel struct {
	DeletedAt gorm.DeletedAt `json:"deleted_at" xml:"DeletedAt" gorm:"index"`
}
