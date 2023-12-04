package user

import "context"

type User struct {
	ID       int64  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserReq struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type CreateUserRes struct {
	ID       string  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginUserReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginUserRes struct {
	accessToken string
	ID       string  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
}

// type ProfileUserReq struct {
// 	ID       int64  `json:"id" db:"id"`
// }

type ChangProfileUserRes struct {
	ID       string  `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type ProfileUserRes struct {
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}


type Repository interface {
	CreateUser(ctx context.Context, user *User) (*User, error)
	GetUserByEmail(ctx context.Context, email string) (*User, error)
	GetUserById(ctx context.Context, id string) (*User, error)
	ChangeProfile(ctx context.Context, user *User) (*User, error)
}

type Service interface {
	CreateUser(ctx context.Context, req *CreateUserReq) (*CreateUserRes, error)
	Login(ctx context.Context,req *LoginUserReq) (*LoginUserRes, error)
	Profile(c context.Context, id string) (*ProfileUserRes, error)
	ChangeProfile(c context.Context, id string, req *CreateUserReq) error
}


