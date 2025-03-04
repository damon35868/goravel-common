package common

import (
	"github.com/goravel/framework/support/carbon"
	"gorm.io/gorm"
)

type OrmId struct {
	ID uint `gorm:"primaryKey" json:"id"`
}
type OrmTimestamps struct {
	CreatedAt *carbon.DateTime `gorm:"autoCreateTime;column:created_at" json:"createdAt"`
	UpdatedAt *carbon.DateTime `gorm:"autoUpdateTime;column:updated_at" json:"updatedAt"`
}
type OrmSoftDeletes struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"-"`
}

type OrmSoftDeletesShow struct {
	DeletedAt gorm.DeletedAt `gorm:"column:deleted_at" json:"deletedAt"`
}
