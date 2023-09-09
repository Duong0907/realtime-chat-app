// Enterprise Business Rules Layer: contains Entities
// An entity can either be a core data structure necessary for the business rules (models)  
// or an object with methods that hold business logic in it (respository)

package user

import (
	// "context"
)

// Data models
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
	ID       string `json:"id" db:"id"`
	Username string `json:"username" db:"username"`
	Email    string `json:"email" db:"email"`
}

type LoginReq struct {
	Email    string `json:"email" db:"email"`
	Password string `json:"password" db:"password"`
}

type LoginRes struct {
	AccessToken string `json:"-"`
	ID       int64  `json:"id" db:"id"`
}

