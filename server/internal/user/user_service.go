package user

import (
	"context"
	"fmt"
	"server/util"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

const secretKey string = "jwt_secret_key"

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

func (s *service) CreateUser(c context.Context, req *CreateUserReq) (*CreateUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// hash password
	hashedPassword, err := util.HashPassword(req.Password)
	if err != nil {
		return nil, err
	}
	u := &User{
		Username: req.Username,
		Email:    req.Email,
		Password: hashedPassword,
	}

	r, err := s.Repository.CreateUser(ctx, u)
	if err != nil {
		return nil, err
	}

	res := &CreateUserRes{
		ID:       strconv.Itoa(int(r.ID)),
		Username: r.Username,
		Email:    r.Email,
	}

	return res, nil
}


type JWTClaims struct {
	ID string `json:"id"`
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func(s *service) Login(ctx context.Context,req *LoginUserReq) (*LoginUserRes, error) {
	ctx, cancel := context.WithTimeout(ctx,s.timeout)
	defer cancel()

	u, err := s.Repository.GetUserByEmail(ctx, req.Email)
	if err != nil {
		return &LoginUserRes{}, err
	}

	err = util.CheckPassword(req.Password, u.Password)
	if err != nil {
		return &LoginUserRes{}, err
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, JWTClaims{
		ID: strconv.Itoa(int(u.ID)),
		Username: u.Username,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: strconv.Itoa(int(u.ID)),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	})

	accessTk, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return &LoginUserRes{}, err
	}
	
	return &LoginUserRes{accessToken: accessTk, Username: u.Username, ID: strconv.Itoa(int(u.ID))}, nil
}

func (s *service) Profile(c context.Context, id string) (*ProfileUserRes, error) {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	r, err := s.Repository.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}

	res := &ProfileUserRes{
		Username: r.Username,
		Email:    r.Email,
		Password: r.Password,
	}

	return res, nil
}


func(s *service) ChangeProfile(c context.Context, id string, req *CreateUserReq) error {
	ctx, cancel := context.WithTimeout(c, s.timeout)
	defer cancel()

	// convet str to int64
	id_int, err := strconv.ParseInt(id, 10, 64)
  if err != nil {
    return err
  }
	u := &User{
		ID: id_int,
		Username: req.Username,
		Password: req.Password,
		Email: req.Email, // not change email
	}
	fmt.Println("Change profile at service")
	_, err = s.Repository.ChangeProfile(ctx, u)
	if err != nil {
		return err
	}

	return nil
}	
