package service

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"secretKeeper/internal/server/middleware/auth"
	"secretKeeper/internal/server/model"
	"secretKeeper/internal/server/storage"
	"secretKeeper/pkg/apperr"
	"secretKeeper/pkg/crypt"
	"secretKeeper/pkg/jwt"
	pb "secretKeeper/proto"
)

type userGrpc struct {
	pb.UnimplementedUserServer

	storage    storage.UserServerStorage
	jwtManager jwt.Manager
	crypter    crypt.Crypter
}

// NewUserGrpc - creates new user grpc service.
func NewUserGrpc(s storage.UserServerStorage, m jwt.Manager, c crypt.Crypter) *userGrpc {
	return &userGrpc{
		storage:    s,
		jwtManager: m,
		crypter:    c,
	}
}

// RegisterService - registers service via grpc server.
func (u *userGrpc) RegisterService(r grpc.ServiceRegistrar) {
	pb.RegisterUserServer(r, u)
}

// Register - registers a new user.
//
// On successful creation returns JwtToken.
func (u *userGrpc) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	m := model.User{Login: in.Login, Password: in.Password}

	validate := validator.New()
	if errV := validate.Struct(m); errV != nil {
		return nil, status.Error(codes.InvalidArgument, errV.Error())
	}

	userModel, err := u.storage.Create(ctx, m)

	if errors.Is(err, apperr.ErrConflict) {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, errToken := u.jwtManager.Issue(userModel.ID.String())
	if errToken != nil {
		return nil, status.Error(codes.Internal, errToken.Error())
	}

	return &pb.RegisterResponse{Token: u.crypter.Encode(token)}, nil
}

// Login - Will return JwtToken on successful authentication via provided login and password.
func (u *userGrpc) Login(ctx context.Context, in *pb.LoginRequest) (*pb.LoginResponse, error) {
	userModel, err := u.storage.GetByLoginAndPassword(ctx, model.User{Login: in.Login, Password: in.Password})

	if errors.Is(err, pgx.ErrNoRows) {
		return nil, status.Error(codes.NotFound, err.Error())
	}

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	token, errToken := u.jwtManager.Issue(userModel.ID.String())
	if errToken != nil {
		return nil, status.Error(codes.Internal, errToken.Error())
	}

	return &pb.LoginResponse{Token: u.crypter.Encode(token)}, nil
}

// Delete - will delete a user from storage by provided ID.
func (u *userGrpc) Delete(ctx context.Context, in *pb.DeleteRequest) (*pb.DeleteResponse, error) {
	token := ctx.Value(auth.JwtTokenCtx{}).(string)

	uid, err := uuid.Parse(token)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	_, errDelete := u.storage.DeleteUser(ctx, model.User{ID: &uid})
	if errDelete != nil {
		if errors.Is(errDelete, pgx.ErrNoRows) {
			return nil, status.Error(codes.NotFound, errDelete.Error())
		}
		return nil, status.Error(codes.Internal, errDelete.Error())
	}

	return &pb.DeleteResponse{}, nil
}
