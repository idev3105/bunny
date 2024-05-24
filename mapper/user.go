package mapper

import (
	"github.com/jackc/pgx/v5/pgtype"
	userentity "org.idev.bunny/backend/domain/user/entity"
	sqlc_generated "org.idev.bunny/backend/generated/sqlc"
)

func MapToDomain(u sqlc_generated.User) *userentity.User {
	user := &userentity.User{}
	user.UserId = u.UserId
	user.CreatedAt = &u.CreatedAt.Time
	user.UpdatedAt = &u.UpdatedAt.Time
	user.CreatedBy = &u.CreatedBy.String
	user.UpdatedBy = &u.UpdatedBy.String
	user.Username = u.Username.String
	return user
}

func MapToSQLModel(u userentity.User) *sqlc_generated.User {
	user := &sqlc_generated.User{}
	user.UserId = u.UserId
	user.CreatedAt = pgtype.Timestamp{Time: *u.CreatedAt, Valid: true}
	user.UpdatedAt = pgtype.Timestamp{Time: *u.UpdatedAt, Valid: true}
	user.CreatedBy = pgtype.Text{String: *u.CreatedBy, Valid: true}
	user.UpdatedBy = pgtype.Text{String: *u.UpdatedBy, Valid: true}
	user.Username = pgtype.Text{String: u.Username, Valid: true}
	return user
}
