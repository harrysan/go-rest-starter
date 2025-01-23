package cmd

import (
	_ "finance-tracker/docs"
	"finance-tracker/internal/server"
	"finance-tracker/internal/wirex"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/redis"
	"log"

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
			log.Println("Load Configurations...")
			cfg := config.LoadConfigs()

			// Initialize Redis client
			log.Println("Initializing Redis...")
			redisClient := redis.NewRedisClient(cfg)

			// Initialize dependencies
			log.Println("Initializing Dependencies...")
			handlers, err := wirex.InitializeDependencies()
			if err != nil {
				return err
			}

			// Create and start server
			log.Println("Start New Server...")
			log.Println("----------------------------------------------------------------------------")
			srv := server.NewServer(handlers, redisClient, cfg)
			server.StartServer(srv)

			return nil
		},
	}
}
