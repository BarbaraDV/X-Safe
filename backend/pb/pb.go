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
	"github.com/pocketbase/pocketbase/core"
	"github.com/pocketbase/pocketbase/models"
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
	fn := func(e *core.ModelEvent) error {
		if e.Model.TableName() != "vulnerabilities" {
			return nil
		}
		data, _ := App.Dao().FindRecordById("vulnerabilities", e.Model.GetId())
		query := App.Dao().RecordQuery("inactive_categories")
		records := []*models.Record{}
		if err := query.All(&records); err != nil {
			log.Println(err)
			return nil
		}
		for _, r := range records {
			record, err := App.Dao().FindRecordById("categories", r.GetId())
			if err != nil {
				log.Println(err)
				return nil
			}
			if data.Get("category") != record.GetId() {
				if err := App.Dao().DeleteRecord(record); err != nil {
					log.Println(err)
					return nil
				}
			}
		}
		return nil
	}
	App.OnModelAfterCreate("vulnerabilities").Add(fn)
	App.OnModelAfterDelete("vulnerabilities").Add(fn)
	App.OnModelAfterUpdate("vulnerabilities").Add(fn)
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
