// Application Bussiness Rules Layer/Use case Layer
// This layer holds Use Cases. As the name suggests, 
// it should provide every use case of the application. 
// i.e. it holds each and every functionality provided by the application: CreateUser, ...

package user

import (
	"context"
	"strconv"
	"time"
	// "errors"
	"go-chat/internal/authorization"
	"go-chat/util"
	"github.com/golang-jwt/jwt/v5"
)



type Service interface {
	CreateUser(context.Context, *CreateUserReq) (*CreateUserRes, error)
	Login(context.Context, *LoginReq) (*LoginRes, error)
}


type service struct {
	Repository
	timeout time.Duration
}

func NewService(repository Repository) Service {
	return &service{
		repository,
		time.Duration(2) * time.Second,
	}
}

func (s *service) CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}

	u := &User{
		Username: req.Username,
		Email: req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err 
	}

	res := &CreateUserRes{
		ID: strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email: r.Email,
	}

	return res, nil 
}

func (s *service) Login(ctx context.Context, req *LoginReq) (*LoginRes, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()
	
	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	if err := util.CheckPassword(req.Password, u.Password); err != nil {
		return nil, err
	}

	ss, err := auth.NewTokenStringWithClaims(auth.MyJWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 24)),
		},
	})

	if err != nil {
		return nil, err
	}

	return &LoginRes{
		AccessToken: ss,
		ID: u.ID,
	}, nil
}
