package seeder

import (
	"time"

	"finance-tracker/internal/mods/auth/schema"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/crypto/hash"
	"finance-tracker/pkg/errors"

	"gorm.io/gorm"
)

func SeedUsers(db *gorm.DB) error {
	cfg := config.LoadConfigs()

	users := []*schema.User{
		{
			Username:  cfg.RootConfig.Username,
			Name:      "Admin User",
			Password:  cfg.RootConfig.Password,
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
