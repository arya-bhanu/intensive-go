package user

import (
	"context"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/arya-bhanu/intensive-go/model"
)

func (s *UserServiceServer) GetUser(ctx context.Context, userID int64) (*model.User, error) {
	user, exists := s.Users[userID]
	if !exists {
		return nil, status.Errorf(codes.NotFound, "user not found")
	}
	return user, nil
}

// ValidateUser checks if a user exists and is active
func (s *UserServiceServer) ValidateUser(ctx context.Context, userID int64) (bool, error) {
	user, exists := s.Users[userID]
	if !exists {
		return false, status.Errorf(codes.NotFound, "user not found")
	}
	return user.Active, nil
}

func (s *UserServiceServer) GetUserRPC(ctx context.Context, req *GetUserRequest) (*GetUserResponse, error) {
	user, err := s.GetUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &GetUserResponse{User: user}, nil
}

func (s *UserServiceServer) ValidateUserRPC(ctx context.Context, req *ValidateUserRequest) (*ValidateUserResponse, error) {
	valid, err := s.ValidateUser(ctx, req.UserId)
	if err != nil {
		return nil, err
	}
	return &ValidateUserResponse{Valid: valid}, nil
}
