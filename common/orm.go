package common

import (
	"time"
)

type OrmId struct {
	ID uint `gorm:"primaryKey" json:"id"`
}
type OrmTimestamps struct {
	CreatedAt *time.Time `json:"createdAt"`
	UpdatedAt *time.Time `json:"updatedAt"`
}
type OrmSoftDeletes struct {
	DeletedAt *time.Time `json:"deletedAt"`
}
