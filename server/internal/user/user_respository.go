package user

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}

func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	var lastInsertID int
	query := "INSERT INTO users(username, password, email) VALUES($1,$2,$3) returning id"
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Password, user.Email).Scan(&lastInsertID)
	if err != nil {
		return &User{}, err
	}

	user.ID = int64(lastInsertID)
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	u := User{}
	query := "SELECT id, username, email, password FROM users WHERE email=$1"
	err := r.db.QueryRowContext(ctx, query, email).Scan(&u.ID, &u.Username, &u.Email, &u.Password)
	if err != nil {
		return &User{}, err
	}

	return &u, nil
}

func (r *repository) GetUserById(ctx context.Context, id string) (*User, error) {
	u := User{}
	query := "SELECT id, username, email FROM users WHERE id=$1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&u.ID, &u.Username, &u.Email)
	if err != nil {
		return &User{}, err
	}

	return &u, nil
}

func (r *repository) ChangeProfile(ctx context.Context, user *User) (*User, error) {
	fmt.Println("Change profile at repo")

	query := "UPDATE users SET username = $2, email = $3, password = $4 WHERE id = $1"

	_, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.Email, user.Password)
	if err != nil {
			return nil, err
	}

	updatedUser := User{}

	query = "SELECT id, username, email, password FROM users WHERE id = $1"
	err = r.db.QueryRowContext(ctx, query, user.ID).Scan(&updatedUser.ID, &updatedUser.Username, &updatedUser.Email, &updatedUser.Password)
	if err != nil {
			return nil, err
	}

	return &updatedUser, nil
}
