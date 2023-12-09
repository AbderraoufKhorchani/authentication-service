package helpers

import "gorm.io/gorm"

type LoginRequest struct {
	UserName string `json:"user_name"  binding:"required"`
	Password string `json:"password"  binding:"required"`
}

type Models struct {
	User            User
	UserJSONBinding UserJSONBinding
}

// User is the structure which holds one user from the database.
type User struct {
	gorm.Model
	UserName  string `json:"user_name"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"-"`
}

type UserJSONBinding struct {
	UserName  string `json:"user_name" binding:"required"`
	FirstName string `json:"first_name,omitempty"`
	LastName  string `json:"last_name,omitempty"`
	Password  string `json:"password"  binding:"required"` // Include the password field for binding
}

type ResetPasswordPlatload struct {
	UserName    string `json:"user_name" binding:"required"`
	Password    string `json:"password" binding:"required"`
	NewPassword string `json:"new_password" binding:"required"`
}
