package auth

import (
	"finance-tracker/internal/mods/auth/api"
	"finance-tracker/internal/mods/auth/schema"
	"fmt"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type Auth struct {
	DB      *gorm.DB
	UserAPI *api.User
}

func (a *Auth) AutoMigrate() error {
	fmt.Println("AutoMigrate")

	return a.DB.AutoMigrate(
		new(schema.User),
	)
}

func (a *Auth) Init() error {
	if err := a.AutoMigrate(); err != nil {
		return err
	}

	return nil
}

func (a *Auth) RegisterV1Routers(v1 *gin.RouterGroup) error {
	user := v1.Group("users")
	{
		user.GET("", a.UserAPI.Query)
		user.GET(":id", a.UserAPI.Get)
		user.POST("", a.UserAPI.Create)
		user.PUT(":id", a.UserAPI.Update)
		user.DELETE(":id", a.UserAPI.Delete)
	}

	return nil
}
