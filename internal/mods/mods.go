package mods

import (
	"finance-tracker/internal/mods/auth"

	"github.com/google/wire"
)

// Collection of wire providers
var Set = wire.NewSet(
	wire.Struct(new(Mods), "*"),
	auth.Set,
)

type Mods struct {
	Auth *auth.Auth
}
