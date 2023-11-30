package pb

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/labstack/echo/v5"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/apis"
	"github.com/pocketbase/pocketbase/plugins/migratecmd"
)

var App *pocketbase.PocketBase

func GenerateMigrations() error {
	path, err := os.Getwd()
	if err != nil {
		return err
	}
	migratecmd.MustRegister(App, App.RootCmd, migratecmd.Config{
		Automigrate: true, // auto creates migration files when making collection changes
		Dir:         fmt.Sprintf("%s/backend/migrations", path),
	})
	return nil
}

func init() {
	homeDir, _ := os.UserHomeDir()
	path := fmt.Sprintf("%s/.xsafe", homeDir)
	App = pocketbase.NewWithConfig(pocketbase.Config{
		DefaultDataDir: path,
	})
}

func disableAdminMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if strings.HasPrefix(c.Path(), "/_") {
				return apis.NewForbiddenError("You are not allowed to access this resource", nil)
			}
			return next(c)
		}
	}
}

func SetServeCommand() {
	port := os.Getenv("PORT")
	if _, err := strconv.Atoi(port); err != nil || port == "" {
		port = "8080"
	}
	App.RootCmd.SetArgs([]string{"serve", "--http", "0.0.0.0:" + port})
}

func StartServer() {
	if err := App.Start(); err != nil {
		log.Fatal(err)
	}
}
