// Enterprise Business Rules Layer: contains Entities
// An entity can either be a core data structure necessary for the business rules (models)  
// or an object with methods that hold business logic in it (respository)

package user

import (
	"context"
	"database/sql"
	"log"
)

// Database Transaction interface
// Implemented by sql.DB
type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context,  string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

// Should create interface right before using it
type Repository interface {
	CreateUser(context.Context, *User) (*User, error)
	GetUserByEmail(context.Context, string) (*User, error)
}

// Object which contains business logic: CreateUser, ...
type repository struct {
	db DBTX
}

func NewRepository(db DBTX) Repository {
	return &repository{db: db}
}


func (r *repository) CreateUser(ctx context.Context, user *User) (*User, error) {
	query := `
		insert into "user" (username, email, password) 
		values ($1, $2, $3)
		returning id;
	`
	var lastInsertID int64
	err := r.db.QueryRowContext(ctx, query, user.Username, user.Email, user.Password).Scan(&lastInsertID) 
	if err != nil {
		log.Printf("Query error: %s\n", err.Error())
		return nil, err
	}

	user.ID = lastInsertID
	return user, nil
}

func (r *repository) GetUserByEmail(ctx context.Context, email string) (*User, error) {
	var user User

	query := `select id, username, password from "user" where email=$1`
	err := r.db.QueryRowContext(ctx, query, email).Scan(
		&user.ID,
		&user.Username,
		&user.Password,
	)
	return &user, err
}

