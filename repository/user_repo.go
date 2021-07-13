package repository

import (
	"context"
	"go-backend/model"
)

type UserPepo interface {
	SaveUser(context context.Context, user model.User) (model.User, error)
	SelectUserByEmail(context context.Context, email string) (model.User, error)
	SelectUsserById(context context.Context, userID string) (model.User, error)
	SelectUsers(context context.Context) ([]model.User, error)
}
