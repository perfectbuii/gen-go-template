package deliveries

import (
	"context"
	"fmt"

	"github.com/perfectbuii/gen-go-template/internal/services"
	"github.com/perfectbuii/gen-go-template/pb"
	"github.com/perfectbuii/gen-go-template/transform"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type userDelivery struct {
	userService services.UserService
	pb.UnimplementedUserServiceServer
}

func NewUserDelivery(userService services.UserService) pb.UserServiceServer {
	return &userDelivery{
		userService: userService,
	}
}

func (d *userDelivery) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	if err := req.Validate(); err != nil {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Errorf("validate failed: %w", err).Error())
	}
	if err := d.userService.Create(ctx, transform.PbToUserPtr(req.GetUser())); err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Errorf("Create: %v", err).Error())
	}
	return &pb.CreateUserResponse{}, nil
}

func (d *userDelivery) UpdateUser(ctx context.Context, req *pb.UpdateUserRequest) (*pb.UpdateUserResponse, error) {
	if err := d.userService.Update(ctx, req.User.Id, transform.PbToUserPtr(req.GetUser())); err != nil {
		return nil, err
	}
	return &pb.UpdateUserResponse{
		Success: true,
	}, nil
}

func (d *userDelivery) GetList(ctx context.Context, req *pb.GetListUserRequest) (*pb.GetListUserResponse, error) {
	users, err := d.userService.GetList(ctx, int(req.Offset), int(req.Limit))
	if err != nil {
		return nil, err
	}
	return &pb.GetListUserResponse{
		Data: transform.UserToPbPtrList(users),
	}, nil
}

func (d *userDelivery) GetUserByID(ctx context.Context, req *pb.GetUserByIDRequest) (*pb.GetUserByIDResponse, error) {
	user, err := d.userService.GetByID(ctx, req.GetUserID())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserByIDResponse{
		Data: transform.UserToPbPtr(user),
	}, nil
}
