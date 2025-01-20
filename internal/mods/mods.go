package mods

import (
	"finance-tracker/internal/mods/auth"
)

const (
	ApiPrefix = "/api/"
)

// Collection of wire providers
// var Set = wire.NewSet(
// 	wire.Struct(new(Mods), "*"),
// 	auth.Set,
// )

type Mods struct {
	Auth *auth.Auth
}

func (a *Mods) Init() error {
	if err := a.Auth.Init(); err != nil {
		return err
	}

	return nil
}
