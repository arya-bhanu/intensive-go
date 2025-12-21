package user

import (
	"github.com/arya-bhanu/intensive-go/model"
)

func NewUserServiceServer() *UserServiceServer {
	users := map[int64]*model.User{
		1: {ID: 1, Username: "alice", Email: "alice@example.com", Active: true},
		2: {ID: 2, Username: "bob", Email: "bob@example.com", Active: true},
		3: {ID: 3, Username: "charlie", Email: "charlie@example.com", Active: false},
	}
	return &UserServiceServer{Users: users}
}
