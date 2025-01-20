package cmd

import (
	"log"
	"net/http"
	"time"

	"finance-tracker/internal/data"
	"finance-tracker/internal/mods"
	"finance-tracker/internal/mods/auth"
	"finance-tracker/internal/mods/auth/api"
	"finance-tracker/internal/mods/auth/biz"
	"finance-tracker/internal/mods/auth/dal"
	"finance-tracker/pkg/config"
	"finance-tracker/pkg/util"

	"github.com/gin-gonic/gin"
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

			userDal := dal.NewUserDal(db)
			userBiz := biz.NewUserBiz(userDal)
			userApi := api.NewUserApi(userBiz)

			// Initiate gin
			e := gin.Default()
			gAPI := e.Group(mods.ApiPrefix)
			v1 := gAPI.Group("v1")

			// Test API
			e.GET("/health", func(c *gin.Context) {
				util.ResOK(c)
			})

			a := auth.Auth{
				UserAPI: userApi,
			}
			a.RegisterV1Routers(v1)

			// Start server
			s := &http.Server{
				Addr:         cfg.AppConfig.Port,
				Handler:      e,
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
