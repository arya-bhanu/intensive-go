package user

import (
	"context"

	"github.com/arya-bhanu/intensive-go/model"
)

type UserService interface {
	GetUser(ctx context.Context, userID int64) (*model.User, error)
	ValidateUser(ctx context.Context, userID int64) (bool, error)
}

type UserServiceServer struct {
	Users map[int64]*model.User
}

type GetUserRequest struct {
	UserId int64 `json:"user_id"`
}

type GetUserResponse struct {
	User *model.User `json:"user"`
}

type ValidateUserRequest struct {
	UserId int64 `json:"user_id"`
}

type ValidateUserResponse struct {
	Valid bool `json:"valid"`
}
