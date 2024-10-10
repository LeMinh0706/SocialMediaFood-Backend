package gapi

import (
	"context"
	"database/sql"
	"strings"

	"github.com/LeMinh0706/SocialMediaFood-Backend/internal/pb"
	"github.com/LeMinh0706/SocialMediaFood-Backend/util"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	hashPassword, err := util.HashPashword(req.GetPassword())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "fail to hash pw: %s", err)
	}
	fullname := req.GetFullname()
	if strings.TrimSpace(fullname) == "" {
		fullname = req.GetUsername()
	}
	var email sql.NullString
	if strings.TrimSpace(req.Email) == "" {
		email = sql.NullString{Valid: false}
	} else {
		email = sql.NullString{String: req.GetEmail(), Valid: true}
	}

	user, err := server.UserService.Register(ctx, req.GetUsername(), hashPassword, req.GetFullname(), email.String, req.GetGender())
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Error: %s", err)
	}
	res := &pb.CreateUserResponse{
		User: &pb.User{
			Username: user.Username,
			Fullname: user.Fullname,
			Email:    user.Email.String,
			Gender:   user.Gender,
			RoleId:   user.RoleID,
		},
	}
	return res, nil
}
