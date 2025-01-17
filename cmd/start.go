package cmd

import (
	"log"
	"net/http"

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
			log.Println("Starting server on port 8080...")
			if err := http.ListenAndServe(":8080", router); err != nil {
				log.Fatalf("Could not start server: %s\n", err.Error())
			}

			return nil
		},
	}
}
