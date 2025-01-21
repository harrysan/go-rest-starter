package mods

import (
	"finance-tracker/internal/mods/auth"
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

const (
	ApiPrefix = "/api/v1/"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Mods), "*"),
	auth.Set,
)

type Mods struct {
	Auth *auth.Auth
}

func (a *Mods) RegisterRouters(e *gin.Engine, handlers *wirex.Handlers) error {
	gAPI := e.Group(ApiPrefix)

	gAPI.GET("/health", func(c *gin.Context) {
		util.ResOK(c)
	})

	user := gAPI.Group("users")
	{
		user.GET("", handlers.UserApi.Query)
		user.GET(":id", handlers.UserApi.Get)
		user.POST("", handlers.UserApi.Create)
		user.PUT(":id", handlers.UserApi.Update)
		user.DELETE(":id", handlers.UserApi.Delete)
	}

	// if err := a.Auth.RegisterV1Routers(gAPI); err != nil {
	// 	return err
	// }

	return nil
}
