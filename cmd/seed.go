package cmd

import (
	"finance-tracker/internal/data"
	"finance-tracker/pkg/config"
	"finance-tracker/seeder"

	"github.com/urfave/cli/v2"
)

func SeedData() *cli.Command {
	return &cli.Command{
		Name:  "seed",
		Usage: "Seed Data for Testing",
		Action: func(c *cli.Context) error {
			cfg := config.LoadConfigs()
			dsn := "host=" + cfg.DatabaseConfig.Host +
				" user=" + cfg.DatabaseConfig.User +
				" password=" + cfg.DatabaseConfig.Password +
				" dbname=" + cfg.DatabaseConfig.Name +
				" port=" + cfg.DatabaseConfig.Port +
				" sslmode=disable"

			db := data.InitDatabase(dsn)

			s := seeder.Seed{
				DB: db,
			}

			if err := s.SeederData(); err != nil {
				return err
			}

			return nil
		},
	}
}
