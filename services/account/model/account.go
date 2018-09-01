package model

import "github.com/jinzhu/gorm"

// define unique constraint for account table
const (
	UniqueConstraintUsername = "accounts_username_key"
)

type Account struct {
	gorm.Model
	Username string `json:"username" form:"username" query:"username"`
	Password string `json:"password" form:"password" query:"password"`
}

type UsernameDuplicateError struct{}

// Error returns formatted error for unique Username constraint
func (e *UsernameDuplicateError) Error() string {
	return "Account with the username you have entered already exists"
}
