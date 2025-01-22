package auth

import (
	"finance-tracker/internal/mods/auth/api"
	"finance-tracker/internal/mods/auth/biz"
	"finance-tracker/internal/mods/auth/dal"

	"github.com/google/wire"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Auth), "*"),
	wire.Struct(new(dal.User), "*"),
	wire.Struct(new(biz.User), "*"),
	wire.Struct(new(api.User), "*"),
	wire.Struct(new(biz.Login), "*"),
	wire.Struct(new(api.Login), "*"),
)
