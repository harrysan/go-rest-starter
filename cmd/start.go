package cmd

import (
	"finance-tracker/internal/data"
	"finance-tracker/pkg/config"
	"finance-tracker/routes"
	"log"
	"net/http"

	"github.com/urfave/cli/v2"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Action: func(c *cli.Context) error {
			// Load configuration
			config.LoadConfigs()

			// Initialize database
			db := data.InitDatabase()
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
