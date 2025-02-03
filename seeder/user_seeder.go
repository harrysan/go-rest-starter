package seeder

import (
	"finance-tracker/internal/mods/auth/schema"
	"finance-tracker/pkg/crypto/hash"
	"finance-tracker/pkg/errors"
	"time"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	users := []*schema.User{
		{
			Username:  "admin",
			Name:      "Admin User",
			Password:  "password",
			Email:     "admin@example.com",
			Status:    "activated",
			CreatedAt: time.Now(),
		},
	}

	for _, user := range users {
		var count int64
		err := db.Model(&schema.User{}).Where("username = ?", user.Username).Count(&count).Error
		if err != nil {
			return err
		}

		if count == 0 {
			hashPass, err := hash.GeneratePassword(user.Password)
			if err != nil {
				return errors.BadRequest("", "Failed to generate hash password: %s", err.Error())
			}
			user.Password = hashPass

			return db.Create(user).Error
		}
	}

	return nil
}
