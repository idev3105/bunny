package commonentity

import "time"

type BaseEntity struct {
	CreatedAt *time.Time `json:"created_at,omitempty"`
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	CreatedBy *string    `json:"created_by,omitempty"`
	UpdatedBy *string    `json:"updated_by,omitempty"`
}
