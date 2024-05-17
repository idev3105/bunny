package userentity

import commonentity "org.idev.bunny/backend/domain/common/entity"

type User struct {
	commonentity.BaseEntity
	UserId   string `json:"userId,omitempty"`
	Username string `json:"username,omitempty"`
}
