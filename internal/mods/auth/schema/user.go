package schema

import (
	"time"

	"finance-tracker/pkg/crypto/hash"
	"finance-tracker/pkg/errors"

	"github.com/go-playground/validator/v10"
)

const (
	UserStatusActivated = "activated"
	UserStatusFreezed   = "freezed"
)

// User management for RBAC
type User struct {
	ID        int       `json:"id" gorm:"size:20;primarykey;"`         // Unique ID
	Username  string    `json:"username" gorm:"size:64,index:,unique"` // Username for login
	Name      string    `json:"name" gorm:"size:64"`                   // Name of user
	Password  string    `json:"-" gorm:"size:64"`                      // Password for login (encrypted)
	Phone     string    `json:"phone" gorm:"size:32"`                  // Phone number of user
	Email     string    `json:"email" gorm:"size:128,unique"`          // Email of user
	Remark    string    `json:"remark" gorm:"size:1024"`               // Remark of user
	Status    string    `json:"status" gorm:"size:20"`                 // Status of user (activated, freezed)
	CreatedAt time.Time `json:"created_at" gorm:"index"`               // Create time
	UpdatedAt time.Time `json:"updated_at" gorm:"index"`               // Update time
}

// Defining the slice of `User` struct.
type Users []*User

// Defining the query result for the `User` struct.
type UserQueryResult struct {
	Data Users
}

// Defining the data structure for creating a `User` struct.
type UserForm struct {
	Username string `json:"username" binding:"required,max=64"`                // Username for login
	Name     string `json:"name" binding:"required,max=64"`                    // Name of user
	Password string `json:"password" binding:"max=64"`                         // Password for login (md5 hash)
	Phone    string `json:"phone" binding:"max=32"`                            // Phone number of user
	Email    string `json:"email" binding:"max=128"`                           // Email of user
	Remark   string `json:"remark" binding:"max=1024"`                         // Remark of user
	Status   string `json:"status" binding:"required,oneof=activated freezed"` // Status of user (activated, freezed)
}

// A validation function for the `UserForm` struct.
func (a *UserForm) Validate() error {
	if a.Email != "" && validator.New().Var(a.Email, "email") != nil {
		return errors.BadRequest("", "Invalid email address")
	}
	return nil
}

// Convert `UserForm` to `User` object.
func (a *UserForm) FillTo(user *User) error {
	user.Username = a.Username
	user.Name = a.Name
	user.Phone = a.Phone
	user.Email = a.Email
	user.Remark = a.Remark
	user.Status = a.Status

	if pass := a.Password; pass != "" {
		hashPass, err := hash.GeneratePassword(pass)
		if err != nil {
			return errors.BadRequest("", "Failed to generate hash password: %s", err.Error())
		}
		user.Password = hashPass
	}

	return nil
}
