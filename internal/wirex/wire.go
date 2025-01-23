//go:build wireinject
// +build wireinject

package wirex

import (
	"finance-tracker/internal/data"
	"finance-tracker/internal/mods"
	"finance-tracker/internal/mods/auth/api"
	"finance-tracker/pkg/config"
	rds "finance-tracker/pkg/redis"

	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

type Handlers struct {
	UserApi  *api.User
	LoginApi *api.Login
}

// ProvideConfig menyediakan konfigurasi aplikasi.
func ProvideConfig() config.App {
	return config.LoadConfigs()
}

// ProvideDatabase menyediakan koneksi database.
func ProvideDatabase(cfg config.App) *gorm.DB {
	dsn := "host=" + cfg.DatabaseConfig.Host +
		" user=" + cfg.DatabaseConfig.User +
		" password=" + cfg.DatabaseConfig.Password +
		" dbname=" + cfg.DatabaseConfig.Name +
		" port=" + cfg.DatabaseConfig.Port +
		" sslmode=disable"

	return data.InitDatabase(dsn)
}

func ProvideRedis(cfg config.App) *redis.Client {
	return rds.NewRedisClient(cfg)
}

func InitializeDependencies() (*Handlers, error) {
	wire.Build(
		ProvideConfig,
		ProvideDatabase,
		ProvideRedis,
		mods.Set,
		wire.Struct(new(Handlers), "*"),
	) // end
	return nil, nil
}
