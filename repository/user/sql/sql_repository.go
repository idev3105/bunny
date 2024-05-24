package usersql

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
	userentity "org.idev.bunny/backend/domain/user/entity"
	sqlc_generated "org.idev.bunny/backend/generated/sqlc"
	"org.idev.bunny/backend/mapper"
)

type UserSqlRepo struct {
	db      sqlc_generated.DBTX
	queries *sqlc_generated.Queries
}

// create new instance of sql repository
func NewSqlRepository(db sqlc_generated.DBTX) *UserSqlRepo {
	return &UserSqlRepo{db: db, queries: sqlc_generated.New(db)}
}

func (r *UserSqlRepo) FindByUserId(ctx context.Context, userId string) (*userentity.User, error) {
	user, err := r.queries.FindUserByUserId(ctx, userId)
	if err != nil {
		return nil, err
	}
	return mapper.MapToDomain(user), nil
}

func (r *UserSqlRepo) Save(ctx context.Context, user *userentity.User) (*userentity.User, error) {
	prams := sqlc_generated.SaveUserParams{UserId: user.UserId,
		Username: pgtype.Text{String: user.Username, Valid: true},
	}
	if user.CreatedBy != nil {
		prams.CreatedBy = pgtype.Text{String: *user.CreatedBy, Valid: true}
	}
	if user.UpdatedBy != nil {
		prams.UpdatedBy = pgtype.Text{String: *user.UpdatedBy, Valid: true}
	}
	_, err := r.queries.SaveUser(ctx, prams)
	if err != nil {
		return nil, err
	}
	return user, nil
}
