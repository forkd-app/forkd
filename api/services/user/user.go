package user

import (
	"context"
	"forkd/db"
	"forkd/graph/model"
	"forkd/services/auth"
	"forkd/util"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

type UserService interface {
	GetByID(ctx context.Context, id uuid.UUID) (*model.User, error)
	GetByEmail(ctx context.Context, email string) (*model.User, error)
	GetCurrent(ctx context.Context) (*model.User, error)
	Update(ctx context.Context, input model.UserUpdateInput) (*model.User, error)
}

type userService struct {
	queries     *db.Queries
	authService auth.AuthService
}

// GetCurrent implements UserService.
func (u userService) GetCurrent(ctx context.Context) (*model.User, error) {
	user, _ := u.authService.GetUserSessionFromCtx(ctx)
	return model.UserFromDBType(*user), nil
}

// GetByEmail implements UserService.
func (u userService) GetByEmail(ctx context.Context, email string) (*model.User, error) {
	result, err := u.queries.GetUserByEmail(ctx, email)
	return util.HandleNoRowsOnNullableType(result, err, model.UserFromDBType)
}

// GetByID implements UserService.
func (u userService) GetByID(ctx context.Context, id uuid.UUID) (*model.User, error) {
	pgId := pgtype.UUID{
		Bytes: id,
		Valid: true,
	}
	result, err := u.queries.GetUserById(ctx, pgId)
	return util.HandleNoRowsOnNullableType(result, err, model.UserFromDBType)
}

// Update implements UserService.
func (u userService) Update(ctx context.Context, input model.UserUpdateInput) (*model.User, error) {
	user, _ := u.authService.GetUserSessionFromCtx(ctx)
	params := db.UpdateUserParams{
		ID: user.ID,
	}
	if input.DisplayName != nil && *input.DisplayName != "" {
		params.DisplayName = *input.DisplayName
	} else {
		params.DisplayName = user.DisplayName
	}
	updatedUser, err := u.queries.UpdateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	return model.UserFromDBType(updatedUser), nil
}

func New(queries *db.Queries, authService auth.AuthService) UserService {
	return userService{
		queries,
		authService,
	}
}
