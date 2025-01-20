package cmd

import (
	"log"
	"net/http"
	"time"

	"finance-tracker/internal/data"
	"finance-tracker/pkg/config"
	"finance-tracker/routes"

	"github.com/urfave/cli/v2"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Action: func(c *cli.Context) error {
			// Load configuration
			cfg := config.LoadConfigs()

			// Initialize database
			dsn := "host=" + cfg.DatabaseConfig.Host + " user=" + cfg.DatabaseConfig.User + " password=" + cfg.DatabaseConfig.Password + " dbname=" + cfg.DatabaseConfig.Name + " port=" + cfg.DatabaseConfig.Port + " sslmode=disable"
			db := data.InitDatabase(dsn)
			// defer db.Close

			// Setup routes
			router := routes.SetupRoutes(db)

			// Start server
			s := &http.Server{
				Addr:         cfg.AppConfig.Port,
				Handler:      router,
				ReadTimeout:  time.Duration(cfg.AppTimeoutConfig.Read * int(time.Second)),
				WriteTimeout: time.Duration(cfg.AppTimeoutConfig.Write * int(time.Second)),
				IdleTimeout:  time.Duration(cfg.AppTimeoutConfig.Idle * int(time.Second)),
			}

			log.Println("Starting server on port " + cfg.AppConfig.Port + "...")
			if err := s.ListenAndServe(); err != nil {
				log.Fatalf("Could not start server: %s\n", err.Error())
			}

			return nil
		},
	}
}
