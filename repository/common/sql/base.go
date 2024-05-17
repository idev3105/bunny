package commonsql

import "time"

const (
	ID        = "id"
	CreatedAt = "created_at"
	UpdatedAt = "updated_at"
	CreatedBy = "created_by"
	UpdatedBy = "updated_by"
	Metadata  = "metadata"
)

type BaseSQLModel struct {
	CreatedAt *time.Time `gorm:"column:created_at"`
	UpdatedAt *time.Time `gorm:"column:updated_at"`
	CreatedBy *string    `gorm:"column:created_by"`
	UpdatedBy *string    `gorm:"column:updated_by"`
	Metadata  *string    `gorm:"column:metadata"`
}
