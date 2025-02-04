package auth

import (
	"finance-tracker/internal/mods/auth/api"

	"gorm.io/gorm"
)

type Auth struct {
	DB       *gorm.DB
	UserAPI  *api.User
	LoginAPI *api.Login
}

func (a *Auth) AutoMigrate() error {
	// fmt.Println("AutoMigrate")

	// return a.DB.AutoMigrate(
	// 	new(schema.User),
	// )

	return nil
}

func (a *Auth) Init() error {
	// if err := a.AutoMigrate(); err != nil {
	// 	return err
	// }

	return nil
}
