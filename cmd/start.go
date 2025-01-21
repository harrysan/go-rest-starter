package cmd

import (
	_ "finance-tracker/docs"
	"finance-tracker/internal/server"
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/config"

	// swagger embed files
	// gin-swagger middleware
	"github.com/urfave/cli/v2"
)

func Start() *cli.Command {
	return &cli.Command{
		Name:  "start",
		Usage: "Start server",
		Action: func(c *cli.Context) error {
			// Load configuration
			cfg := config.LoadConfigs()

			// Initialize dependencies
			handlers, err := wirex.InitializeDependencies()
			if err != nil {
				return err
			}

			// Create and start server
			srv := server.NewServer(handlers, cfg)
			server.StartServer(srv)

			return nil
		},
	}
}
